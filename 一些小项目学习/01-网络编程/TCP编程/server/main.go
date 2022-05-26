/*
	TCP服务端程序的处理流程：
		1 - 监听端口
		2 - 接收客户端请求建立连接
		3 - 创建 goroutine 处理连接

以下是 TCP server 端代码实现：
*/

package main

import (
	"bufio"
	"fmt"
	"net"
)

// 处理函数
func process(conn net.Conn) {
	// 关闭链接
	defer conn.Close()

	// 死循环，一直监听
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte

		// 读取数据
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Println("read from client failed, err: ", err)
			break
		}

		recvStr := string(buf[:n])
		fmt.Println("收到 client 端发来的数据: ", recvStr)

		// 发送数据
		conn.Write([]byte(recvStr))
	}
}

// 主函数
func main() {
	// listen 监听
	listen, err := net.Listen("tcp", "127.0.0.1:10010")
	if err != nil {
		fmt.Println("listen failed, err: ", err)
		return
	}

	// 无限循环，在里面一直保持读取状态
	for {
		// 建立连接
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, err: ", err)
			continue
		}
		go process(conn)
	}

}
