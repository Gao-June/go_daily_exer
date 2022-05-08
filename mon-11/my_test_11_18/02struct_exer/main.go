package main

import "fmt"

type person struct {
	name string
	age  int
	city string
}

func main() {
	a1 := person{
		"张三",
		18,
		"湖北",
	}

	a2 := &person{
		name: "李四",
		age:  10,
		city: "重庆",
	}

	fmt.Println(a1)
	fmt.Println(a2)

	fmt.Printf("%+v\n", a1)
	fmt.Printf("%+v\n", a2)
}
