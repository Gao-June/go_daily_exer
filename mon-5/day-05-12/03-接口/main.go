/*
接口
*/

package main

import "fmt"

type Sayer interface {
	say()
}

type Namer interface {
	name()
}

type dog struct {
	age   int
	names string
}

type cat struct {
	age   int
	names string
}

func (d *dog) say() {
	fmt.Println("狗叫:\t", d.age)
}
func (c *cat) say() {
	fmt.Println("猫叫:\t", c.age)
}

func (d *dog) name() {
	fmt.Println("狗名:\t", d.names)
}

func (c *cat) name() {
	fmt.Println("猫名:\t", c.names)
}

func main() {
	var s Sayer
	var n Namer

	dd := dog{
		names: "小旺财",
		age:   18,
	}

	cc := cat{
		names: "小凯缇",
		age:   20,
	}

	fmt.Println("test No.1")
	s = &dd
	s.say()
	n = &dd
	n.name()

	fmt.Println("test No.2")
	s = &cc
	s.say()
	n = &cc
	n.name()
}
