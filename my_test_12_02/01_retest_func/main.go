package main

import (
	"fmt"
	f "my_owner_src/func_test_sum"
)

func main() {

	var a, b int = 10, 11
	var c = f.Two_sums(a, b)

	fmt.Println(c)  
}
