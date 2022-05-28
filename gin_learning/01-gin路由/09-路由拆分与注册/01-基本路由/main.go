/*
	最基础的gin路由注册方式，适用于路由条目比较少的简单项目或者项目demo。

	GET后面的参数类似一个 函数指针/lambda，拆开写好理解一些
*/

package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func hello_fun(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello welcome!",
	})
}

func main() {
	// 创建默认路由
	r := gin.Default()

	// 绑定路由，并运行定义的函数
	r.GET("/togo", hello_fun)
	if err := r.Run(); err != nil {
		fmt.Println("startup service failed, err: %v\n", err)
	}

}
