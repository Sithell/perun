package main

import (
	"context"
	"fmt"
	"github.com/sithell/perun/manager/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedManagerServer
}

func (s *server) RegisterProvider(context.Context, *pb.ProviderInfo) (*pb.RegisterProviderResponse, error) {
	fmt.Println("got a request to register a new provider")
	return &pb.RegisterProviderResponse{
		Ok:      true,
		Message: "this method is not yet implemented",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterManagerServer(grpcServer, &server{})
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
