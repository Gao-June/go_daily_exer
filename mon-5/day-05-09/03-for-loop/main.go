/**
 * 练习 for 循环
 */

package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello_world")

	for i, j := 1, 1; i <= 5 && j <= 5; i, j = i+1, j+1 {
		println(i, "  ", j)
	}
}
