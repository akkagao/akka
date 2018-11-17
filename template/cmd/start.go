package cmd

import (
	"github.com/akkagao/akka/template/common"
	"github.com/akkagao/akka/template/route"
	"github.com/gin-gonic/gin"
)

func start(configName string) {
	common.InitConfig(configName)
	common.StartDb()
	common.StartRedis()
	common.StartService()
	startWeb()
}

/**
* 启动web服务
 */
func startWeb() {
	engine := gin.Default()
	engine.Use(gin.Logger())
	// 设置图标
	// engine.Use(favicon.New("./favicon.ico"))
	// engine.StaticFile("/index", "./page/index.html")

	route.RouterInit(engine)
	port := common.Config.GetString("sys.port")
	engine.Run(":" + port)
}
