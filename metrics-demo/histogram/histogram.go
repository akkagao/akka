package main

import (
	"log"
	"os"
	"time"

	"github.com/rcrowley/go-metrics"
)

func main() {
	s := metrics.NewExpDecaySample(1028, 0.015) // or metrics.NewUniformSample(1028)
	h := metrics.NewHistogram(s)
	metrics.Register("baz", h)
	h.Update(1)

	go metrics.Log(metrics.DefaultRegistry,
		2*time.Second,
		log.New(os.Stdout, "metrics: ", log.Lmicroseconds))

	i := int64(0)
	for true {
		time.Sleep(time.Second * 1)
		i++
		h.Update(i)
	}
}
