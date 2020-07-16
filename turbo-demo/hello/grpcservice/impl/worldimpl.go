package impl

import (
	"github.com/akkagao/turbo-demo/hello/gen/proto"
	"golang.org/x/net/context"
)

// Hello is the struct which implements generated interface
type World struct {
}

// SayWorld is an example entry point
func (s *World) SayWorld(ctx context.Context, req *proto.SayWorldRequest) (*proto.SayWorldResponse, error) {
	return &proto.SayWorldResponse{Message: "[grpc server]Hello, " + req.YourWold + "-" + req.Password}, nil
}
