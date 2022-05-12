/*
	包的导入
	写路径的时候，从src下面的开始写
*/

package main

import (
	"code_test/code_exercise/mon-5/day-05-12/04-package/Sum_nums"
	"fmt"
)

func main() {
	var a, b int = 10, 1500
	fmt.Println("a: ", a, "\tb: ", b)

	var c = Sum_nums.Sum(a, b)
	fmt.Println("a: ", a, "\tb: ", b, "\tc:", c)
}
