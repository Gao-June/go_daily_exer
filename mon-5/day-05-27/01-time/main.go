// 练习使用 time 时间类型

package main

import (
	"fmt"
	"time"
)

func main( ){
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

}