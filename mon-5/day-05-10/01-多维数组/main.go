/*
	练习使用多维数组
*/

package main

import (
	"fmt"
)

func main() {
	// 注意最后一行 后面有一个 ,
	// 如果不想要的话，把最后的 } 写上去
	a := [...][2]string{
		{"北京", "上海"},
		{"成都", "武汉"},
		{"广州", "重庆"},
	}

	fmt.Println(a)
}
