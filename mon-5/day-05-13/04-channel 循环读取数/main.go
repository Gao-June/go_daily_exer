/*
用 for range 循环读取 channel里的数
	Go语言的并发模型是CSP（Communicating Sequential Processes），
	提倡通过通信实现共享内存（将数据通过管道发送进行通信）
	而不是通过共享内存而实现通信（通过修改某个共同的变量进行通信）。
*/

package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	// 开启 goroutine 将 0 ~ 100的数发送到 ch1 中
	go func() {
		for i := 1; i < 20; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	// 开启 goroutine 从 ch1 中接收值，并将值求平方后放到 ch2 中
	go func() {
		for {
			i, ok := <-ch1
			if !ok {
				break
			}
			ch2 <- i * i
		}
		close(ch2)
	}()

	// 在主 goroutine 中从 ch2 中接收值并打印
	for value := range ch2 {
		fmt.Println(value)
	}
}
