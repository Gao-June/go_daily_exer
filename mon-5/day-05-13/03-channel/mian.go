/*
	练习管道 channel
*/

package main

import(
	"fmt"
	"sync"
) 

var wg sync.WaitGroup

// channel
func Recv(c chan int) {
	ret := c
	fmt.Println("接收成功：", ret)
	wg.Done()
}

func main() {
	ch := make(chan int, 1) // channel
	wg.Add(1)
	go Recv(ch)
	ch <- 10
	fmt.Println("发送成功")

	wg.Wait()
}
