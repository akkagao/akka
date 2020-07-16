package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("test", func(c *gin.Context) {
		c.JSON(200, "hello world...")
	})
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "index page")

	})
	r.Run(":8080")
}
