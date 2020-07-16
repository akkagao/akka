package route

import "github.com/gin-gonic/gin"

func RouterInit(engine *gin.Engine) {
	genesisRouterInit(engine)
}
