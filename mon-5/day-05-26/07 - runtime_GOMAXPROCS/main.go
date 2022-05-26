// 练习使用 tuntime.GOMAXPROCS
// 为什么我电脑上看不出区别？

package main

import (
	"fmt"
	"runtime"
	"time"
)

func a() {
	for i := 1; i < 10; i++ {
		fmt.Println("A: ", i)
	}
}

func b() {
	for i := 1; i < 10; i++ {
		fmt.Println("B: ", i)
	}
}

func main() {
	fmt.Println("练习使用 runtime.GOMAXPROCS")

	// runtime.GOMAXPROCS(1)
	runtime.GOMAXPROCS(3)

	go a()
	go b()

	time.Sleep(time.Second)

}
