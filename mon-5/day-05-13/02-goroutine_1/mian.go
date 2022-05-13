/*
	练习 goroutine, 练习 sync.WaitGroup
*/

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func Hello(count int) {
	defer wg.Done() // defer add - 1
	fmt.Println("Hello:\t", count)
}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go Hello(i)
	}

	wg.Wait()
}
