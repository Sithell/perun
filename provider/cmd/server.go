package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/sithell/perun/provider/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	port int
)

func init() {
	flag.IntVar(&port, "port", 9002, "port for grpc server to listen on")
	flag.Parse()
}

type server struct {
	pb.UnimplementedProviderServer
}

func (s *server) RunContainer(ctx context.Context, runContainerParams *pb.RunContainerParams) (*pb.ContainerInfo, error) {
	fmt.Printf("got a request to run a container: %v", runContainerParams)
	return &pb.ContainerInfo{Id: ""}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
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
