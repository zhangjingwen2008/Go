package main

import (
	"fmt"
	"time"
)

/*
定义周期定时器：myTicker = time.NewTicker(n)
	- 定时时长到达后，系统会自动向Ticker中的C写入系统当前时间
	- 每隔一个定时时长后，循环写入系统当前时间
	- 在子go程中循环读取C，获取系统写入时间

Timer和Ticker的区别：
	1.Timer系统只向C写入1次时间，而Ticker会不断循环写入
	2.Timer只可读取1次，Ticker可被不断循环读取
*/
func main() {
	fmt.Println("-------now=", time.Now())

	quit := make(chan bool) //判断退出go程的channel

	myTimer := time.NewTicker(time.Second) //定义一个周期定时器Ticker
	i := 0

	go func() {
		for {
			nowTime := <-myTimer.C
			fmt.Println(nowTime)
			i++
			if i == 8 { //i==8时退出周期定时
				quit <- true
				break
			}
		}
	}()

	<-quit //当i==8前代码都会阻塞在这里
}
