package main

import (
	"fmt"
	"log"

	"github.com/guonaihong/gout"
)

func main() {
	g := gout.New(nil)
	var httpCode int
	var result interface{}

	if err := g.GET("https://c-p.dubaner.com/api/v1/accesstoken/parents").Code(&httpCode).BindJSON(&result).Do(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(httpCode)
	fmt.Println(result)
}
