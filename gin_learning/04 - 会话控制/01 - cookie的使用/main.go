/*
	练习 cookie 的使用
*/

package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Use_Cookies(c *gin.Context) {
	// 获取客户端是否携带 cookie
	cookie, err := c.Cookie("key_cookie")
	if err != nil {
		cookie = "NOtSet"

		// 给客户端设置 cookie
		// maxAge int, 单位：秒
		// path: cookie所在目录
		// domain string.域名
		// secure 是否只能通过 http 访问
		// httpOnly bool :是否允许别人通过 js 获取自己的 cookie
		c.SetCookie("key_cookie", "value_cookie", 60, "/", "localhost", false, true)
	}
	fmt.Println("cookie 的值是： ", cookie)
}

func main() {
	// 创建默认路由
	r := gin.Default()

	// 使用 GET 请求，并定义执行的规则
	r.GET("cookie", Use_Cookies)

	// 执行默认路由： 8080
	r.Run()
}
