/*
	日志文件

	会在根路径生成 gin.log 文件来记录日志
*/

package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func My_fun(c *gin.Context) {
	c.String(200, "ping")
}

func main() {
	gin.DisableConsoleColor()

	// Loggin to a file
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	// 如果需要同时将 日志写入文件和控制台， 请使用以下代码
	// gin.DefaultWriter = io.MultiWriter( f, os.Stdout )
	r := gin.Default()

	r.GET("/ping", My_fun)

	r.Run()

}
