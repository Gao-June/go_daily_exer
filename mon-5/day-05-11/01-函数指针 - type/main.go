/*
	golang 中也能实现  C++里的函数指针
	要先 type 声明
*/

package main

import "fmt"

type calculation_nums func(int, int) int

func add_nums(a, b int) int {
	return a + b
}

func sub_nums(a, b int) int {
	return a - b
}

func main() {
	var c calculation_nums
	c = add_nums

	fmt.Println(c(1, 2))

	c = sub_nums
	fmt.Println(c(1, 2))

	// 也能直接做类型推导
	d := add_nums
	fmt.Println(d(100, 200))
}
