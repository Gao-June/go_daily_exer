/*
	练习结构体
*/

package main

import "fmt"

type person_information struct {
	age        int
	name, city string
}

func main() {
	fmt.Println("test No.1")
	var str person_information
	str.age = 18
	str.city = "BeiJing"
	str.name = "张三"

	fmt.Printf("%v\n", str)
	fmt.Printf("%#v\n", str)

	fmt.Println("test No.2")
	str2 := person_information{20, "NanJing", "李四"}

	fmt.Printf("%v\n", str2)
	fmt.Printf("%#v\n", str2)
}
