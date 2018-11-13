package main

import (
	"fmt"
	"log"
	"os"

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
}
