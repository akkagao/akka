package main

import (
	"github.com/gin-gonic/gin"
	"github.com/akkagao/akka/gin-demo/route"
)

func main() {
	r := gin.Default()
	route.RouterInit(r)


	r.Run(":8080")
}
