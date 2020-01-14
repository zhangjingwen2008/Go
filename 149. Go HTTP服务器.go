package main

import (
	"fmt"
	"net/http"
)

func main() {
	//回调函数
	http.HandleFunc("/itcast", myHandler)

	//监听
	http.ListenAndServe("127.0.0.1:8000", nil)
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is a Web Server"))

	fmt.Println("Header:", r.Header)
	fmt.Println("URL:", r.URL)
	fmt.Println("Method:", r.Method)
	fmt.Println("Host:", r.Host)
	fmt.Println("RemoteAddr:", r.RemoteAddr)
	fmt.Println("Body:", r.Body)
}
