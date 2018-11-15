package main

import (
	"github.com/akkagao/akka/docker-go/route"
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
)

/**
启动gin服务
*/
func main() {
	engine := gin.Default()
	gin.SetMode(gin.ReleaseMode)

	engine.Use(gin.Logger())
	// 设置图标
	engine.Use(favicon.New("./favicon.ico"))
	engine.StaticFile("/index", "./page/index.html")

	route.RouterInit(engine)
	engine.Run(":8088")
}
