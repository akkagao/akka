# 使用trace 观察系统go程序性能

参考文章：

https://mp.weixin.qq.com/s/nf_-AH_LeBN3913Pt6CzQQ



```go
package main

import (
	"fmt"
	"os"
	"runtime/trace"
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

	fmt.Println("=======以上为模版代码直接使用就好，下面随便写点东西作为测试使用========")

	count := 0
	for i := 0; i < 100; i++ {
		count += i
	}
	fmt.Println(count)
}

```

运行上面程序后会生成一个trace.out 文件然后使用 `go tool trace trace.out`就能生成自动打开浏览器查看信息了。

但是发现一个问题：启动GIN 作为web服务后 生成的trace.out 文件是空的。

代码如下：

```go
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
```

