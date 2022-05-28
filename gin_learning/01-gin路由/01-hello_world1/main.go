// 练习使用 Gin 框架
// 访问该 localhost: 在终端 curl http://localhost:8090/hello

package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func this_fun(context *gin.Context) {
	fmt.Println("请求路径", context.FullPath())
	context.Writer.Write([]byte("Hello, gin\n"))
}

func main() {
	// 创建 默认路由
	engine := gin.Default()

	// 绑定路由规则，定义执行的函数
	// gin.Context 封装了 request 和 response
	/* 这里的 func(){...} 类似lambda，其实可以再写一个函数，然后把函数名放在这里。

	engine.GET("/hello", func(context *gin.Context) {
		fmt.Println("请求路径", context.FullPath())
		context.Writer.Write([]byte("Hello, gin\n"))
	})

	*/
	engine.GET("/hello", this_fun)

	// 监听端口，默认是在 8080
	// Run 可以修改端口
	if err := engine.Run(":8090"); err != nil {
		log.Fatal(err.Error())
	}

}
