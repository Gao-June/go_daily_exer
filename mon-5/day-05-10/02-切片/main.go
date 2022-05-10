/*
练习切片
	通过对比 cap() 可知，切片的容量是 cap - 当前索引

	切片的本质就是对底层数组的封装，
	它包含了三个信息：底层数组的指针、切片的长度（len）和切片的容量（cap）

切片的复制：
	切片的复制是 “浅拷贝” 共享了底层数据， 一旦更改值的话 a/t/y 都会更改
	解释：切片是引用类型，所以s1和s2b其实都指向了同一块内存地址。修改s2的同时s1的值也会发生变化。

切片的遍历
*/

package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5}
	t := a[1:3:5]
	y := a[0:2:5]

	fmt.Println("test No.1")
	fmt.Printf("a:%v  len(a):%v, cap(a):%v\n", a, len(a), cap(a))
	// t:[2 3]  len(t):2, cap(t):4
	fmt.Printf("t:%v  len(t):%v, cap(t):%v\n", t, len(t), cap(t))
	// y:[1 2]  len(y):2, cap(y):5
	fmt.Printf("y:%v  len(y):%v, cap(y):%v\n", y, len(y), cap(y))

	// 切片的复制是 “浅拷贝” 共享了底层数据， 一旦更改值的话 a/t/y 都会更改
	// 解释：切片是引用类型，所以s1和s2b其实都指向了同一块内存地址。修改s2的同时s1的值也会发生变化。
	fmt.Println("test No.2")
	t[0] = 100
	fmt.Printf("a:%v  len(a):%v, cap(a):%v\n", a, len(a), cap(a))
	fmt.Printf("t:%v  len(t):%v, cap(t):%v\n", t, len(t), cap(t))
	fmt.Printf("y:%v  len(y):%v, cap(y):%v\n", y, len(y), cap(y))

	// 切片的遍历
	fmt.Println("test No.3")
	for i := 0; i < len(a); i++ {
		fmt.Printf("%v\t", a[i])
	}
	fmt.Println()

	// 切片的遍历
	fmt.Println("test No.4")
	for _, value := range a {
		fmt.Printf("%v\t", value)
	}
	fmt.Println()

}
