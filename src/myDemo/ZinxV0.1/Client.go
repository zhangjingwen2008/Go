package main

import (
	"fmt"
	"net"
	"time"
)

/*
	模拟客户端
*/
func main() {
	fmt.Println("client start...")

	time.Sleep(1 * time.Second)

	//1.连接远程服务器，得到一个conn连接
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("client start err, exit!")
		return
	}

	for {
		//2.连接调用write 写数据
		conn.Write([]byte("Hello Zinx V0.1.."))
		if err != nil {
			fmt.Println("write conn err!")
			return
		}

		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read buf err!")
			return
		}

		fmt.Println(string(buf[:cnt]))

		//cpu阻塞
		time.Sleep(time.Second)
	}

}
