/*
scan 读取输入
*/

package main

import (
	"fmt"
	"io"
)

func main() {
	var a, b int
	_, err := fmt.Scan(&a, &b)
	if err == io.EOF {
		fmt.Println("err")
	}

	fmt.Println("a = ", a, "  b = ", b)

}
