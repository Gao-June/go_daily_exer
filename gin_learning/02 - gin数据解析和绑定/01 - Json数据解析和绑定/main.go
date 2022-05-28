/*
	Json 数据解析和绑定

	客户端传参，后端接收 并 解析到结构体

	没看明白
*/

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 定义接收数据的结构体
type Login struct {
	// binding : "required"修饰的字段
	//			若接收为空值，则报错，是必须字段
	User    string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
	Pssword string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

func My_fun(c *gin.Context) {
	// 声明接收的变量
	var json Login

	// 将request的body中的数据，自动按照json格式解析到结构体
	if err := c.ShouldBindJSON(&json); err != nil {
		// 返回错误信息
		// gin.H封装了生成json数据的工具
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	// 判断用户名密码是否正确
	if json.User != "root" || json.Pssword != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "200"})
}

func main() {
	// 创建路由
	// 默认使用了 2 个中间件 Logger(), Recovery()
	r := gin.Default()

	// JSON 绑定
	r.POST("loginJSON", My_fun)

	// 执行默认端口 8080
	r.Run()
}
