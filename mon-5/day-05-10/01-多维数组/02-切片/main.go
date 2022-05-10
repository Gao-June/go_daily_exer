/*
	练习切片
	通过对比 cap() 可知，切片的容量是 cap - 当前索引
*/

package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5}
	t := a[1:3:5]
	y := a[0:2:5]

	// t:[2 3]  len(t):2, cap(t):4
	fmt.Printf("t:%v  len(t):%v, cap(t):%v\n", t, len(t), cap(t))
	// y:[1 2]  len(y):2, cap(y):5
	fmt.Printf("y:%v  len(y):%v, cap(y):%v\n", y, len(y), cap(y))
}
