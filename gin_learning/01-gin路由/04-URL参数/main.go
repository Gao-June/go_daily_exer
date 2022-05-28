/*
	URL 参数可以通过 DefaultQuery() 或 Query() 方法获取
	DefaultQuery() 若不传参数，返回默认值，Query() 若不存在，返回空串
	API ? name=非默认值

	【输出】
		1 - 若输入： localhost:8080/user
			输出：hello 枯藤
		2 - 若输入： localhost:8080/user?name=张三
			输出：hello 张三
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

	// 绑定路由规则，定义执行的函数
	r.GET("/user", func(c *gin.Context) {
		// 指定默认值
		// 默认将 name 设置为 "枯藤"
		name := c.DefaultQuery("name", "枯藤")

		c.String(http.StatusOK, fmt.Sprintf("hello %s", name))
	})

	//默认监听 8080 端口
	r.Run()
}
