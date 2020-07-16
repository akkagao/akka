package handler

import (
	"net/http"

	"github.com/akkagao/akka/template/common"
	"github.com/gin-gonic/gin"
)

/**
* 健康检查
 */
// curl 'http://localhost:8088/genesis/health'
func Health(c *gin.Context) {
	genesis := common.GenesisService.GetById(1)
	c.JSON(http.StatusOK, gin.H{
		"web":     "ok",
		"env":     common.Config.Get("env"),
		"genesis": genesis,
	})
}
