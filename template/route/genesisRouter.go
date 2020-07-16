package route

import (
	"github.com/akkagao/akka/template/handler"
	"github.com/gin-gonic/gin"
)

/**
* 初始基准测试
 */
func genesisRouterInit(engine *gin.Engine) {
	genesis := engine.Group("/genesis")
	{
		genesis.GET("/health", handler.Health)
	}
}
