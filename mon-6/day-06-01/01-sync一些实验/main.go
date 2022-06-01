// 验证一下 sync WaitGrop 里的参数

// state[0] waiter: 等待者计数
// state[1] counter: 任务计数
// state[2] sema: 信号量

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Printf("WaitGroup statep: %v\n", wg) // [0 0 0]
	wg.Add(1)

	fmt.Printf("WaitGroup statep: %v\n", wg) // [0 1 0]
	wg.Done()

	fmt.Printf("WaitGroup statep: %v\n", wg) // [0 0 0]
	wg.Add(1)

	fmt.Printf("WaitGroup statep: %v\n", wg) // [0 1 0]
	wg.Add(1)

	fmt.Printf("WaitGroup statep: %v\n", wg) // [0 2 0]
	wg.Done()

	fmt.Printf("WaitGroup statep: %v\n", wg) // [0 1 0]
	wg.Done()

	fmt.Printf("WaitGroup statep: %v\n", wg) // [0 0 0]
	wg.Wait()

	fmt.Printf("WaitGroup statep: %v\n", wg) // [0 0 0]
}
