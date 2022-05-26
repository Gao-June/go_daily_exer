// 再写一个来练习 runtime.Goexit()

package main

import(
	"fmt"
	"runtime"
	"time"
)

func f1( ) ( ){
	for i := 0; i < 5; i ++{
		if i == 2{
			runtime.Goexit()
		}

		fmt.Println(i)
	}
}

func f2( ) ( ){
	for i := 10; i < 15; i ++{
		fmt.Println(i)
	}
}

func main( ){
	go f1()
	go f2()

	time.Sleep( time.Second*4 )
}