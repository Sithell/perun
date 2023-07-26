package main

import (
	"context"
	"fmt"
	"github.com/sithell/perun/provider/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedProviderServer
}

func (s *server) RunContainer(ctx context.Context, runContainerParams *pb.RunContainerParams) (*pb.ContainerInfo, error) {
	fmt.Printf("got a request to run a container: %v", runContainerParams)
	return &pb.ContainerInfo{Id: ""}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterProviderServer(grpcServer, &server{})
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
