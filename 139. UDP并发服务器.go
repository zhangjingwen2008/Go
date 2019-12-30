package main

import (
	"fmt"
	"net"
	"time"
)

/*
在接收UDP数据往后阶段，使用for循环即可
*/
func main() {
	//1.创建监听地址
	udpAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8001")
	if err != nil {
		fmt.Println("net.ResolveUDPAddr err:", err)
		return
	}

	//2.创建通信socket
	udpConn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println("net.ListenUDP err:", err)
		return
	}
	defer udpConn.Close()

	fmt.Println("等待接收UDP数据中......")
	//3.接收UDP数据
	readBuf := make([]byte, 1024*4) //实现并行：在这里外层嵌套for循环即可
	for {
		n, readAddr, err := udpConn.ReadFromUDP(readBuf)
		if err != nil {
			fmt.Println("udpConn.Read err:", err)
			return
		}

		//处理接收到的数据
		fmt.Printf("接收来自%s的数据：%s\n", readAddr, string(readBuf[:n]))
		dateNow := time.Now().String() + "\n"

		//4.写出数据到UDP
		_, err = udpConn.WriteToUDP([]byte(dateNow), readAddr)
		if err != nil {
			fmt.Println("udpConn.WriteToUDP err:", err)
			return
		}
	}

}
