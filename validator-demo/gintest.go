package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

//validate 设置方式 https://godoc.org/gopkg.in/go-playground/validator.v10

type User struct {
	ID int64 `form:"id" json:"id" validate:"gte=0,lte=130"`
}

func main() {
	validate := validator.New()
	engin := gin.Default()
	engin.GET("/getuser/:id", func(c *gin.Context) {
		user := &User{}
		c.Bind(user)
		idStr := c.Param("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		fmt.Println(err)
		user.ID = id
		fmt.Println(user)
		err = validate.Struct(user)
		if err != nil {
			fmt.Println(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"1": 1,
		})
	})
	engin.Run(":8080")
}
