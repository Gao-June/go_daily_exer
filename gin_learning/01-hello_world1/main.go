// 练习使用 Gin 框架
// 访问该 localhost: 在终端 curl http://localhost:8090/hello

package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建 服务器 引擎
	engine := gin.Default()

	engine.GET("/hello", func(context *gin.Context) {
		fmt.Println("请求路径", context.FullPath())
		context.Writer.Write([]byte("Hello, gin\n"))
	})

	// Run 可以修改端口
	if err := engine.Run(":8090"); err != nil {

		log.Fatal(err.Error())
	}

}
