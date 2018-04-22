package route

import (
	"github.com/akkagao/akka/gin-demo/handler"
	"github.com/gin-gonic/gin"
)

/**
初始化路由
*/
func RouterInit(r *gin.Engine) {
	testGroup := r.Group("/test")
	/**
	获取restFul 参数
	http://localhost:8080/test/getuser/1345
	*/
	testGroup.GET("/getuser/:id", handler.GetUserById)
	/**
	获取？后面的参数
	http://localhost:8080/test/queryStrParameter?firstname=123&lastname=ggg
	*/
	testGroup.GET("/queryStrParameter", handler.QueryStrParameter)

}
