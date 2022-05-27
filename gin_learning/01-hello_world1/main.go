// 练习使用 Gin 框架
// 访问该 localhost: 在终端 curl http://localhost:8090/hello

package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建 默认路由
	engine := gin.Default()

	// 绑定路由规则，执行的函数
	// gin.Context 封装了 request 和 response
	engine.GET("/hello", func(context *gin.Context) {
		fmt.Println("请求路径", context.FullPath())
		context.Writer.Write([]byte("Hello, gin\n"))
	})

	// 监听端口，默认是在 8080
	// Run 可以修改端口
	if err := engine.Run(":8090"); err != nil {
		log.Fatal(err.Error())
	}

}
