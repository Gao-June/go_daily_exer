/*
	读取文件
*/

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("D:/golang/golang_test/src/code_test/code_exercise/mon-5/day-05-12/06-file/info.txt")
	if err != nil {
		fmt.Println("文件打开失败， err: ", err)
		return
	}
	defer file.Close()

	var tmp = make([]byte, 128)
	n, err := file.Read(tmp)
	if err == io.EOF {
		fmt.Println("文件读完了")
		return
	}
	if err != nil {
		fmt.Println("读取文件的时候遇到错误, err: ", err)
		return
	}

	fmt.Printf("总共读取了 %d 字节数据\n", n)
	fmt.Println(string(tmp[:n]))
}
