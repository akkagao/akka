package main

import (
	"github.com/akkagao/akka/gin-demo/route"
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
)

/**
启动gin服务
*/
func main() {
	engine := gin.Default()
	//gin.SetMode(gin.ReleaseMode)
	engine.Use(favicon.New("./favicon.ico"))
	route.RouterInit(engine)
	engine.Run(":8080")
}
