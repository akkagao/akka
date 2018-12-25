package main

import (
	"strconv"

	"github.com/akkagao/goutil"
	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	goutil.ChkErr(err)
	defer conn.Close()

	ch, err := conn.Channel()
	goutil.ChkErr(err)
	q, err := ch.QueueDeclare("work", false, false, false, false, nil)
	goutil.ChkErr(err)
	body := "咱们工人有力量"

	for i := 0; i < 50; i++ {
		err = ch.Publish("", q.Name, false, false, amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(body + strconv.Itoa(i)),
		})

		goutil.ChkErr(err)
	}

}
