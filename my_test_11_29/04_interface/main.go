/*
	练习 接口
*/

package main

import "fmt"

type say_hi interface {
	hi()
}

type dog struct{}

func (*dog) hi() {
	fmt.Println("旺财")
}

type cat struct{}

func (*cat) hi() {
	fmt.Println("杰瑞")
}

func main() {
	w := dog{}

	c := cat{}
	var x say_hi
	x = &w
	x.hi()

	x = &c
	x.hi()
}
