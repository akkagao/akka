package main

//client.go

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"

	pb "github.com/akkagao/akka/grpc-demo/proto"
	"github.com/hashicorp/consul/api"
	"github.com/sirupsen/logrus"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/naming"
)

const (
	defaultName = "world"
)

func main() {
	conn, err := grpc.Dial(
		"",
		grpc.WithInsecure(),
		// 负载均衡，使用 consul 作服务发现
		grpc.WithBalancer(grpc.RoundRobin(NewConsulResolver(
			"127.0.0.1:8500", "helloService",
		))),
	)

	if err != nil {
		log.Fatal("did not connect: %v", err)
	}
	defer conn.Close()

	var serviceStart = make(chan struct{})
	// 如果不等待conn的State变化就调用远程方法，第一次调用可能会失败，因为获取服务列表是异步的
	go func() {
		for {
			if conn.GetState() == connectivity.Ready {
				serviceStart <- struct{}{}
			}
		}
	}()
	<-serviceStart

	c := pb.NewHelloServiceClient(conn)
	rand.Seed(time.Now().UnixNano())
	name := fmt.Sprintf("%s%d", defaultName, rand.Int())
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	for {
		r, err := c.SayHello(context.Background(), &pb.HelloRequest{Greeting: name})
		if err != nil {
			log.Fatal("could not greet: %v", err)
		}
		log.Printf("Greeting: %s", r.Reply)
		time.Sleep(1 * time.Second)
	}

}

// NewConsulResolver consul resolver
func NewConsulResolver(address string, service string) naming.Resolver {
	return &consulResolver{
		address: address,
		service: service,
	}
}

type consulResolver struct {
	address string
	service string
}

// Resolve implement
func (r *consulResolver) Resolve(target string) (naming.Watcher, error) {
	fmt.Println("=====Resolve=======", target)
	config := api.DefaultConfig()
	config.Address = r.address
	client, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}

	return &consulWatcher{
		client:  client,
		service: r.service,
		addrs:   map[string]struct{}{},
	}, nil
}

type consulWatcher struct {
	client    *api.Client
	service   string
	addrs     map[string]struct{}
	lastIndex uint64
}

func (w *consulWatcher) Next() ([]*naming.Update, error) {

	for {
		services, metainfo, err := w.client.Health().Service(w.service, "", true, &api.QueryOptions{
			WaitIndex: w.lastIndex, // 同步点，这个调用将一直阻塞，直到有新的更新
		})

		if err != nil {
			logrus.Warn("error retrieving instances from Consul: %v", err)
		}
		w.lastIndex = metainfo.LastIndex

		addrs := map[string]struct{}{}
		for _, service := range services {
			addrs[net.JoinHostPort(service.Service.Address, strconv.Itoa(service.Service.Port))] = struct{}{}
		}

		var updates []*naming.Update
		for addr := range w.addrs {
			if _, ok := addrs[addr]; !ok {
				updates = append(updates, &naming.Update{Op: naming.Delete, Addr: addr})
			}
		}

		for addr := range addrs {
			if _, ok := w.addrs[addr]; !ok {
				updates = append(updates, &naming.Update{Op: naming.Add, Addr: addr})
			}
		}

		if len(updates) != 0 {
			w.addrs = addrs
			return updates, nil
		}
	}
}

func (w *consulWatcher) Close() {
	// nothing to do
}
