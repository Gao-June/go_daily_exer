/*
上传单个文件
	multipart/form-data 格式用于文件上传
	gin文件上传与原生的 net/http 方法类似，不同在于gin把原生的request封装到c.Request中
*/

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建默认 路由
	r := gin.Default()

	// 限制上传最大尺寸
	r.MaxMultipartMemory = 8 << 20

	// 绑定路由规则（POST），定义执行的函数
	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.String(500, "上传图片出错")
		}

		c.SaveUploadedFile(file, file.Filename)
		c.String(http.StatusOK, file.Filename)
	})

	// 运行默认路由 8080
	r.Run()
}
