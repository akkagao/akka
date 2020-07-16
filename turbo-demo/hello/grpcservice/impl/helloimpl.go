package impl

import (
	"github.com/akkagao/turbo-demo/hello/gen/proto"
	"golang.org/x/net/context"
)

// Hello is the struct which implements generated interface
type Hello struct {
}

// SayHello is an example entry point
func (s *Hello) SayHello(ctx context.Context, req *proto.SayHelloRequest) (*proto.SayHelloResponse, error) {
	return &proto.SayHelloResponse{Message: "[grpc server]Hello, " + req.YourName}, nil
}
