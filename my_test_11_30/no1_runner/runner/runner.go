/*
	runner 扩展包
	本项目用于展示 如何使用通道来监视程序的执行时间，如果程序运行时间太长，也可以用 runner 来终止
*/

package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

// Runner 在给定的超时时间内执行一组任务
// 	并且在操作系统中发送中断信号时结束这些任务
type Runner struct {
	// interrupt 通道报告从 操作系统 发送的信号
	interrupt chan os.Signal

	// complete 通道报告处理任务已经完成
	complete chan error

	// timeout 报告处理任务已经超时
	timeout <-chan time.Time

	// tasks 持有一组以索引顺序依次执行的函数
	tasks []func(int)
}

// ErrTimeout 会在任务执行超时返回
var ErrTimeout = errors.New("received timeout")

// ErrInterrupt 会在接收到 操作系统 的事件时返回
var ErrInterrupt = errors.New("received interrupt")

// New 返回一个新的准备使用的 Runner
func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}

// Add 将一个任务附加到 Runner 上
// 这个任务是一个 接收 一个 int 类型的 ID 作为参数的函数
// " ... " 的意思是 可变参数
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

// Start 执行所有任务， 并监视通道事件
func (r *Runner) Start() error {
	// 我们希望接受所有 终端信号
	signal.Notify(r.interrupt, os.Interrupt)

	// 用不同的 goroutine 执行不同的任务
	go func() {
		r.complete <- r.run()
	}()

	select {
	// 当任务处理完时发出的信号
	case err := <-r.complete:
		return err

	// 当任务处理程序运行超时发出的信号
	case <-r.timeout:
		return ErrTimeout
	}
}

// Run 执行每一个已注册的函数
func (r *Runner) run() error {
	for id, task := range r.tasks {
		// 检测操作系统的 中断信号
		if r.gotInterrupt() {
			return ErrInterrupt
		}

		// 执行已注册的任务
		task(id)
	}

	return nil
}

// gotInterrupt
func (r *Runner) gotInterrupt() bool {
	select {
	// 当中断事件被触发时
	case <-r.interrupt:
		// 停止接受后续的任何信号
		signal.Stop(r.interrupt)
		return true

	// 继续 正常运行
	default:
		return false
	}
}
