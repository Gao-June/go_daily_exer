/*
练习使用 select
	如果多个 channel 同时 ready，则随机选择执行一个
*/

package main

import (
	"fmt"
	ff "go_daily_exer/mon-5/day-05-27/04-test_import/Select"
	"time"
)

func main() {
	// 创建2个管道
	int_chan := make(chan int, 1)
	strint_chan := make(chan string, 1)

	go func() {
		int_chan <- 1
	}()
	go func() {
		strint_chan <- "hello"
	}()

	select {
	case value := <-int_chan:
		fmt.Println("int: ", value)
	case value := <-strint_chan:
		fmt.Println("string: ", value)
	}

	fmt.Println("main 结束")
	time.Sleep(time.Second)

	ff.Ssum()
}
