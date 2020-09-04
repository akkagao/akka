package main

import (
	"log"
	"os"
	"time"

	"github.com/rcrowley/go-metrics"
)

func main() {
	g := metrics.NewGauge()
	metrics.Register("bar", g)
	g.Update(1)

	go metrics.Log(metrics.DefaultRegistry,
		1*time.Second,
		log.New(os.Stdout, "metrics: ", log.Lmicroseconds))

	j := int64(1)

	for true {
		time.Sleep(time.Second * 1)
		g.Update(j)
		j++
	}
}
