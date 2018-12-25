package main

import (
	"fmt"

	"github.com/akkagao/goutil"
	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	goutil.ChkErr(err)
	defer conn.Close()
	ch, err := conn.Channel()
	goutil.ChkErr(err)
	msgs, err := ch.Consume("work", "", true, false, false, false, nil)
	goutil.ChkErr(err)
	fc := make(chan struct{})
	go func() {
		for msg := range msgs {
			fmt.Printf("%s\n", msg.Body)
		}
	}()
	<-fc

}
