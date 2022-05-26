/*
	练习使用 func() strings.Trim

	func Trim(s, cutset string) string
	Trims返回字符串s的一个切片，删除剪切集中包含的所有前导和尾随 cutset 码位。
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "abcdefgabcrtesabc"
	str2 := "abc"

	str3 := strings.Trim(str1, str2) // defgabcrtes

	fmt.Println("str1: ", str1, "\nstr2: ", str2, "\nstr3: ", str3)
}
