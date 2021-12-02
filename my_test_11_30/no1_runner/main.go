/*
	这个实例程序演示如何使用通道来监视
	程序运行的时间，以及程序运行时间过长时 如何终止程序
*/

package main

import (
	"code_test/code_exercise/my_test_11_30/no1_runner/runner"
	"log"
	"os"
	"time"
)

// timeout 规定了必须在多少秒内处理完成
const timeout = 3 * time.Second

// createTask() 返回一个根据 id 休眠指定秒数的示例任务
func createTask() func(int) {
	return func(id int) {
		log.Printf("processor  -  Task #%d. ", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}

// main 入口
func main() {
	log.Println("Starting work. ")

	// 为本次执行分配超时时间
	r := runner.New(timeout)

	// 加入要执行的任务
	r.Add(createTask(), createTask(), createTask())

	// 执行任务并处理结果
	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrInterrupt:
			log.Println("Terminatinf 超时. ")
			os.Exit(1)
		case runner.ErrInterrupt:
			log.Println("Terminating 中断.")
			os.Exit(2)
		}
	}

	log.Println("process ended.")
}
