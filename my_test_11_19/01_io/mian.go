// 无法执行，报错 CreateFile main.go: The system cannot find the file specified.

package main

import (
	"fmt"
	"os"
)

func printMessage(filePath string) {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		fmt.Println("err: ", err.Error())
	} else {
		fmt.Printf("数据类型是： %T \n", fileInfo)
		fmt.Println("文件名：  ", fileInfo.Name())
		fmt.Println("文件大小： ", fileInfo.Size())
		fmt.Println("文件权限： ", fileInfo.Mode())
		fmt.Println("文件最后修改时间： ", fileInfo.ModTime())
	}

}

func main() {
	path := "D:/golang/golang_test/src/code_test/code_exercise/my_test_11_19/1234.txt"

	printMessage(path)
}
