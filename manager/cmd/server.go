package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/sithell/perun/manager/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	port int
)

func init() {
	flag.IntVar(&port, "port", 9001, "port for grpc server to listen on")
	flag.Parse()
}

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
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterManagerServer(grpcServer, &server{})
	log.Printf("Manager gRPC API serving at %s", lis.Addr().String())
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
