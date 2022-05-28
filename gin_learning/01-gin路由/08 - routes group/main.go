/*
	routes group 管理一些相同的 URL

	其实就是做了一些简写， eg 04-URL参数那节，是一样的。

	【结】 POST处的代码 无法正确运行，不知道为啥
*/

package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func login(c *gin.Context) {
	name := c.DefaultQuery("name", "jack")
	c.String(200, fmt.Sprintf("hello %s\n", name))
}

func submit(c *gin.Context) {
	name := c.DefaultQuery("name", "lily")
	c.String(200, fmt.Sprintf("hello %s\n", name))
}

func main() {
	// 创建默认路由
	// 默认使用了 2 个中间件 Logger(), Recovery()
	r := gin.Default()

	// 路由 1， 处理 GET请求
	v1 := r.Group("/v1")
	//  { } 是书写规范
	{
		v1.GET("/login", login)
		v1.GET("/submit", submit)
	}

	v2 := r.Group("/v2")
	{
		v2.POST("/login", login)
		v2.POST("/submit", submit)
	}

	// 选择默认路由
	r.Run()
}
