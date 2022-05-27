// 练习使用 time 时间类型

package main

import (
	"fmt"
	"time"
)

func main( ){
// 简单表述时间
	now := time.Now()
	fmt.Println("time now: ", now)

	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Println("time: ", year, "-", month,
				 ".", day, "-", hour, 
				 ":", minute, ":", second)

// 记录时间差
	time_begin := time.Now()
	var num int = 1		// 添加操作，让时间长一点
	for i := 1; i < 1000000; i ++{
		num++
	}
	time_end  := time.Now()
	time_dur := time_end.Sub(time_begin)
	fmt.Println("时间间隔：", time_dur)

}