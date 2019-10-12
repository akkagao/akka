package main

import (
	"context"
	"log"
	"time"

	proto "github.com/micro/examples/helloworld/proto"
	"github.com/micro/go-micro"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.hello"),
		micro.RegisterTTL(time.Second*30),
	)

	service.Init()
	proto.RegisterGreeterHandler(service.Server(), new(Greeter))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

// protoc --go_out=plugins=grpc:. --micro_out==plugins=grpc:. greeter.proto
