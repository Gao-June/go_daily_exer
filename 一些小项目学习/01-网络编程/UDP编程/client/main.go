/*
	UDP客户端 client

	对代码改了下，可以一直会话（直到用户输入”Q“退出）
*/

package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("连接服务器失败, err: ", err)
		return
	}

	defer socket.Close()

	// 我在这里改一下，让它能够一直通信，直到输入 "Q"才退出
	for {
		//sendData := []byte("Hello Server!")
		// 这里改了一下，改成自定义输入
		var infomations string
		fmt.Scan(&infomations)
		sendData := strings.Trim(infomations, "\r\n")
		if strings.ToUpper(sendData) == "Q" {
			fmt.Println("会话结束")
			return
		}

		_, err = socket.Write([]byte(sendData)) // 发送数据
		if err != nil {
			fmt.Println("发送数据失败, err: ", err)
			return
		}

		data := make([]byte, 4096)
		n, remoteAddr, err := socket.ReadFromUDP(data) // 接收数据
		if err != nil {
			fmt.Println("接收数据失败, err: ", err)
			return
		}

		fmt.Printf("recv: %v \taddr: %v \t count: %v\n", string(data[:n]), remoteAddr, n)
	}

}
