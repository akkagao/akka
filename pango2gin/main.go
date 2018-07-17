package main

import (
	"net/http"

	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
	"gitlab.com/go-box/pongo2gin"
)

func main() {
	engine := gin.Default()

	// 设置模版解析器
	engine.HTMLRender = pongo2gin.Default()

	// 设置模版路径（使用默认模版解析器指定模版路径）
	// engine.LoadHTMLGlob("static/*.html")

	engine.GET("/", indexHandler)

	engine.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", pongo2.Context{"name": "crazywolf"})
	})

	engine.GET("/indextest", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", pongo2.Context{})
	})

	engine.Run(":8888")
}

func indexHandler(c *gin.Context) {
	c.HTML(200, "index.html", pongo2.Context{"name": "hahahha"})
}
