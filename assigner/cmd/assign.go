package main

import (
	"context"
	"fmt"
	"github.com/sithell/perun/internal/database"
	"github.com/sithell/perun/provider/pb"
	flag "github.com/spf13/pflag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"strings"
)

var (
	dbHost     string
	dbUser     string
	dbPassword string
	dbPort     uint
	dbName     string
)

func init() {
	flag.StringVar(&dbHost, "db-host", "localhost", "database host")
	flag.StringVar(&dbUser, "db-user", "perun", "database user")
	flag.UintVar(&dbPort, "db-port", 5432, "database port")
	flag.StringVar(&dbName, "db-name", "perun", "database name")
	dbPassword = os.Getenv("DATABASE_PASSWORD")
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

	var jobs []database.Job
	result = db.Find(&jobs)
	if result.Error != nil {
		log.Fatalf("failed to fetch jobs from db: %v", result.Error)
	}
	for i, job := range jobs {
		fmt.Printf("Job(id=%d, image=%s, command=%s, created_at=%s)\n", job.ID, job.Image, job.Command, job.CreatedAt)
		var runs []database.Run
		result = db.Where("runs.job_id = ?", job.ID).Find(&runs)
		if result.Error != nil {
			log.Printf("WARN: failed to fetch runs from db: %v", result.Error)
			continue
		}
		if len(runs) == 0 {
			// job was never run before
			assignedProvider := providerClients[i%len(providerClients)]
			run := database.Run{
				Job:      job,
				Provider: assignedProvider.Provider,
				Status:   "started",
			}
			result := db.Save(&run)
			if result.Error != nil {
				log.Printf("WARN: failed to save run in db: %v", result.Error)
				continue
			}
			containerInfo, err := assignedProvider.Client.RunContainer(ctx, &pb.RunContainerParams{
				Image: job.Image,
				Cmd:   strings.Split(job.Command, " "),
			})
			if err != nil {
				log.Printf("WARN: failed to run container: %v", err)
				continue
			}
			fmt.Printf("Started container id=%s", containerInfo.Id)
		}
	}
}
