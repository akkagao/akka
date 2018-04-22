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
