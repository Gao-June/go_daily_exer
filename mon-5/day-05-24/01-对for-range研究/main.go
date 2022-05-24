/*
	在 for-range中，range仅会复制目标数据，并不会对源数据进行改动
	【wrong】 是可以修改的！
	验证：
		下面写两个for循环进行验证
*/

package main

import (
	"fmt"
)

func main() {
	data1 := [5]int{1, 2, 3, 4, 5}
	data2 := [5]int{1, 2, 3, 4, 5}

	// 非 for-range， 可以修改切片元素
	fmt.Println("test-1")
	for i := 0; i < 5; i++ {
		data1[i] += 100
		fmt.Printf("%d\t", data1[i])
	}
	fmt.Println()
	fmt.Println(data1)

	// for-range
	fmt.Println("test-2")
	for index, _ := range data2 {
		data2[index] += 100
		//val += 100
		fmt.Printf("%d\t", data2[index])
	}
	fmt.Println()
	fmt.Println(data2)

}
