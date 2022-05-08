package main

import (
	"fmt"
	"os"
)

func main() {
	fileName := "./abc.test"

	// 创建文件
	file1, err := os.Create(fileName)
	if err != nil {
		fmt.Println("create err: ", err.Error())
	} else {
		fmt.Printf("%s 创建成功！ %v  \n", fileName, file1)
	}
}
