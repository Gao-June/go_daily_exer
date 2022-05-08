/*
	测试 switch case 中单条件为空的情况
	单条件，内容为空、隐式 “ case x==10 : break; ”
*/

package main

import "fmt"

func main() {
	switch x := 10; {
	case x == 1:
		fmt.Println(1)
	case x == 10: // 单条件，内容为空、隐式 “ case x==10 : break; ”

	case x > 5:
		fmt.Println(3)
	}

	fmt.Println("hello world")

}
