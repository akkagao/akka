package impl

import (
	"fmt"

	"github.com/akkagao/turbo-demo/hello/gen/proto"
	"google.golang.org/grpc"
)

// RegisterServer registers a service struct to a server
func RegisterServer(s *grpc.Server) {
	fmt.Println("1")
	proto.RegisterHelloServer(s, &Hello{})
	fmt.Println("2")
	proto.RegisterWorldServer(s, &World{})
	fmt.Println("3")
}
