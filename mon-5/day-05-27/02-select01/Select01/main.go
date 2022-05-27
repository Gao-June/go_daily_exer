/*
练习使用 select
	select可以同时监听一个或多个channel，直到其中一个channel ready
	如下代码可以发现 select 在接受完 s2 后就退出了
*/

package main

import (
	"fmt"
	"time"
)

func Test1(ch chan string) {
	time.Sleep(time.Second * 5)
	ch <- "test1"
}

func Test2(ch chan string) {
	time.Sleep(time.Second * 2)
	ch <- "test2"
}

func Show_to_main02() {
	fmt.Println("这是来自 mian01的代码")
}

func main() {
	fmt.Println("this is main01")

	// 2个管道
	output1 := make(chan string)
	output2 := make(chan string)

	// 跑2个子协程，写入数据
	go Test1(output1)
	go Test2(output2)

	// 用 select 监控
	select {
	case s1 := <-output1:
		fmt.Println("s1 = ", s1)
	case s2 := <-output2:
		fmt.Println("s2 = ", s2)
	}
	Show_to_main02()
}
