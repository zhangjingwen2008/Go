package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

func HttpGetDB(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
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
		fmt.Println(string(buf[:n]))
		result += string(buf[:n])
	}
	return
}

func Save2file(idx int, filmName, filmScore, peopleNum [][]string) {
	filePath := "D:/Temp/" + "第" + strconv.Itoa(idx) + "页.txt"
	f, err := os.Create(filePath)
	if err != nil {
		fmt.Println("os.Create ERR:", err)
		return
	}
	defer f.Close()

	n := len(filmName)
	f.WriteString("电影名称\t\t\t评分\t\t\t评分人数\n")
	for i := 0; i < n; i++ {
		f.WriteString(filmName[idx][1] + "\t\t\t" + filmScore[idx][1] + "\t\t\t" + peopleNum[idx][1] + "\n")
	}
}

func SpiderPageDB(i int, page chan int) {
	url := "https://movie.douban.com/top250?start=" + strconv.Itoa((i-1)*25) + "&filter="

	result, err := HttpGetDB(url)
	if err != nil {
		fmt.Println("HttpGet ERR:", err)
		return
	}

	ret1 := regexp.MustCompile(`<img width="100" alt="(.*?)"`)
	filmName := ret1.FindAllStringSubmatch(result, -1)
	ret2 := regexp.MustCompile(`<span class="rating_num" property="v:average">(?s:(.*? ))</span>`)
	filmScore := ret2.FindAllStringSubmatch(result, -1)
	ret3 := regexp.MustCompile(`<span>(.*?)人评价</span>`)
	peopleNum := ret3.FindAllStringSubmatch(result, -1)

	Save2file(i, filmName, filmScore, peopleNum)

	page <- i //爬取结束，向channel写入数据
}

func workingDB(start int, end int) {
	fmt.Printf("开始爬取第%d页到第%d页的数据……\n", start, end)

	page := make(chan int) //构建channel

	for i := start; i <= end; i++ {
		go SpiderPageDB(i, page)
	}

	for i := start; i <= end; i++ { //循环监控读取channel
		fmt.Printf("第%d页爬取完成\n", <-page)
	}

}

func main() {
	var start, end int
	fmt.Print("请输入要爬取的起始页：")
	fmt.Scan(&start)
	fmt.Print("请输入要爬取的终止页：")
	fmt.Scan(&end)

	workingDB(start, end)

}
