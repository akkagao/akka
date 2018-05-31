package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
测试获取restFul 参数
http://localhost:8080/test/getuser/1345
*/
// @Summary test
// @Param id path int true "id"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /test/getuser/{id} [get]
func GetUserById(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, fmt.Sprintf("id :%s user's name is akkagao...", id))
}

/**
获取？后面的参数
http://localhost:8080/test/queryStrParameter?firstname=123&lastname=ggg
DefaultQuery 方法如果获取不到参数则返回默认值，如下地址则 firstName 的值将为 akka
http://localhost:8080/test/queryStrParameter?lastname=ggg
*/
func QueryStrParameter(c *gin.Context) {
	firstName := c.DefaultQuery("firstname", "akka")
	lastName := c.Query("lastname")
	c.JSON(http.StatusOK, fmt.Sprintf("firstName: %s, lastName: %s", firstName, lastName))
}

/**
获取form表单提交的参数
用PostMan 测试的时候不能参数不能放在 params 中，应该放在body中作为参数传过来 否则无法接受到参数

// todo PostForm 这种参数swagger无法传参数
*/
// @Description 获取form表单提交的参数
// @Param name query string true "名称"
// @Param password query string true "密码"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /test/form_post [post]
func FormPostParameter(c *gin.Context) {
	// 如果没有任何获取其他参数的方法，Default方法也不能获取默认参数，
	// 至少有一个获取其他参数的方法，没有name 参数的时候才能设置默认值
	//name := c.DefaultPostForm("name", "akka")
	//password := c.PostForm("password")
	nameQuery := c.Query("name")
	passwordQuery := c.Query("password")
	//c.JSON(http.StatusOK, fmt.Sprintf("user %s login success, it password is : %s", name, password))
	c.JSON(http.StatusOK, fmt.Sprintf("user %s login success, it password is : %s", nameQuery, passwordQuery))
}

/**
上传文件测试
*/
func Uploadfile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		fmt.Errorf("upload file error, msg: %s", err)
		c.JSON(http.StatusInternalServerError, "system error,place retry later...")
		return
	}
	c.JSON(http.StatusOK, fmt.Sprintf("upload file name is %s", file.Filename))
}

type UserLogin struct {
	Name    string `json:"name" form:"name" binding:"required"`
	Address string `json:"address" form:"address" binding:"required"`
}

/**
JSON 请求参数，用ShouldBindJSON方法bing到struct对象
*/
func JsonParameter(c *gin.Context) {
	var userLogin UserLogin
	if err := c.ShouldBindJSON(&userLogin); err == nil {
		if userLogin.Name == "akka" {
			fmt.Println("user %s's address is %s", userLogin.Name, userLogin.Address)
			c.JSON(http.StatusOK, fmt.Sprintf("%s login success", userLogin.Name))
		} else {
			c.JSON(http.StatusUnauthorized, "unauthorized")
		}
	} else {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("bad request, err: ", err.Error()))
	}
}

/**
form 表单请求参数，用ShouldBind方法bind到struct对象
*/
func BingFormParameter(c *gin.Context) {
	var userLogin UserLogin
	if err := c.ShouldBind(&userLogin); err == nil {
		if userLogin.Name == "akka" {
			fmt.Println("user %s's address is %s", userLogin.Name, userLogin.Address)
			c.JSON(http.StatusOK, fmt.Sprintf("%s login success", userLogin.Name))
		} else {
			c.JSON(http.StatusUnauthorized, "unauthorized")
		}
	} else {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("bad request, err: ", err.Error()))
	}
}
