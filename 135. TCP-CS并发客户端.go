package main

import (
	"fmt"
	"net"
	"os"
)

/*
并发客户端
判断服务器关闭：conn.Read时读到的为0
*/
func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8001")
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return
	}
	defer conn.Close()
	fmt.Println("客户端连接成功！")

	//子go程：并发写入
	go func() {
		writeBuf := make([]byte, 1024*4)
		for {
			n, err := os.Stdin.Read(writeBuf) //os.Stdin.Read：接收从键盘输入
			if err != nil {
				fmt.Println("os.Stdin.Read err:", err)
				continue
			}
			conn.Write(writeBuf[:n])
		}
	}()

	//主go程：并发读取
	readBuf := make([]byte, 1024*4)
	for {
		n, err := conn.Read(readBuf)
		if n == 0 { //若服务端先关闭，则Read时的channel写入的那一端就消失了，此时从channel中读到的只有0
			fmt.Println("服务器断开，客户端关闭")
			return
		}
		if err != nil {
			fmt.Println("conn.Read err:", err)
			return
		}
		fmt.Println("读取到数据：", string(readBuf[:n]))
	}

}
