/*
	在抖音看到一个有意思的线程时间问题

	描述：
		在线程函数中执行  a++, b++
		在主函数中 判断是否有 a < b 的情况
	分析：
		在处理 < 的时候可能会有时间间隙：
		首先读到 a 的时候，此时 a == b
		然后处理 < 的时候，a、b 可能执行了 ++ 操作，这时候导致 a < b；
	解决办法：
		设置定值， a_value = a; b_value = b
*/

package main

import (
	"fmt"
	"sync"
)

var a, b int = 0, 0
var wg sync.WaitGroup

func Add_nums() {
	for i := 1; i < 10000; i++ {
		a++
		b++
	}

	wg.Done()
}

func Is_equal() {
	/*
		a_value = a
		b_value = b
	*/

	if a < b {
		fmt.Println("a = ", a, "\tb = ", b)
	}
}

func main() {
	wg.Add(1)

	// 开始执行2个携程
	fmt.Println("程序开始")
	go Add_nums()
	go Is_equal()

	wg.Wait()
	fmt.Println("程序结束")
}
