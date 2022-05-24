/*
	探究 defer 和 return 的顺序

	说实话，这个代码没看懂 0.0
*/

package main

import (
	"fmt"
)

func Func_print() (num int) {
	fmt.Println("No.2")
	// 这里用了一个 匿名函数
	defer func() {
		fmt.Println("No.4")
		fmt.Println("defer:\t", num)
		num += 100
	}()
	fmt.Println("No.3")
	return 100
}

func main() {
	fmt.Println("No.1")
	fmt.Println("test:\t", Func_print())
	fmt.Println("No.5")
}
