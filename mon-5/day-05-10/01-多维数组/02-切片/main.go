/*
	练习切片
*/

package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5}
	t := a[1:3:5]

	fmt.Printf("t:%v  len(t):%v, cap(t):%v\n", t, len(t), cap(t))
}
