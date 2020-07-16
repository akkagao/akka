package main

import (
	"github.com/akkagao/turbo-demo/hello/grpcservice/impl"
	"github.com/vaporz/turbo"
	"os/signal"
	"os"
	"syscall"
	"fmt"
)

func main() {
	s := turbo.NewGrpcServer(nil, "/Users/crazywolf/go/gopath/src/github.com/akkagao/turbo-demo/hello/service.yaml")
	s.StartGrpcService(impl.RegisterServer)

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGQUIT)

	select {
	case <-exit:
		fmt.Println("Service is stopping...")
	}

	s.Stop()
	fmt.Println("Service stopped")
}
