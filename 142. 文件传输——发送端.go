package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	//1.发送文件名
	list := os.Args //从命令行获取文件路径
	if len(list) != 2 {
		fmt.Println("文件格式错误")
		return
	}

	filePath := list[1]
	fileInfo, fileInfoErr := os.Stat(filePath) //根据路径获取文件状态
	if fileInfoErr != nil {
		fmt.Println("os.Stat ERROR:", fileInfoErr)
		return
	}
	filename := fileInfo.Name()
	fmt.Println("要发送的文件名为：", filename)

	conn, err := net.Dial("tcp", "127.0.0.1:8008") //创建连接
	if err != nil {
		fmt.Println("net.Dial ERROR:", err)
		return
	}
	fmt.Println("发送端连接创建成功")
	defer conn.Close()

	_, err2 := conn.Write([]byte(filename)) //发送文件名
	if err2 != nil {
		fmt.Println("conn.Write ERROR:", err2)
		return
	}
	fmt.Println("文件名发送完毕")

	//2.接收ok信息
	readBuf := make([]byte, 16)
	n, err3 := conn.Read(readBuf) //接收消息
	if err3 != nil {
		fmt.Println("conn.Read ERROR:", err3)
		return
	}

	//3.发送文件
	if "ok" == string(readBuf[:n]) {
		fmt.Println("接收到ok信息！开始发送！")

		writeFile, err4 := os.Open(filePath) //打开要发送的文件
		if err4 != nil {
			fmt.Println("os.Open ERROR:", err4)
			return
		}
		defer writeFile.Close()

		fileBuf := make([]byte, 1024*4)
		for {
			n, err5 := writeFile.Read(fileBuf) //循环读取文件内容
			if err5 != nil {
				if err5 == io.EOF {
					fmt.Println("文件读取完毕")
				} else {
					fmt.Println("writeFile.Read:", err5)
				}
				return
			}
			_, err := conn.Write(fileBuf[:n]) //将文件发送
			if err != nil {
				fmt.Println("conn.Write ERROR:", err)
				return
			}
		}
	}

}
