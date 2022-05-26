/*
	练习使用 有缓冲 chan
*/

package main

import (
	"fmt"
	"time"
)

func recv(ch chan int) {
	for {
		index, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(index * index)
	}
}

func main() {
	ch := make(chan int, 1)

	for i := 1; i < 10; i++ {
		// 读数进去
		ch <- i
		go recv(ch) // 取
	}
	fmt.Println("读成功")

	time.Sleep(time.Second)
	fmt.Println("取成功")
	close(ch)
}
