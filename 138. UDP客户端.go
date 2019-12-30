package main

import (
	"fmt"
	"net"
)

/*
UDP的客户端，和TCP的客户端实现原理，基本一致
*/
func main() {

	conn, err := net.Dial("udp", "127.0.0.1:8001") //仅需修改这里network为udp即可
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return
	}
	defer conn.Close()

	conn.Write([]byte("Are U OK?!"))

	buf := make([]byte, 1024*4)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err:", err)
		return
	}
	fmt.Println("服务端回传:", string(buf[:n]))

}
