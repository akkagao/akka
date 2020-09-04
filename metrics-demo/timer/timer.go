package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/rcrowley/go-metrics"
)

func main() {
	t := metrics.GetOrRegisterTimer("account.create.latency", nil)
	// t.Update(1)

	go metrics.Log(metrics.DefaultRegistry,
		1*time.Second,
		log.New(os.Stdout, "metrics: ", log.Lmicroseconds))

	for true {

		t.Time(func() {
			rand.Seed(time.Now().UnixNano())
			n := rand.Intn(10)
			fmt.Println(n)
			time.Sleep(time.Second * time.Duration(n))
		})
	}
}
