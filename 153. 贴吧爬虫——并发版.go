package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

/*
并发版：
	1.封装爬取一个网页内容的代码到SpiderPage(index)函数中
	2.在working函数for循环启动go程调用SpiderPage() ——> n个待爬取页面，对应n个go程
	3.为防止主go程提前结束，引入channel实现主go程与子go程的同步
	4.在SpiderPage()结尾处（一个页面爬取完成），向channel中写内容
	5.在working函数添加新的for循环，从channel不断的读取各个子go程写入的数据。n个子go程————写n次channel————读n次channel

防止主go程提前结束的方法：
	1.主go程结尾使用for死循环
	2.引入channel来阻塞主go程，子go程达成条件再放通
*/
func main() {
	var start, end int
	fmt.Print("请输入爬取的起始页(start>=1):")
	fmt.Scan(&start)
	fmt.Print("请输入爬取的终止页(end>=start):")
	fmt.Scan(&end)

	working(start, end)

}

func working(start int, end int) {
	fmt.Printf("正在爬取第%d页到第%d页……\n", start, end)

	//3.防止主go程提前结束，引入channel
	page := make(chan int)

	//2.使用循环启动子go程来实现并发爬取
	for i := start; i <= end; i++ {
		go SpiderPage(i, page)
	}

	//5.循环channel来读取子go程的数据
	for i := start; i <= end; i++ {
		fmt.Printf("第%d个页面完成\n", <-page)
	}

}

//1.爬取单个页面的函数
func SpiderPage(i int, page chan int) {
	url := "https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn=" + strconv.Itoa((i-1)*50)
	result, err := HttpGet(url)
	if err != nil {
		fmt.Println("HttpGet ERR:", err)
		return
	}

	f, err := os.Create("D:/Temp/第 " + strconv.Itoa(i) + " 页.html")
	if err != nil {
		fmt.Println("os.Create ERR:", err)
		return
	}
	f.WriteString(result)
	f.Close()

	//4.一个子go程爬取完成，就向channel写入数据
	page <- i
}

func HttpGet(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
	}
	defer resp.Body.Close()

	buf := make([]byte, 1024*4)
	for {
		n, err2 := resp.Body.Read(buf)
		if n == 0 {
			break
		}
		if err2 != nil && err2 != io.EOF {
			err = err2
			return
		}
		result += string(buf[:n])
	}
	return
}
