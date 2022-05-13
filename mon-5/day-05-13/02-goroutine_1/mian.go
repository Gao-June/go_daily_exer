/*
	练习使用 goroutine, 及 sync.WaitGroup
*/

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup // 自定义

func Hello(count int) {
	defer wg.Done() // defer add - 1
	fmt.Println("Hello:\t", count)
}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go Hello(i)
	}

	wg.Wait() // 执行阻塞，直到所有的WaitGroup数量变成0
}
