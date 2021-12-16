/*
	再试着写一次 goroutine + channel without buffer

	上一个测试代码是用了 for 无限循环来时间 channel通信的，这次直接看看

	可以，并不是要 for 无限循环
	重点是需要 Sleep()
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	// 两个 goroutine，一个发送、一个接收
	go func() {
		ch1 <- 10

		fmt.Println("已发送")
	}()

	go func() {
		num := <-ch1
		ch2 <- num

		fmt.Println("已接收")
	}()

	time.Sleep(time.Second)

	n := <-ch2
	fmt.Println("接收到的数据是： ", n)

	time.Sleep(time.Second)

	close(ch1)
	close(ch2)

}
