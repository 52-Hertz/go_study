package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		fmt.Println("等待客户端发送消息"+conn.RemoteAddr().String())
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("服务器的Read err = ", err)
			return
		}
		fmt.Print(string(buf[:n]))
	}
}

func main() {
	fmt.Println("服务端已经启动.....")
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("监听出错 err=", err)
		return
	}
	defer listen.Close()
	for {
		fmt.Println("等待客户端连接.....")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("客户端连接失败 err=%v\n", err)
		} else {
			fmt.Printf("客户端连接成功 coon=%v\n", conn)
		}
		go process(conn)
	}
}
