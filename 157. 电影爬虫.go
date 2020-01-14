package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

/*
豆瓣电影爬虫（现因触发反爬，返回的是错误代码418，无法进行爬虫）

横向爬取规律:
	第二页 https://movie.douban.com/top250?start=25&filter=
	第三页 https://movie.douban.com/top250?start=50&filter=
	第四页 https://movie.douban.com/top250?start=75&filter=
	总结规律：参数start以每页25的数值增加，那么第一页的start就是0

纵向爬取规律:
	电影标题：<img width="100" alt="【电影标题】"
	电影评分：<span class="rating_num" property="v:average">【电影评分】</span>
	评分人数：<span>【评分人数】人评价</span>
	总结规律：通过chrome浏览器的网页审查，发现所需数据的前后标识符
*/

func HttpGetDB(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}
	defer resp.Body.Close()

	//循环爬取数据项
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
	//先打印抬头（电影名称 评分评分人数）
	f.WriteString("电影名称\t\t\t评分\t\t\t评分人数\n")
	for i := 0; i < n; i++ {
		f.WriteString(filmName[idx][1] + "\t\t\t" + filmScore[idx][1] + "\t\t\t" + peopleNum[idx][1] + "\n")
	}
}

//爬取一个豆瓣页面信息
func SpiderPageDB(i int) {
	url := "https://movie.douban.com/top250?start=" + strconv.Itoa((i-1)*25) + "&filter="

	//横向爬取
	result, err := HttpGetDB(url)
	if err != nil {
		fmt.Println("HttpGet ERR:", err)
		return
	}

	//纵向爬取
	//电影名称
	ret1 := regexp.MustCompile(`<img width="100" alt="(.*?)"`)
	filmName := ret1.FindAllStringSubmatch(result, -1)
	//评分
	ret2 := regexp.MustCompile(`<span class="rating_num" property="v:average">(?s:(.*?))</span>`)
	filmScore := ret2.FindAllStringSubmatch(result, -1)
	//评分人数
	ret3 := regexp.MustCompile(`<span>(.*?)人评价</span>`)
	peopleNum := ret3.FindAllStringSubmatch(result, -1)

	//将提取的有用信息，封装到文件中
	Save2file(i, filmName, filmScore, peopleNum)
}

func workingDB(start int, end int) {
	fmt.Printf("开始爬取第%d页到第%d页的数据……\n", start, end)

	for i := start; i <= end; i++ {
		SpiderPageDB(i)
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
