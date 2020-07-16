package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/garyburd/redigo/redis"
)

func main() {
	log.Println("test redis")
	c, err := redis.Dial("tcp", "127.0.0.1:6379")

	if err != nil {
		os.Exit(0)
	}
	defer c.Close()
	c.Do("set", "aa", 3)
	result, err := redis.Int(c.Do("get", "aa"))
	if err != nil {
		fmt.Println("get result from redis error")
	}
	fmt.Println("result:", result)

	go subMsg()
	channel := make(chan struct{})
	go pubMsg(channel, c)
	<-channel

}
func subMsg() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		os.Exit(0)
	}
	defer c.Close()
	sub := redis.PubSubConn{Conn: c}
	sub.Subscribe("channel_test")
	for {
		switch v := sub.Receive().(type) {
		case redis.Message:
			fmt.Printf("%s\n", v.Data)
		case error:
			fmt.Println("----", v)
		}
	}
}

func pubMsg(c chan struct{}, conn redis.Conn) {
	result, err := redis.Int(conn.Do("get", "aa"))
	fmt.Println("result1:", result, err)

	for i := 0; i < 10; i++ {
		conn.Do("PUBLISH", "channel_test", time.Now())
		time.Sleep(time.Second)
	}
	c <- struct{}{}
}
