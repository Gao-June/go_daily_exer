/*
	并发在 无缓冲管道 读取值（任务）
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// runner 模拟接力跑步
func Runner(baton chan int) {
	var newRunner int

	// 等接力棒
	runner := <-baton

	// 开始跑步
	fmt.Printf("Runner %d Running witg baton\n", runner)

	// 创建下一位跑步者
	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("runner %d to the line\n", runner)
		go Runner(baton)
	}

	// 围绕跑道跑
	time.Sleep(100 * time.Microsecond)

	// 比赛结束了吗？
	if runner == 4 {
		fmt.Printf("runner %d finished, race over \n", runner)
		wg.Done()
		return
	}

	// 比赛未结束，将接力棒交给下一位选手
	fmt.Printf("runner %d exchange with runner %d \n", runner, newRunner)

	baton <- newRunner
}

func main() {
	// 创建一个无缓冲通道
	baton := make(chan int)

	// 为最后一位跑步者将计数 +1
	wg.Add(1)

	// 第一位跑步者持有接力棒
	go Runner(baton)

	// 开始比赛
	baton <- 1

	// 等待比赛结束
	wg.Wait()
}
