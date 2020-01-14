package main

import (
	"fmt"
	"net/http"
)

func main() {
	//获取服务器内容
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println("http.Get ERR:", err)
		return
	}
	defer resp.Body.Close()

	//查看答应包
	fmt.Println("Header:", resp.Header)
	fmt.Println("Status:", resp.Status)
	fmt.Println("StatusCode:", resp.StatusCode)
	fmt.Println("Protocol:", resp.Proto)
	buf := make([]byte, 1024*4)
	var result string
	for {
		n, _ := resp.Body.Read(buf) //【注意】关闭主体是调用者的责任
		if n == 0 {
			fmt.Println("Read Finish----------------")
			//return		//不可用return，否则会来不及打印所有结果就退出了
			break
		}
		//if err!=nil{
		//	fmt.Println("resp.Body.Read ERR:",err)
		//	//return
		//	break
		//}
		result += string(buf[:n])
	}
	fmt.Println("|", result, "|")
}
