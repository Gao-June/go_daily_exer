/*
	使用 goroutine 并发
*/

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func cA() {
	defer wg.Done()

	for count := 1; count <= 9; count++ {
		for ch := 'A'; ch <= 'A'+25; ch++ {
			fmt.Printf("%c  ", ch)
		}
	}

}

func ca() {
	defer wg.Done()

	for count := 1; count <= 9; count++ {
		for ch := 'a'; ch <= 'a'+25; ch++ {
			fmt.Printf("%c  ", ch)
		}
	}

}

func main() {
	wg.Add(2)

	go cA()
	go ca()

	wg.Wait()
}
