package main

import "fmt"

var i int

func adder() func(int) int {
	var x int

	return func(y int) int {
		fmt.Println("i = ", i, " x = ", x)

		i += 1
		x += y

		return x
	}
}

func main() {
	f := adder()

	fmt.Println(f(10))
	fmt.Println(f(10))
	fmt.Println(f(10))
}
