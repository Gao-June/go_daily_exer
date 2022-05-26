/*
	TCP客户端程序的处理流程：
		1.建立与服务端的链接
    	2.进行数据收发
    	3.关闭链接

以下是 TCP client 端代码实现：
*/

package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:10010")
	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	// 关闭连接
	defer conn.Close()

	// 从 os.Stdin里读取数据信息
	inputReader := bufio.NewReader(os.Stdin)

	for {
		// ReadString读取用户输入，直到读到 '\n'。出错的话返回 err，这里没写
		input, _ := inputReader.ReadString('\n')

		// 返回一个string切片，扇区input字符前后的  "\r\n"
		inputInfo := strings.Trim(input, "\r\n")

		// 输入 Q 就退出
		if strings.ToUpper(inputInfo) == "Q" {
			return
		}

		// 发送数据
		_, err := conn.Write([]byte(inputInfo))
		if err != nil {
			return
		}

		buf := [512]byte{}
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("recv failed, err: ", err)
			return
		}

		fmt.Println(string(buf[:n]))
	}
}
