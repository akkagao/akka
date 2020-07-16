package main

// server.go

import (
	"fmt"
	"log"
	"net"

	"github.com/akkagao/akka/grpc-demo/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *proto.HelloRequest) (*proto.HelloResponse, error) {
	fmt.Println(in.GetGreeting())
	return &proto.HelloResponse{Reply: "Hello " + in.Greeting}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterHelloServiceServer(s, &server{})
	s.Serve(lis)
}
