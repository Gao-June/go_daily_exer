/**
 * 练习 for 循环、数组、for-range
 */

package main

import (
	"fmt"
)

func main() {
	fmt.Println("No.1 test")

	for i, j := 1, 1; i <= 5 && j <= 5; i, j = i+1, j+1 {
		println(i, "  ", j)
	}

	fmt.Println("No.2 test")
	// 这里涉及到了数组
	nums := []int{1, 5, 2, 3, 7, 3, 4, 8}
	for i, num := range nums {
		fmt.Println("index: ", i, "\tvalue: ", num)
	}

}
