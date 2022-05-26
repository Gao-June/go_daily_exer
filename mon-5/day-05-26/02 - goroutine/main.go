// 写一个 携程 并发

package main

import (
	"sync"
	"fmt"
)

var wg sync.WaitGroup

func Print_num( num int ) () {
	defer wg.Done()
	fmt.Println("num : ", num)
}

func main( ){
	for i:=0; i < 10; i ++{
		wg.Add(1)
		go Print_num(i)
	}

	wg.Wait()
}