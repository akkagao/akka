package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/rcrowley/go-metrics"
)

func main() {
	fmt.Println(fmt.Sprintf("%.2f", 99999.99))
	c := metrics.NewCounter()
	metrics.Register("foo", c)
	c.Inc(45)

	go metrics.Log(metrics.DefaultRegistry,
		1*time.Second,
		log.New(os.Stdout, "metrics: ", log.Lmicroseconds))

	for true {
		time.Sleep(time.Second * 1)
		c.Inc(1)
	}
}
