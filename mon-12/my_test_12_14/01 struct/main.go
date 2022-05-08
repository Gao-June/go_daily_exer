/*
	练习使用 struct
*/

package main

import "fmt"

type stu struct {
	age   int
	name  string
	class int
}

func main() {
	// test 1
	var Wu stu
	Wu.age = 15
	Wu.name = "小吴"
	Wu.class = 7
	fmt.Printf("%#v\n", Wu)

	// test 2
	Ming := stu{
		age:   18,
		name:  "小明",
		class: 4,
	}
	fmt.Printf("%#v\n", Ming)

	// test 3
	Zhang := stu{
		17,
		"小张",
		5,
	}
	fmt.Printf("%#v\n", Zhang)

	// test 4
	// 一般这样写,不过当数据量多的时候不建议。
	Hong := stu{19, "小红", 13}
	fmt.Printf("%#v\n", Hong)
}
