package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/sithell/perun/backend/connector/internal"
	"github.com/sithell/perun/backend/connector/pb"
	"github.com/sithell/perun/backend/internal/database"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"gorm.io/gorm"
	"io"
	"log"
	"net"
	"os"
	"strconv"
)

var (
	dbHost     string
	dbUser     string
	dbPassword string
	dbPort     uint
	dbName     string
	port       int
)

func init() {
	flag.StringVar(&dbHost, "db-host", "localhost", "database host")
	flag.StringVar(&dbUser, "db-user", "perun", "database user")
	flag.UintVar(&dbPort, "db-port", 5432, "database port")
	flag.StringVar(&dbName, "db-name", "perun", "database name")
	dbPassword = os.Getenv("DATABASE_PASSWORD")
	flag.IntVar(&port, "port", 9002, "port for grpc server to listen on")
	flag.Parse()
}

type providerServer struct {
	pb.UnimplementedProviderServer
	db *gorm.DB
}

func (s *providerServer) InitConnection(srv pb.Provider_InitConnectionServer) error {
	ctx := srv.Context()
	p, _ := peer.FromContext(ctx)
	send := make(chan *pb.ServerRequest)
	receive := make(chan *pb.ClientResponse)
	provider := &internal.Provider{
		Ctx:     ctx,
		Send:    send,
		Receive: receive,
	}
	providerModel := &database.Provider{
		Host:   p.Addr.String(),
		Status: "active",
	}
	s.db.Create(providerModel)
	internal.AddProvider(strconv.Itoa(int(providerModel.ID)), provider)
	log.Printf("Opened connection with provider id=%d", providerModel.ID)

	ctx, cancel := context.WithCancel(ctx)
	go func(ctx context.Context, cancel context.CancelFunc) {
		for {
			request := <-send
			err := srv.Send(request)
			if err != nil {
				log.Printf("failed to send request: %v", err)
			}
		}
	}(ctx, cancel)
	go func(ctx context.Context, cancel context.CancelFunc) {
		for {
			req, err := srv.Recv()
			if err == io.EOF {
				log.Printf("Closed connection with provider id=%d", providerModel.ID)
				cancel()
				return
			}
			if err != nil {
				log.Printf("receive error %v", err)
				continue
			}
			provider.Receive <- req
		}
	}(ctx, cancel)
	<-ctx.Done()
	providerModel.Status = "disconnected"
	s.db.Save(providerModel)
	internal.DeleteProvider(strconv.Itoa(int(providerModel.ID)))
	return ctx.Err()
}

type apiServer struct {
	pb.UnimplementedApiServer
}

func (s *apiServer) Ping(context.Context, *pb.Empty) (*pb.PingResponse, error) {
	return &pb.PingResponse{}, nil
}

func (s *apiServer) RunContainer(_ context.Context, params *pb.RunContainerParams) (*pb.ContainerInfo, error) {
	provider, ok := internal.GetProviders()[strconv.FormatUint(params.ProviderId, 10)]
	if !ok {
		return nil, fmt.Errorf("no connection with provider_id=%d", params.ProviderId)
	}
	response, err := provider.RunContainer(&pb.RunContainerRequest{Image: params.Image, Cmd: params.Cmd})
	if err != nil {
		return nil, err
	}
	return &pb.ContainerInfo{Id: response.Id}, nil
}

func (s *apiServer) GetActiveConnections(context.Context, *pb.Empty) (*pb.ProviderConnections, error) {
	var result []*pb.ProviderConnection
	for id := range internal.GetProviders() {
		intId, err := strconv.Atoi(id)
		if err != nil {
			log.Printf("failed to convert string id to int: %v", err)
			continue
		}
		result = append(result, &pb.ProviderConnection{ProviderId: uint64(intId)})
	}
	return &pb.ProviderConnections{Connections: result}, nil
}

func main() {
	db, err := database.InitDB(dbHost, dbUser, dbPassword, dbName, dbPort)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterProviderServer(grpcServer, &providerServer{db: db})
	pb.RegisterApiServer(grpcServer, &apiServer{})
	log.Printf("Provider gRPC API serving at %s", lis.Addr().String())

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
