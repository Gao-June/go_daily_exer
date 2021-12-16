/*
	写一个 go 协程， 用 sync 做同步
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	for i := 1; i <= 10; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			fmt.Println("now is  ", id)
			time.Sleep(time.Millisecond * 500)
		}(i)
	}

	fmt.Println("main begin")

	wg.Wait()

	fmt.Println("main end")
}
