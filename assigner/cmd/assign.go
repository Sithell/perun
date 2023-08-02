package main

import (
	"context"
	"encoding/json"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sithell/perun/api/models"
	"github.com/sithell/perun/internal/database"
	"github.com/sithell/perun/provider/pb"
	flag "github.com/spf13/pflag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"strings"
	"time"
)

var (
	dbHost     string
	dbUser     string
	dbPassword string
	dbPort     uint
	dbName     string
	mqHost     string
	mqUser     string
	mqPassword string
	mqPort     uint
)

func init() {
	flag.StringVar(&dbHost, "db-host", "localhost", "database host")
	flag.StringVar(&dbUser, "db-user", "perun", "database user")
	flag.UintVar(&dbPort, "db-port", 5432, "database port")
	flag.StringVar(&dbName, "db-name", "perun", "database name")
	dbPassword = os.Getenv("DATABASE_PASSWORD")
	flag.StringVar(&mqHost, "mq-host", "localhost", "message queue host")
	flag.StringVar(&mqUser, "mq-user", "guest", "message queue user")
	flag.UintVar(&mqPort, "mq-port", 5672, "message queue port")
	mqPassword = os.Getenv("MESSAGE_QUEUE_PASSWORD")
}

type Provider struct {
	Provider database.Provider
	Client   pb.ProviderClient
}

func main() {
	ctx := context.Background()
	db, err := database.InitDB(dbHost, dbUser, dbPassword, dbName, dbPort)
	if err != nil {
		log.Fatalf("failed to init db: %v", err)
	}

	var providers []database.Provider
	result := db.Find(&providers)
	if result.Error != nil {
		log.Fatalf("failed to fetch providers from db: %v", result.Error)
	}
	var providerClients []Provider
	for _, provider := range providers {
		conn, err := grpc.Dial(provider.Host, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Printf("WARN: failed to connect to provider #%d: %v", provider.ID, err)
		}
		providerClients = append(providerClients, Provider{
			Provider: provider,
			Client:   pb.NewProviderClient(conn),
		})
	}
	if len(providerClients) == 0 {
		log.Fatalf("No providers to assign jobs to!")
	}

	mq, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/", mqUser, mqPassword, mqHost, mqPort))
	if err != nil {
		log.Fatalf("failed to init mq: %v", err)
	}
	defer func(mq *amqp.Connection) {
		err := mq.Close()
		if err != nil {
			log.Printf("WARN: failed to close mq connection: %v", err)
		}
	}(mq)
	ch, err := mq.Channel()
	if err != nil {
		log.Fatalf("failed to open channel: %v", err)
	}
	defer func(ch *amqp.Channel) {
		err := ch.Close()
		if err != nil {
			log.Printf("WARN: failed to close channel: %v", err)
		}
	}(ch)
	q, err := ch.QueueDeclare(
		"run-container",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("failed to declare queue: %v", err)
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			var job models.Job
			err := json.Unmarshal(d.Body, &job)
			if err != nil {
				log.Printf("WARN: failed to unmarshal job: %v", err)
				continue
			}
			fmt.Printf("Job(id=%d, image=%s, command=%s)\n", job.ID, job.Image, job.Command)
			var runs []database.Run
			result = db.Where("runs.job_id = ?", job.ID).Find(&runs)
			if result.Error != nil {
				log.Printf("WARN: failed to fetch runs from db: %v", result.Error)
				continue
			}
			if len(runs) == 0 {
				// job was never run before
				assignedProvider := providerClients[time.Now().Nanosecond()%len(providerClients)]
				run := database.Run{
					JobID:    uint(job.ID),
					Provider: assignedProvider.Provider,
					Status:   "started",
				}
				result := db.Save(&run)
				if result.Error != nil {
					log.Printf("WARN: failed to save run in db: %v", result.Error)
					continue
				}
				containerInfo, err := assignedProvider.Client.RunContainer(ctx, &pb.RunContainerParams{
					Image: *job.Image,
					Cmd:   strings.Split(job.Command, " "),
				})
				if err != nil {
					log.Printf("WARN: failed to run container: %v", err)
					continue
				}
				fmt.Printf("Started container id=%s", containerInfo.Id)
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
