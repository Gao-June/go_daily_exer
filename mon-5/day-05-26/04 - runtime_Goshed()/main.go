// Gosched让出处理器，允许运行其他goroutines。它不会挂起当前的goroutine，因此会自动继续执行。

package main

import (
    "fmt"
    "runtime"
)

func main() {
    go func(s string) {
        for i := 0; i < 5; i++ {
            fmt.Println(s)
        }
    }("world")

    // 主协程
    for i := 0; i < 5; i++ {
        // 切一下，再次分配任务
        runtime.Gosched()
        fmt.Println("hello")
    }
}