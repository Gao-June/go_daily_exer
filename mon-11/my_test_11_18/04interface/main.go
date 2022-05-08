package main

import (
	"fmt"
)

type dog struct {
	age  int
	name string
}

type Mover interface {
	move()
}

func (d *dog) move() {
	fmt.Println(" ddd ")

	fmt.Printf("%+v  \n", d)
	d.age = 10
	fmt.Printf("%+v  \n", d)

}

func main() {
	var i Mover
	var dd = dog{18, "旺财"}
	fmt.Printf("%+v  \n", dd)

	i = &dd

	i.move()

	fmt.Printf("%+v  \n", dd)
}
