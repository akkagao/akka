package main

// server.go

import (
	"fmt"
	"log"
	"net"
	"strings"
	"sync"

	"github.com/akkagao/akka/grpc-demo/proto"
	"github.com/hashicorp/consul/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
)

const (
	port = 50053
)

var (
	service_locker  = new(sync.Mutex)
	consul_client   *api.Client
	my_service_id   string
	my_service_name string
	my_kv_key       string
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *proto.HelloRequest) (*proto.HelloResponse, error) {
	fmt.Println(in.GetGreeting())
	return &proto.HelloResponse{Reply: "Hello22 " + in.Greeting}, nil
}

// HealthImpl 健康检查实现
type HealthImpl struct{}

// Check 实现健康检查接口，这里直接返回健康状态，这里也可以有更复杂的健康检查策略，比如根据服务器负载来返回
func (h *HealthImpl) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	return &grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	}, nil
}

func (h *HealthImpl) Watch(req *grpc_health_v1.HealthCheckRequest, watchServer grpc_health_v1.Health_WatchServer) error {
	return nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	grpc_health_v1.RegisterHealthServer(s, &HealthImpl{})
	proto.RegisterHelloServiceServer(s, &server{})
	DoRegistService("helloService", "127.0.0.1", port)
	s.Serve(lis)
}

func DoRegistService(service_name string, ip string, port int) {
	my_service_id := fmt.Sprintf("%s-%d", service_name, port)
	var tags []string
	service := &api.AgentServiceRegistration{
		ID:      my_service_id,
		Name:    service_name,
		Port:    port,
		Address: ip,
		Tags:    tags,
		Check: &api.AgentServiceCheck{
			// HTTP:     "http://" + monitor_addr + "/status",
			GRPC:     fmt.Sprintf("%v:%v/%v", ip, port, service_name), // grpc 支持，执行健康检查的地址，service 会传到 Health.Check 函数中
			Interval: "5s",
			Timeout:  "1s",
		},
	}

	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		log.Fatal(err)
	}
	consul_client = client
	if err := consul_client.Agent().ServiceRegister(service); err != nil {
		log.Fatal(err)
	}
	log.Printf("Registered service %q in consul with tags %q", service_name, strings.Join(tags, ","))
}
