package main

import (
	"github.com/vaporz/turbo"
	"github.com/akkagao/turbo-demo/hello/gen"
	gcomponent "github.com/akkagao/turbo-demo/hello/grpcapi/component"
	gimpl "github.com/akkagao/turbo-demo/hello/grpcservice/impl"
	//tcomponent "github.com/akkagao/turbo-demo/hello/thriftapi/component"
	//timpl "github.com/akkagao/turbo-demo/hello/thriftservice/impl"
	"os/signal"
	"os"
	"syscall"
	"fmt"
)

func main() {
	s := turbo.NewGrpcServer(&gcomponent.ServiceInitializer{}, "/Users/crazywolf/go/gopath/src/github.com/akkagao/turbo-demo/hello/service.yaml")
	s.Start(gcomponent.GrpcClient, gen.GrpcSwitcher, gimpl.RegisterServer)

	//s := turbo.NewThriftServer(&tcomponent.ServiceInitializer{}, "/Users/crazywolf/go/gopath/src/github.com/akkagao/turbo-demo/hello/service.yaml")
	//s.Start(tcomponent.ThriftClient, gen.ThriftSwitcher, timpl.TProcessor)

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGQUIT)

	select {
	case <-exit:
		fmt.Println("Service is stopping...")
	}

	s.Stop()
	fmt.Println("Service stopped")
}
