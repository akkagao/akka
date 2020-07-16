package route

import (
	"github.com/akkagao/akka/docker-go/handler"
	"github.com/gin-gonic/gin"
)

func RouterInit(engine *gin.Engine) {
	testRoot := engine.Group("/test")
	{
		testRoot.GET("/showJson", handler.ShowJson)
		testRoot.GET("/showJson1", handler.ShowJson)
	}
}
