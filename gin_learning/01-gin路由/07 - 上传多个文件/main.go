/*
	上传多个文件

	p.s. 我这边代码出错，无法访问
*/

package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建默认路由
	// 默认使用了 2个中间件 Logger(), Recovery()
	r := gin.Default()

	// 限制表单上传大小 8MB， 默认为 32MB
	// r.MaxMultipartMemory = 8 << 20
	r.POST("/upload", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get err %s", err.Error()))
		}

		// 获取所有图片
		files := form.File["files"]

		// 遍历所有图片
		for _, file := range files {
			// 逐个存储
			if err := c.SaveUploadedFile(file, file.Filename); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("upload err: %s", err.Error()))
				return
			}
		}

		c.String(200, fmt.Sprintf("upload ok %d files", len(files)))
	})

	// 执行默认 路由
	r.Run()
}
