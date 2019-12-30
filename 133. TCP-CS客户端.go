package main

import (
	"fmt"
	"net"
)

/*
客户端
1. conn,err:=net.Dial(network,address)
2. conn.Read() / conn.Write()
*/
func main() {

	//1.指定 服务器IP+port创建 通信套接字
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return
	}
	defer conn.Close()

	//2.主动写数据给服务器
	conn.Write([]byte("Are U OK?!"))

	//3.接收服务器回发的数据
	buf := make([]byte, 1024*4)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err:", err)
		return
	}
	fmt.Println("服务端回传:", string(buf[:n]))

}
