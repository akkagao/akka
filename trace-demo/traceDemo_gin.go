package main

import (
	"fmt"
	"os"
	"runtime/trace"

	"github.com/gin-gonic/gin"
)

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()
	fmt.Println("go tool trace")

	count := 0
	for i := 0; i < 100; i++ {
		count = count + i
	}
	fmt.Println(count)

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "success")
	})
	r.Run(":8080")

}
