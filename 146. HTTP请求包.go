package main

import (
	"fmt"
	"net"
	"os"
)

//本程序获取到的HTTP请求包内容：

//GET / HTTP/1.1
//Host: 127.0.0.1:8000
//Connection: keep-alive
//Cache-Control: max-age=0
//Upgrade-Insecure-Requests: 1
//User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.142 Safari/537.36
//Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3
//Accept-Encoding: gzip, deflate, br
//Accept-Language: zh-CN,zh;q=0.9

/*
HTTP请求包格式：

请求行：请求方法 (空格) 请求文件URL (空格) 协议版本 (\r\n)
	- GET、POST
请求头：语法格式 key：value
空行：\r\n ———— 代表HTTP请求头结束
请求包体：请求方法对应的数据内容，GET方法没有内容
*/

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8000") //构建监听
	errMethod(err, "net.Listen")
	defer listener.Close()

	conn, err := listener.Accept() //获取监听连接
	errMethod(err, "listener.Accept()")
	defer conn.Close()

	buf := make([]byte, 1024*4) //捕获HTTP请求包
	n, err := conn.Read(buf)
	if n == 0 {
		return
	}
	errMethod(err, "conn.Read")
	fmt.Println(string(buf[:n]))
}

//封装异常
func errMethod(err error, errInfo string) {
	if err != nil {
		fmt.Println(errInfo, err)
		os.Exit(1)
	}
}
