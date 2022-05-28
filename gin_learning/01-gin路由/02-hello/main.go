// 练习一下基本路由

package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// 创建 默认路由
	r := gin.Default()

	// 绑定路由规则，定义执行的函数
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// 监听端口默认为 8080, 0.0.0.0:8080
	r.Run()
}
