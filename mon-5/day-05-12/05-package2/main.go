/*
	¡∑œ∞ π”√ package
*/

package main

import (
	"code_test/code_exercise/mon-5/day-05-12/05-package2/Fun_change_num"
	"fmt"
)

func main() {
	var a, b int = 100, 200

	c, d := Fun_change_num.Exchange(a, b)
	fmt.Println("\ta:\t", a, "\tb:\t", b)
	fmt.Println("\tc:\t", c, "\td:\t", d)
}
