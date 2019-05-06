package main

import (
	"log"

	"github.com/gomodule/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "172.16.18.31:6378")
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer c.Close()
	if _, err := c.Do("AUTH", "library20171012PRD"); err != nil {
		log.Fatalln(err)
		return
	}

	if _, err := c.Do("SET", "aa", "aa"); err != nil {
		log.Println(err)
		return
	}
	if result, err := redis.String(c.Do("GET", "aa")); err != nil {
		log.Println(err)
		return
	} else {
		log.Println(result)
	}

	scriptStr := `return redis.call('set',KEYS[1],ARGV[1])`
	script := redis.NewScript(1, scriptStr)
	r, err := script.Do(c, "aa", "123")
	log.Println(r, err)

	// 获取 redis 中 aa 的值
	if result, err := redis.String(c.Do("GET", "aa")); err != nil {
		log.Println(err)
		return
	} else {
		log.Println(result)
	}

}
