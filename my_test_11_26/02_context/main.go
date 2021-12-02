package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var exitChan = make(chan bool, 1)
var num int = 1

func worker() {
	defer wg.Done()

LOOPP:
	for {
		fmt.Println("hello, ", num)
		num += 1
		time.Sleep(time.Microsecond * 100)
		select {
		case <-exitChan:
			break LOOPP

		default:
		}
	}
}

func main() {
	wg.Add(1)

	go worker()

	time.Sleep(time.Second * 1)
	exitChan <- true
	wg.Wait()
}
