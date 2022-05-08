/*
	写了一个调用 自定义库 的测试案例

	能正常运行，不过不知道为啥 调用库包 Sum_func 会提示问题。
*/

package main

import (
	"fmt"

	"go_daily_exer/my_test_12_15/01_customize_library/Sum_func"
)

func main() {
	a, b := 12, 13

	c := Sum_func.Ssum(a, b)

	fmt.Println(c)
}
