/*
	探究下 golang 对类型转换的支持（隐式转换、显示转换）
	【注】 %t 可以打印类型

	结论：
		可以强制转换，但是要注意丢失精度问题；
		另， string 不能和 int/float32/等转换（可以使用 strconv包）
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int = 1
	var b float32 = 2.768

	fmt.Printf("类型是：%T\t", a)
	fmt.Println("a = ", a)
	fmt.Printf("类型是：%T\t", b)
	fmt.Println("b = ", b)

	fmt.Println("接下来试验强制转化")
	c := float32(a)
	d := int(b)
	fmt.Printf("类型是：%T\t", c)
	fmt.Println("c = ", c)
	fmt.Printf("类型是：%T\t", d)
	fmt.Println("d = ", d)

	// 转化为 string 需要 strconv包
	fmt.Println("string 类型转化")
	var s1 string = strconv.Itoa(a)
	fmt.Printf("类型是：%T\t", s1)
	fmt.Println("s1 = ", s1)
	var s2 string = strconv.Itoa(int(b)) // Itoa只接受 int型作为参数
	fmt.Printf("类型是：%T\t", s2)
	fmt.Println("s2 = ", s2)

}
