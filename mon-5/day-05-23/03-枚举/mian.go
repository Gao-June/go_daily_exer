/*
	iota 默认从0开始计数，可以从自定义的位置开始计数吗？
	=>
		可以的，不过要注意一下几个细节：
		用 iota去表示枚举数，类型是 const
		中间可以任意定义数和其值，再重新使用iota的时候会按照最开始排序（从0开始）
		该解决办法可以同样 iota + 10
*/

package main

import (
	"fmt"
)

const (
	a = iota + 10
	b
	c = 100
	d = iota
)

func main() {
	fmt.Println("a = ", a, "\tb = ", b, "\tc = ", c, "\td = ", d)
}
