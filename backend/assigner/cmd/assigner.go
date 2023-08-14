package main

import (
	"context"
	"encoding/json"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sithell/perun/backend/api/models"
	connector "github.com/sithell/perun/backend/connector/pb"
	"github.com/sithell/perun/backend/internal/database"
	flag "github.com/spf13/pflag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

var (
	dbHost        string
	dbUser        string
	dbPassword    string
	dbPort        uint
	dbName        string
	mqHost        string
	mqUser        string
	mqPassword    string
	mqPort        uint
	connectorHost string
	connectorPort uint
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
	flag.StringVar(&connectorHost, "connector-host", "localhost", "connector host")
	flag.UintVar(&connectorPort, "connector-port", 9002, "connector port")
}

func availableProviders(ctx context.Context, db *gorm.DB) ([]database.Provider, error) {
	var providers []database.Provider
	q := db.Find(&providers)
	if q.Error != nil {
		return nil, fmt.Errorf("failed to fetch providers from db: %w", q.Error)
	}

	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", connectorHost, connectorPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("WARN: failed to connect to connector: %w", err)
	}
	client := connector.NewApiClient(conn)
	connections, err := client.GetActiveConnections(ctx, &connector.Empty{})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch active connections from connector: %w", err)
	}

	var result []database.Provider
	// O(n^2)
	for _, provider := range providers {
		for _, connection := range connections.Connections {
			if uint64(provider.ID) == connection.ProviderId {
				result = append(result, provider)
			}
		}
	}
	return result, nil
}

func main() {
	ctx := context.Background()
	db, err := database.InitDB(dbHost, dbUser, dbPassword, dbName, dbPort)
	if err != nil {
		log.Fatalf("failed to init db: %v", err)
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
			result := db.Where("runs.job_id = ?", job.ID).Find(&runs)
			if result.Error != nil {
				log.Printf("WARN: failed to fetch runs from db: %v", result.Error)
				continue
			}
			if len(runs) == 0 {
				// job was never run before
				var providers []database.Provider
				providers, err = availableProviders(ctx, db)
				if err != nil {
					log.Printf("ERROR: failed to get avaiable providers: %v", err)
					continue
				}

				assignedProvider := providers[time.Now().Nanosecond()%len(providers)]

				run := database.Run{
					JobID:    uint(job.ID),
					Provider: assignedProvider,
					Status:   "assigned",
				}
				result = db.Save(&run)
				if result.Error != nil {
					log.Printf("WARN: failed to save run in db: %v", result.Error)
					continue
				}

				conn, err := grpc.Dial(fmt.Sprintf("%s:%d", connectorHost, connectorPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
				if err != nil {
					fmt.Printf("WARN: failed to connect to connector: %v", err)
					continue
				}
				client := connector.NewApiClient(conn)

				containerInfo, err := client.RunContainer(ctx, &connector.RunContainerParams{
					Image:      *job.Image,
					Cmd:        job.Command,
					ProviderId: uint64(assignedProvider.ID),
				})
				if err != nil {
					log.Printf("WARN: failed to run container: %v", err)
					continue
				}

				run.Status = "started"
				run.ContainerID = containerInfo.Id
				result = db.Save(&run)
				if result.Error != nil {
					log.Printf("WARN: failed to update run status: %v", result.Error)
					continue
				}

				fmt.Printf("Started container id=%s", containerInfo.Id)
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
