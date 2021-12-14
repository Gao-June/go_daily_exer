/*
	无论是普通循环，还是 range迭代，其定义的局部变量都会重复使用

	i :  0   num:  10
	&i :  0xc000012078   &num:  0xc000012090
	i :  1   num:  20
	&i :  0xc000012078   &num:  0xc000012090
	i :  2   num:  30
	&i :  0xc000012078   &num:  0xc000012090
	i :  3   num:  40
	&i :  0xc000012078   &num:  0xc000012090
	i :  4   num:  50
	&i :  0xc000012078   &num:  0xc000012090
	i :  5   num:  60
	&i :  0xc000012078   &num:  0xc000012090
	i :  6   num:  70
	&i :  0xc000012078   &num:  0xc000012090
*/

package main

import "fmt"

func main() {
	nums := []int{10, 20, 30, 40, 50, 60, 70}

	for i, num := range nums {
		fmt.Println("i : ", i, "  num: ", num)
		fmt.Println("&i : ", &i, "  &num: ", &num)
	}
}
