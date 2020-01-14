package main

import (
	"fmt"
	"net/http"
	"os"
)

//打开文件并解析到客户端
func OpenSendFile(fName string, w http.ResponseWriter) {
	pathFileName := "D:/Temp" + fName //拼接服务器本地物理地址进行访问
	f, err := os.Open(pathFileName)   //若打开失败，则证明文件不存在
	if err != nil {
		fmt.Println("Open File ERR:", err)
		w.Write([]byte("No such file!"))
		return
	}
	defer f.Close()

	buf := make([]byte, 1024*4) //解析文件
	for {
		n, _ := f.Read(buf)
		if n == 0 {
			return
		}
		w.Write(buf[:n]) //传输到客户端
	}
}

//回调函数
func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("客户端请求：", r.URL)
	OpenSendFile(r.URL.String(), w)
}

func main() {
	//注册回调函数
	http.HandleFunc("/", myHandler)
	//绑定监听地址
	http.ListenAndServe("127.0.0.1:8000", nil)
}
