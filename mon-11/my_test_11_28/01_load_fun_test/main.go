package main

import (
	"fmt"

	"code_test/code_exercise/my_test_11_28/01_load_fun_test/Fun_sum"
)

func main() {
	var a, b int = 10, 20

	fmt.Println(a, "  ", b)
	var c = Fun_sum.Ssum(a, b)
	fmt.Println(c)
}
