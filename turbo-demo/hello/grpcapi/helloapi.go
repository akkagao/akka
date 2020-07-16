package main

import (
	"github.com/akkagao/turbo-demo/hello/gen"
	"github.com/akkagao/turbo-demo/hello/grpcapi/component"
	"github.com/vaporz/turbo"
	"os/signal"
	"os"
	"syscall"
	"fmt"
)

func main() {
	s := turbo.NewGrpcServer(&component.ServiceInitializer{}, "/Users/crazywolf/go/gopath/src/github.com/akkagao/turbo-demo/hello/service.yaml")
	s.StartHTTPServer(component.GrpcClient, gen.GrpcSwitcher)

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGQUIT)

	select {
	case <-exit:
		fmt.Println("Service is stopping...")
	}

	s.Stop()
	fmt.Println("Service stopped")
}
