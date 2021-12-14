/*
	练习使用 eerors.New 创造新的错误
*/

package main

import (
	"errors"
	"fmt"
)

func main( ){
	err1 := errors.New("参数 x 出现了新错误")

	x := 17

	if x > 20{
		fmt.Println("yes")
	} else{
		fmt.Println(err1)
	}


}