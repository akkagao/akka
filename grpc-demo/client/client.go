package main

//client.go

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	pb "github.com/akkagao/akka/grpc-demo/proto"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:80"
	defaultName = "world"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewHelloServiceClient(conn)
	rand.Seed(time.Now().UnixNano())
	name := fmt.Sprintf("%s%d", defaultName, rand.Int())
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Greeting: name})
	if err != nil {
		log.Fatal("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Reply)
}
