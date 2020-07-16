package handlers

import (
	"context"

	pb "github.com/akka/gokit-demo/hello"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.HelloServiceServer {
	return helloserviceService{}
}

type helloserviceService struct{}

// SayHello implements Service.
func (s helloserviceService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	var resp pb.HelloResponse
	resp = pb.HelloResponse{
		// Reply:
		// Number:
	}
	return &resp, nil
}
