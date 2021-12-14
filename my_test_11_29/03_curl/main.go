/*
	书上的示例, http 简单使用
*/

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func init() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./example <url>")
		os.Exit(-1)
	}
}

func main() {
	r, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	io.Copy(os.Stdout, r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
}
