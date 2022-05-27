/*
	可以通过Context的Param方法来获取API参数
	localhost:8000/xxx/zhangsan
【知识点】
	1- 使用了 strings.Trim() 进行字符串切割
	2-  "/user/:name/*action"
		这里没弄懂  :name、*action 是什么意思，只知道做了切割
		=> ":" 这个只取1个， "*" 取后面全部
		如果没写 ":" 的话，就要按要求一模一样输入，而不是自定义输入
*/

package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建默认路由
	r := gin.Default()

	// 绑定路由规则
	// 【注】 Param : 路由变量
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")

		// 截取
		action = strings.Trim(action, "/")
		c.String(http.StatusOK, name+"  is  "+action)
	})

	//默认监听 8080 端口
	r.Run(":8080")
}
