/*
	并发在 带缓冲管道 读取值（任务）
*/

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

const (
	numberGoroutines = 4  // 要使用的 goroutine 数量
	taskLoad         = 10 // 要处理的任务数
)

// 初始化
func init() {
	// 随机初始化种子
	rand.Seed(time.Now().Unix())
}

// worker 作为 goroutine 启动来处理
// 从有缓冲的通道传入的工作
func worker(tasks chan string, worker int) {
	// 通知函数已经返回
	defer wg.Done()

	// 无限循环
	for {
		// 等待分配工作
		task, ok := <-tasks
		// 通道是空的，并且已经被关闭了。
		// p.s. 通道空 且 未关闭的话， 会处于阻塞状态
		if !ok {
			fmt.Printf("worker %d 工作完了\n", worker)
			return
		}

		// 开始工作
		fmt.Printf("worker %d 正在工作 ： %s\n", worker, task)

		// 模拟一段时间 作为 工作的时间
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		// 显示工作完成了
		fmt.Printf("worker: %d 完成了 工作 %s\n", worker, task)
	}
}

func main() {
	// 创建一个有缓冲的通道
	tasks := make(chan string, taskLoad)

	// 启动 goroutine 来完成任务
	wg.Add(numberGoroutines)
	for i := 1; i < numberGoroutines; i++ {
		go worker(tasks, i)
	}

	// 往 缓冲通道 里增加任务
	for pos := 1; pos <= taskLoad; pos++ {
		tasks <- fmt.Sprintf("任务 %d ", pos)
	}

	// 所有任务完成时关闭通道
	// 以便所有 goroutine 退出
	close(tasks)

	wg.Wait()
}
