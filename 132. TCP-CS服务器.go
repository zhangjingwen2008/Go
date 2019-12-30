package main

import (
	"fmt"
	"net"
)

/*
服务器
1. listener,err:=net.Listen(network,address)
2. conn,err:=listener.Accept()
3. conn.Read() / conn.Write()
*/
func main() {
	//1.指定服务器的：通信协议+IP地址+port
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Listen err:", err)
	}
	defer listener.Close()

	//2.阻塞监听客户端连接请求，成功建立连接，返回用于通信的socket
	fmt.Println("客户端已启动，准备接收连接...")
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("listener.Accept err:", err)
		return
	}
	defer conn.Close()

	//3.读取客户端传来的数据
	fmt.Println("已连接客户端！")
	buff := make([]byte, 1028*4)
	n, err := conn.Read(buff)
	if err != nil {
		fmt.Println("conn.Read err:", err)
		return
	}
	conn.Write(buff[:n]) //将读到的数据回传给客户端
	fmt.Println("接收消息:", string(buff[:n]))

}
