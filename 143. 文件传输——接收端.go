package main

import (
	"fmt"
	"net"
	"os"
)

/*
判断读取文件已到末尾的标志
	- 从本地读取：io.EOF
	- 从网络读取：n==0
*/

func main() {
	//1.创建监听，并Accept阻塞
	listener, err := net.Listen("tcp", "127.0.0.1:8008") //创建连接
	if err != nil {
		fmt.Println("net.Listener ERROR:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Accept()阻塞中")
	conn, err := listener.Accept() //进行Accept()阻塞
	if err != nil {
		fmt.Println("listener.Accept ERROR:", err)
		return
	}
	fmt.Println("监听连接成功")
	defer conn.Close()

	//2.接收消息，回传ok信息
	readBuf := make([]byte, 1024) //接收消息
	n, err := conn.Read(readBuf)
	if err != nil {
		fmt.Println("conn.Read ERROR:", err)
		return
	}
	filename := string(readBuf[:n])
	fmt.Println("接收到的fineName为:", filename)

	conn.Write([]byte("ok")) //回传ok

	//3.接收并保存文件
	receiveFile(conn, filename)
}

func receiveFile(conn net.Conn, filename string) {
	fmt.Println("开始保存文件")
	//创建文件
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println("os.Create ERROR:", err)
		return
	}
	defer f.Close()

	//写入文件
	buf := make([]byte, 1024*4)
	for {
		n, _ := conn.Read(buf)
		if n == 0 {
			fmt.Println("文件读取完成")
			return
		}
		f.Write(buf[:n])
	}
}
