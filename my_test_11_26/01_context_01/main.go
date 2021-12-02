package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func worker() {
	defer wg.Done()

	for {
		fmt.Println("hello")
		time.Sleep(time.Microsecond * 200)
	}
}

func main() {
	wg.Add(1)

	go worker()

	wg.Wait()
	fmt.Println("finish!")
}
