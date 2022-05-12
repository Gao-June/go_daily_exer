/*
方法和接收者
	方法限定了该函数的调用者
	只有满足类型的 调用者才可以调用该方法
*/

package main

import "fmt"

type person struct {
	age  int
	name string
}

// 这里的传参必须是指针类型， 不然修改操作只是针对副本，无法修改接收者变量本身。
func (p *person) set_age(age int) {
	p.age = 18
	fmt.Println(p)
}

func main() {
	wang := person{
		age:  100,
		name: "小王",
	}

	fmt.Println(wang)
	wang.set_age(20)
	fmt.Println(wang)

}
