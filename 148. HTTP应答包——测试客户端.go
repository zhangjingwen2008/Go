package main

import (
	"fmt"
	"net"
	"os"
)

//本程序获取的HTTP应答包内容：

//HTTP/1.1 200 OK
//Date: Mon, 13 Jan 2020 23:55:45 GMT
//Content-Length: 12
//Content-Type: text/plain; charset=utf-8

/*
HTTP应答包格式：

状态行：协议版本号 (空格) 状态码 (空格) 状态码描述 (\r\n)
响应头：语法格式：key：value
空行：\r\n
响应包体：请求内容存在：返回请求页面内容；不存在：返回错误页面描述
*/

//模拟浏览器
func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	errMethod(err, "Dial")
	defer conn.Close()

	httpRequest := "GET /itcast HTTP/1.1\r\nHost:127.0.0.1:8000\r\n\r\n" //构建最简单的HTTP请求（行+头+空行）
	conn.Write([]byte(httpRequest))                                      //发送数据给服务器

	buf := make([]byte, 1024*4) //接收请求
	n, _ := conn.Read(buf)
	if n == 0 {
		return
	}
	fmt.Println(string(buf[:n]))
}

//封装异常
func errMethod(err error, errInfo string) {
	if err != nil {
		fmt.Println(errInfo, err)
		os.Exit(1)
	}
}
