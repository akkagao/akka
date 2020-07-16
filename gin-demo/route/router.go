package route

import (
	"log"
	"net/http"

	"github.com/akkagao/akka/gin-demo/handler"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func TimeCase() gin.HandlerFunc {
	return func(context *gin.Context) {
		log.Println("before into index page")
		context.Next()
		log.Println("after index page")
	}
}

/**
初始化路由
*/
func RouterInit(engine *gin.Engine) {
	engine.GET("/", TimeCase(), func(c *gin.Context) {
		c.JSON(http.StatusOK, "index page")
	})

	// 增加swagger
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 设置路由组
	testGroup := engine.Group("/test")
	{
		/**
		获取restFul 参数
		curl GET http://localhost:8080/test/getuser/1345
		*/
		testGroup.GET("/getuser/:id", handler.GetUserById)

		/**
		获取？后面的参数
		curl GET http://localhost:8080/test/queryStrParameter?firstname=123&lastname=ggg
		*/
		testGroup.GET("/queryStrParameter", handler.QueryStrParameter)

		/**
		获取form表单提交参数
		curl -X POST http://localhost:8080/test/form_post -F name=gggg -F password=aaa
		*/
		testGroup.POST("/form_post", handler.FormPostParameter)

		/**
		上传文件
		  curl -X POST http://localhost:8080/test/uploadfile \
		  -F "file=@/Users/crazywolf/myapp/canal/canal.deployer-1.0.25.tar.gz" \
		  -H "Content-Type: multipart/form-data"
		*/
		testGroup.POST("/uploadfile", handler.Uploadfile)

		/**
		JSON 请求参数，用ShouldBindJSON方法bing到struct对象
		curl -X POST http://localhost:8080/test/jsonParameter \
		-H 'Content-Type: application/json' \
		-d '{
			"name":"akka",
			"address":"lanzhou"
		}'
		*/
		testGroup.POST("/jsonParameter", handler.JsonParameter)
		/**
		form 表单请求参数，用ShouldBind方法bind到struct对象
		curl -X POST \
		  http://localhost:8080/test/bingFormParameter \
		  -F name=akka \
		  -F address=lanzhou
		*/
		testGroup.POST("/bingFormParameter", handler.BingFormParameter)

		// 生成二维码
		testGroup.GET("/genQrCode", handler.GenQrCode)
	}

}
