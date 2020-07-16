package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowJson(c *gin.Context) {
	c.JSON(http.StatusOK, &struct {
		Name    string
		Address string
		Mobile  string
	}{
		Name:    "CrazyWolf",
		Address: "北京",
		Mobile:  "13913870540",
	})
}
