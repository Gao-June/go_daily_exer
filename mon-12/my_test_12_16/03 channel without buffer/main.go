/*
	写一个无缓冲的 channel

	这里要使用 for 循环进行 无限循环
*/

package main

import (
	"fmt"
)

func main() {

	// ch := make(chan int)
	// // 接收数据
	// ch <- 10
	// fmt.Println("接收成功")
	// 如果程序到这里结束，会报错： fatal error: all goroutines are asleep - deadlock!
	// 原因是缓冲区里的数据没有处理，在接收完数据之后就 阻塞 了！

	// 释放数据，下面代码仍然无法执行，仍然会报错：fatal error: all goroutines are asleep - deadlock!
	// num := <-ch
	// fmt.Println("数据发送完毕， num : ", num)

	// 正确做法，使用 gotoutine 并发执行程序

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := 1; i <= 10; i++ {
			ch1 <- i
			fmt.Println("发送数据： ", i)
		}
		fmt.Println("数据发送成功")

		close(ch1)
	}()

	go func() {
		for {
			i, ok := <-ch1
			if !ok {
				break
			}
			ch2 <- i
			fmt.Println("接收数据： ", i)
		}
		fmt.Println("数据接收成功 ")

		close(ch2)
	}()

	for num := range ch2 {
		fmt.Println("ch2 处接收数据： ", num)
	}

	fmt.Println("程序执行完毕")
}
