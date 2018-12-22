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
	msgs, err := ch.Consume("hello", "", true, false, false, false, nil)
	goutil.ChkErr(err)
	go func() {
		for msg := range msgs {
			fmt.Println(string(msg.Body))
		}
	}()
}
