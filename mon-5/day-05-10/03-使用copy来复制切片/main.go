/*
	在 02 中说明了切片是引用类型，类似于 C++ 里的浅拷贝，所复制的对象指向了同一块地址。
	本节使用 copy() 复制，实现深拷贝

	Go语言的 copy() 函数可以迅速将一个切片的数据复制到另一个切片中。
*/

package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := make([]int, 5) // len 不能为空，  cap可以为空
	copy(b, a)

	fmt.Println("test No.1")
	fmt.Println(a)
	fmt.Println(b)

	b[2] = 200
	fmt.Println("test No.2")
	fmt.Println(a)
	fmt.Println(b)
}
