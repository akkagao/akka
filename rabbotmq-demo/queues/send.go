package main

import (
	"github.com/akkagao/goutil"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	goutil.ChkErr(err)
	defer conn.Close()
	ch, err := conn.Channel()
	goutil.ChkErr(err)
	q, err := ch.QueueDeclare("hello", false, false, false, false, nil)
	body := "world"
	ch.Publish("", q.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(body),
	})
}
