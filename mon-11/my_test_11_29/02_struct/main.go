/*
	结构体 声明
	推荐使用方法-3

*/

package main

import "fmt"

type stu struct {
	name   string
	age    int
	gender string
}

func main() {
	// eg.1
	fmt.Println("no.1")
	eg1 := stu{
		name:   "one",
		age:    11,
		gender: "boy",
	}
	fmt.Printf("%#v \n", eg1)

	// eg. 2
	fmt.Println("eg.2")
	eg2 := stu{"two", 12, "girl"}
	fmt.Printf("%#v \n", eg2)

	// eg. 3
	fmt.Println("eg.3")
	eg3 := stu{name: "eg3", age: 18, gender: "boy"}
	fmt.Printf("%#v\n", eg3)
}
