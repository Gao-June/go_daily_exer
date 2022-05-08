/* 使用 _ 可以表示匿名变量
Go 有其独特的特点，可以 return 多个变量
*/

package main

import (
	"fmt"
)

func foo() (int, string) {
	return 10, "hello_world"
}
func main() {
	var x int
	var y string

	x, _ = foo()
	_, y = foo()

	fmt.Println("x = ", x)
	fmt.Println("y = ", y)

}
