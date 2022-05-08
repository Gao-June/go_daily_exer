/*
	写一个有缓冲 channel
*/

package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	// 接收数据
	// p.s. 如果是 无缓冲channel，会阻塞在接收处，除非发送数据
	ch <- 10
	fmt.Println("接收成功")

	// 发送数据
	num := <- ch
	fmt.Println("发送成功,  num =  ", num)

}
