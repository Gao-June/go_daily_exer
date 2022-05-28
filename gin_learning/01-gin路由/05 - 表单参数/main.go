/*
	表单传输为 post 请求，http常见的传输格式为四种：
		- application/json
		- application/x-www-form-urlencoded
		- application/xml
		- multipart/form-data
	表单参数可以通过 PostForm() 方法获取，
		该方法默认解析的是x-www-form-urlencoded或from-data格式的参数
【代码运行】
	先运行 main.go，然后执行 name.html文件。
	在里面输入账号密码
*/

package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建默认路由
	r := gin.Default()

	// 绑定路由规则（POST请求），定义执行的函数
	r.POST("/form", func(c *gin.Context) {
		//types := c.DefaultPostFrom("type", "post")
		types := c.DefaultPostForm("type", "post")
		username := c.PostForm("username")
		password := c.PostForm("userpassword")

		c.String(http.StatusOK, fmt.Sprintf("username:%s, password:%s, type:%s",
			username, password, types))
	})

	// 运行默认路由
	r.Run()
}
