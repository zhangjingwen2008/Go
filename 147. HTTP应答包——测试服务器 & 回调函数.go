package main

import "net/http"

/*
使用net/http包 创建web服务器 步骤：
	1. 注册回调函数：http.HandleFunc()
	2. 绑定服务器监听地址：http.ListenAndServer()

回调函数
	- 本质：函数指针。通过地址，在某一特定位置，调用函数
	- 在程序中，定义一个函数，但不显示调用。当某一条件满足时，该函数由操作系统自动调用。
*/

func main() {
	//1.注册回调函数
	//参数1：用户访问文件位置
	//参数2：回调函数名
	http.HandleFunc("/itcast", handler)

	//2.绑定服务器监听地址
	//参数1：监听的地址
	//参数2：指定回调函数，通常使用nil，使其调用自身的回调函数
	http.ListenAndServe("127.0.0.1:8000", nil)
}

//回调函数，参数必须为 http.ResponseWriter 和 *http.Request
func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world!"))
}
