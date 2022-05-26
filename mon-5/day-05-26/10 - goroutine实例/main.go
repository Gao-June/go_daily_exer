/*
	这里写一个 goroutine 实例， 同时练习一下 chan
		本质上是生产者消费者模型
		可以有效控制goroutine数量，防止暴涨
	需求：
		计算一个数字的各个位数之和，例如数字123，结果为1+2+3=6
		随机生成数字进行计算
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 生产者
type Job struct {
	Id      int // 编号
	RandNum int // 产生的随机值
}

// 消费者
type Result struct {
	job *Job // 传入对象实例	why？
	sum int  // 求和值
}

// 创建工作池
// 参数1： 开几个协程
func createPool(num int, jobChan chan *Job, resultChan chan *Result) {
	// 根据协程个数 去运行
	for i := 0; i < num; i++ {
		go func(jC chan *Job, rC chan *Result) {
			// 执行运算
			// 遍历 job 管道所有数据，进行相加
			for job := range jC {
				// 随机数接过来
				r_num := job.RandNum

				// 随机数每一位相加
				// sum 位数求和返回值
				var sum int = 0
				for r_num != 0 {
					tmp := r_num % 10
					sum += tmp
					r_num /= 10
				}

				// 想要结果放入 Result
				r := &Result{
					job: job,
					sum: sum,
				}
				// 运算结果 给 chan
				rC <- r
			}
		}(jobChan, resultChan)
	}
}

func main() {
	// 需要 2 个chan用来传入数 和 求值
	// 1. job 管道
	jobChan := make(chan *Job, 128)
	// 2. result 管道
	resultChan := make(chan *Result, 128)

	// 3. 创建工作池
	createPool(64, jobChan, resultChan)

	// 4. 利用 goroutine 打印数据
	go func(ch chan *Result) {
		// 遍历结果管道打印
		for result := range ch {
			fmt.Printf("job id: %v\trandnum: %v\tresult: %v\n", result.job.Id,
				result.job.RandNum, result.sum)
		}
	}(resultChan)

	var id int = 0
	var count int = 20 //只计算 20 组数据
	// 循环创建 job， 输入到管道
	for count > 0 {
		
		time.Sleep(time.Second) // 不然协程运行20次太快
		count--
		id++

		// 生成随机数
		r_num := rand.Int() // 这里是生成一个非负的数， [0, 2^64)
		job := &Job{
			Id:      id,
			RandNum: r_num,
		}
		jobChan <- job

	}
}
