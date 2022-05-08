package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}

	newslice := slice[1:3]

	fmt.Println(slice)    // [1 2 3 4 5]
	fmt.Println(newslice) //[2 3]

	newslice = append(newslice, 10)

	fmt.Println(slice)    // [1 2 3 10 5]
	fmt.Println(newslice) //[2 3 10]

	slice = append(slice, 200)

	fmt.Println(slice)
	fmt.Println(newslice)

	newslice = append(newslice, 11, 12, 13)
	fmt.Println(slice)
	fmt.Println(newslice)
}
