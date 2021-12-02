package main

import (
	"fmt"
	"os"
)

func main() {
	fileName := "./abc.test"

	err := os.Remove(fileName)
	if err != nil {
		fmt.Println(err)
	} else {

		fmt.Printf("%s 删除成功 ", fileName)
	}
}
