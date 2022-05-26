/*
	练习使用 无缓冲 chan
	有一个细节：无缓冲 chan需要有接收的时候才能发送，不然会 deadlock
	=> 可以使用 协程 来接收
*/

package main

import "fmt"

func recv(c chan int) {
	ret := <-c
	fmt.Println("接收成功： ", ret)
}

func main() {
	ch := make(chan int) // 无缓冲 chan
	go recv(ch)          // 启用 goroutine 从通道接收值
	ch <- 10
	fmt.Println("发送成功")
}
