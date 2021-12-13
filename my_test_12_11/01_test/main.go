package main

import "fmt"

var x = 0x12

func main() {
	var y int = x
	//y = x
	
	fmt.Println(y, "  ", x)
	
}
