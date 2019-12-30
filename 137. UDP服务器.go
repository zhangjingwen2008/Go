package main

import (
	"fmt"
	"net"
	"time"
)

/*
TCP通信：面向连接的，可靠的数据包传输
UDP传输：无连接的，不可靠的报文传输，但会比TCP快

UDP相比TCP来说：
	- 优点：效率高、开销小、开发复杂度低
	- 稳定性差、安全低、无序

使用场景：
	TCP：对数据传输安全性、稳定性要求较高的场合。例如网络文件传输、下载、上传
	UDP：对数据实时传输要求较高的场合。例如视频直播、在线电话会议、游戏

创建UDP服务器步骤：
	1.创建监听地址：ResolveUPDAddr(network, string) (*UDPAddr, error)
	2.创建用户通信的socket：ListenUDP(network, *UDPAddr) (*UDPConn, error)
	3.接收UDP数据：(c *UDPConn) ReadFromUDP(b []byte) (int, *UDPAddr, error)
	4.写出数据到UDP：(c *UDPConn) WriteToUDP(b []byte, *UDPAddr) (int, error)

【注意】将time.Now()转换成string
	- string(time.Now())：会报错，无法转换
	- time.Now().String()：可行

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
	readBuf := make([]byte, 1024*4)
	n, readAddr, err := udpConn.ReadFromUDP(readBuf)
	if err != nil {
		fmt.Println("udpConn.Read err:", err)
		return
	}

	//处理接收到的数据
	fmt.Printf("接收来自%s的数据：%s\n", readAddr, string(readBuf[:n]))
	dateNow := time.Now().String()

	//4.写出数据到UDP
	_, err = udpConn.WriteToUDP([]byte(dateNow), readAddr)
	if err != nil {
		fmt.Println("udpConn.WriteToUDP err:", err)
		return
	}

}
