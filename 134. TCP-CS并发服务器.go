package main

import (
	"fmt"
	"net"
	"strings"
)

/*
使用for循环Accept()的方式，来实现并发的接收客户端
客户端输入exit关闭时
	- nc或Linux等发出的字符串，结尾都带有\n
	- Windows发出的字符串，结尾都带有\r\n
*/
func main() {
	//配置服务器
	listener, err := net.Listen("tcp", "127.0.0.1:8001")
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	defer listener.Close()

	//并发监听客户端
	for {
		fmt.Println("Connecting......")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err:", err)
			return
		}
		go HandlerClient(conn) //循环开启go程创建服务器
	}

}

//客户端处理
func HandlerClient(conn net.Conn) {
	//连接远程客户端
	defer conn.Close()
	addr := conn.RemoteAddr() //读取客户端连接IP
	fmt.Println(addr, " 连接成功！")

	//循环读取客户端发送数据
	buf := make([]byte, 1024*4)
	for {
		n, err := conn.Read(buf)
		in := strings.ToLower(string(buf[:n]))
		if n == 0 { //关闭连接提示————当Read返回为0时，即为写入端关闭
			fmt.Println(addr, " 服务器已关闭，断开连接！！！")
			return
		}
		if "exit\n" == in || "exit\r\n" == in { //【客户端输入exit关闭】：若使用nc或Linux访问时，输入exit后还会带一个\n，所以要放入考虑；Windows系统则为\r\n
			fmt.Println(addr, " 主动退出")
			return
		}
		if err != nil {
			fmt.Println("conn.Read err:", err)
			return
		}
		fmt.Println("服务端接收到的数据：", string(buf[:n]))

		//将数据传回给客户端
		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
	}
}
