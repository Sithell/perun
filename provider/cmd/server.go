package main

import (
	"context"
	"flag"
	"fmt"
	connector "github.com/sithell/perun/connector/pb"
	"github.com/sithell/perun/provider/internal/docker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"strings"
)

var (
	connectorHost string
	connectorPort uint
)

func init() {
	flag.StringVar(&connectorHost, "connector-host", "localhost", "connector host")
	flag.UintVar(&connectorPort, "connector-port", 9002, "connector port")
	flag.Parse()
}

func main() {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", connectorHost, connectorPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to connector: %v", err)
	}
	client := connector.NewProviderClient(conn)
	ctx := context.Background()
	connection, err := client.InitConnection(ctx)
	if err != nil {
		log.Fatalf("failed to init connection: %v", err)
	}
	fmt.Printf("Waiting for commands from connector at %s:%d", connectorHost, connectorPort)
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			break
		default:
		}
		request, err := connection.Recv()
		if err != nil {
			log.Fatalf("failed to recieve request: %v", err)
		}
		switch request.Body.(type) {
		case *connector.ServerRequest_RunContainer:
			runContainerRequest := request.GetRunContainer()
			containerID, err := docker.RunContainer(runContainerRequest.Image, strings.Split(runContainerRequest.Cmd, " "))
			if err != nil {
				log.Fatalf("failed to run docker container: %v", err)
			}
			err = connection.Send(&connector.ClientResponse{
				ResponseTo: request.Id,
				Body: &connector.ClientResponse_RunContainer{
					RunContainer: &connector.RunContainerResponse{Id: containerID},
				},
			})
			if err != nil {
				log.Fatalf("got error from connector: %v", err)
			}
		}
	}
}
