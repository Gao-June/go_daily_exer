package main

import "fmt"

func main() {
	a := make(map[int]int)

	a[1] = 11
	a[2] = 22
	a[3] = 33
	a[4] = 44

	for index, value := range a {
		fmt.Println(index, value)
	}

}
