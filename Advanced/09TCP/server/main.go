package main

import (
	"fmt"
	"net"
)

//tcp server端

func processConn(conn net.Conn) {
	//3、与客户端通信
	var temp [128]byte
	n, err := conn.Read(temp[:])
	if err != nil {
		fmt.Println("read from conn failed,err:", err)
		return
	}
	fmt.Println(string(temp[:n]))
}
func main() {
	//1、本地端口启动服务
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("start tcp failed,err:", err)
		return
	}
	for {
		//2、等待别人连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept failed,err:", err)
			return
		}
		go processConn(conn)
	}

}
