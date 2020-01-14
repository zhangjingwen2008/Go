package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

/*
爬虫概念：访问web服务器，获取指定数据信息的一段程序
工作流程：
	1. 明确目标URL
	2. 发送请求，获取应答数据包
	3. 保存、过滤数据，提取有用信息
	4. 使用、分析得到数据信息
双向爬取：
	- 横向爬取：以页为单位
	- 纵向爬取：以一个页面内的条目为单位

百度贴吧爬虫
*/

func main() {
	//指定爬取起始页、终止页
	var start, end int
	fmt.Print("请输入爬取的起始页(start>=1):")
	fmt.Scan(&start)
	fmt.Print("请输入爬取的终止页(end>=start):")
	fmt.Scan(&end)

	working(start, end)
}

func working(start int, end int) {
	fmt.Printf("正在爬取第%d页到第%d页……\n", start, end)

	//循环爬每一页数据
	for i := start; i <= end; i++ {
		url := "https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn=" + strconv.Itoa((i-1)*50)
		result, err := HttpGet(url) //调用方法获取爬虫数据
		if err != nil {
			fmt.Println("HttpGet ERR:", err)
			continue //当前页出错而已，还能继续，所以不用return而是continue
		}

		//3.将读取到的网页数据，保存成一个文件
		//fmt.Println("result=",result)
		f, err := os.Create("D:/Temp/第 " + strconv.Itoa(i) + " 页.html")
		if err != nil {
			fmt.Println("os.Create ERR:", err)
			continue
		}
		f.WriteString(result)
		f.Close() //这里不使用defer来关闭的原因：因为是在for循环内，需要保存好一个文件，就关闭一个文件
	}
}

func HttpGet(url string) (result string, err error) {
	//1.访问指定url
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1 //将封装函数的内部错误，传出给调用者
	}
	defer resp.Body.Close() //别忘了关闭Body

	//2.循环读取网页数据，传出给调用者
	buf := make([]byte, 1024*4)
	for {
		n, err2 := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("读取网页完成")
			break
		}
		if err2 != nil && err2 != io.EOF {
			err = err2
			return
		}
		//累加每一次循环读取到的buf数据，存入result一次性返回
		result += string(buf[:n])
	}
	return
}
