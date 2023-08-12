package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/sithell/perun/connector/internal"
	"github.com/sithell/perun/connector/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"io"
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

type providerServer struct {
	pb.UnimplementedProviderServer
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
	providerId := p.Addr.String()
	internal.AddProvider(providerId, provider)
	log.Printf("Opened connection with provider %s", providerId)

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
				log.Printf("Closed connection with provider %s", providerId)
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
	internal.DeleteProvider(providerId)
	return ctx.Err()
}

type apiServer struct {
	pb.UnimplementedApiServer
}

func (s *apiServer) Ping(context.Context, *pb.Empty) (*pb.PingResponse, error) {
	return &pb.PingResponse{}, nil
}

func (s *apiServer) RunContainer(_ context.Context, params *pb.RunContainerParams) (*pb.ContainerInfo, error) {
	var provider *internal.Provider
	for _, v := range internal.GetProviders() {
		provider = v
		break
	}
	if provider == nil {
		return nil, fmt.Errorf("no providers available at the moment")
	}
	response, err := provider.RunContainer(&pb.RunContainerRequest{Image: params.Image, Cmd: params.Cmd})
	if err != nil {
		return nil, err
	}
	return &pb.ContainerInfo{Id: response.Id}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterProviderServer(grpcServer, &providerServer{})
	pb.RegisterApiServer(grpcServer, &apiServer{})
	log.Printf("Provider gRPC API serving at %s", lis.Addr().String())

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
