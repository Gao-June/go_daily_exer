/*
	切片是 引用类型
*/

package main

import "fmt"

func fun_change(s_temp []int) {
	s_temp[2] = 100
	fmt.Println(s_temp)
}

func main() {
	s := make([]int, 0)
	s = append(s, 1, 2, 3, 4, 5)

	fmt.Println(s)

	fun_change(s)

	fmt.Println(s)

	// 可以直接用索引修改值
	s[1] = 101
	fmt.Println(s)
}
