/*
	练习使用 管道 channel

	无缓冲 channel
*/

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

// player 模拟一个选手在打网球
func player(name string, court chan int) {
	// 在函数退出时调用 Done 来通知 mian 函数工作已经完成
	defer wg.Done()

	// 无限循环
	for {
		// 等到信号从 channel 传来
		ball, ok := <-court
		if !ok {
			// 通道关闭状态，赢了！
			fmt.Printf("player %s win\n", name)
			return
		}

		// 选随机数，作为退出的条件
		n := rand.Intn(100)
		//var n = 13
		if n%13 == 0 {
			//if n >= 90 {
			fmt.Printf("player %s missed! \n", name)
			fmt.Println("now, n  = ", n)
			close(court)

			return
		}

		// 显示击球数， 并将数量 + 1
		fmt.Printf("player %s hit %d \n", name, ball)
		ball++

		court <- ball

	}
}

func main() {
	// 创建一个 无缓冲通道
	court := make(chan int)

	// 计数 +2，表示要等待 2 个 goroutine
	wg.Add(2)

	// 启动 2 个选手
	go player("tom", court)
	go player("jerry", court)

	// 发球
	court <- 1

	// 等待结束
	wg.Wait()
}
