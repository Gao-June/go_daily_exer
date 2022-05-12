package main

import (
	"code_test/code_exercise/mon-11/my_test_11_28/01_load_fun_test/Fun_sumss"
	"fmt"
)

func main() {
	var a, b int = 10, 2000

	fmt.Println(a, "  ", b)
	var c = Fun_sumss.Ssub(a, b)
	fmt.Println(c)
}
