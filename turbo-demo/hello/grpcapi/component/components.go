package component

import (
	"fmt"

	"github.com/akkagao/turbo-demo/hello/gen/proto"
	"github.com/vaporz/turbo"
	"google.golang.org/grpc"
)

// GrpcClient returns a grpc client
func GrpcClient(conn *grpc.ClientConn) interface{} {
	return proto.NewHelloClient(conn)
}

type ServiceInitializer struct {
}

// InitService is run before the service is started, do initializing staffs for your service here.
// For example, init turbo components, such as interceptors, pre/postprocessors, errorHandlers, etc.
func (i *ServiceInitializer) InitService(s turbo.Servable) error {
	// TODO
	fmt.Println("InitService")
	return nil
}

// StopService is run after both grpc server and http server are stopped,
// do your cleaning up work here.
func (i *ServiceInitializer) StopService(s turbo.Servable) {
	// TODO
	fmt.Println("StopService")
}
