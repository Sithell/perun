package internal

import (
	"context"
	"fmt"
	"github.com/sithell/perun/backend/connector/pb"
	"log"
	"math/rand"
	"strconv"
)

type Provider struct {
	Ctx     context.Context
	Send    chan *pb.ServerRequest
	Receive chan *pb.ClientResponse
}

var providers map[string]*Provider

// AddProvider returns false if provider with given id already exists and true otherwise
func AddProvider(id string, provider *Provider) bool {
	if _, ok := providers[id]; ok {
		return false
	}
	providers[id] = provider
	return true
}

// DeleteProvider returns false if provider with given id doesn't exist and true otherwise
func DeleteProvider(id string) bool {
	if _, ok := providers[id]; !ok {
		return false
	}
	delete(providers, id)
	return true
}

func GetProviders() map[string]*Provider {
	return providers
}

func init() {
	providers = make(map[string]*Provider)
}

func (p *Provider) makeRequest(request *pb.ServerRequest) (*pb.ClientResponse, error) {
	if request.Id == "" {
		request.Id = strconv.FormatInt(rand.Int63n(256*256*256), 16)
		log.Printf("request id = %s", request.Id)
	}
	p.Send <- request
	for {
		select {
		case <-p.Ctx.Done():
			log.Printf("makeRequest: %v", p.Ctx.Err())
			return nil, p.Ctx.Err()
		case response := <-p.Receive:
			log.Printf("response_to=%s, request_id=%s", response.ResponseTo, request.Id)
			if response.ResponseTo == request.Id {
				return response, nil
			}
			return nil, fmt.Errorf("client returned response with mismatching request_id")
		default:
		}
	}
}

func (p *Provider) RunContainer(params *pb.RunContainerRequest) (*pb.RunContainerResponse, error) {
	response, err := p.makeRequest(&pb.ServerRequest{
		Body: &pb.ServerRequest_RunContainer{
			RunContainer: params,
		},
	})
	if err != nil {
		return nil, err
	}
	log.Printf("response is: '%s', body is: '%s'", response.String(), response.Body)
	return response.GetRunContainer(), nil
}
