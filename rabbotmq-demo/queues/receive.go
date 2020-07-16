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
	q, err := ch.QueueDeclare("hello", false, false, false, false, nil)

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	goutil.ChkErr(err)
	go func() {
		for msg := range msgs {
			fmt.Println(string(msg.Body))
		}
	}()
	select {}
}
