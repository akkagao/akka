package main

import (
	"github.com/akkagao/akka/gin-demo/route"
	"github.com/gin-gonic/gin"
)

/**
启动gin服务
*/
func main() {
	r := gin.Default()
	//gin.SetMode(gin.ReleaseMode)
	route.RouterInit(r)
	r.Run(":8080")
}
