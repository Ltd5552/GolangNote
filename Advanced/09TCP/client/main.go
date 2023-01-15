package main

import (
	"fmt"
	"net"
)

// tcp client
func main() {

	//	1.与server端建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("dial 127.0.0.1 failed,err:", err)
		return
	}
	conn.Write([]byte("hello world!"))
	conn.Close()
	//	2.发送数据
}
