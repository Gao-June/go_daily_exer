/*
	练习调用 自定义库
*/

package main

import(
	"fmt"
	"go_daily_exer/mon-5/day-05-13/01-package/Fun_add_num"
)

func main( ){
	a, b := 12, 13
	var c int = Fun_add_num.Add_two_nums(a, b)

	fmt.Println("a + b = ", c)

}