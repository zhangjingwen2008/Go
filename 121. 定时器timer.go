package main

import (
	"fmt"
	"time"
)

/*
定时器
	- 创建定时器：myTimer := time.NewTimer(n)
	- 停止：myTimer.Stop()		//定时器归零。读端<-myTimer.C会阻塞，因为channel写端的Timer被停止了，channel无数据写入但却未关闭
	- 重置：myTimer.Reset(time.Second)		//例如，将10秒的定时器重置为指定的3秒
*/

func main() {

	fmt.Println(time.Now()) //现在时间

	//定时器的：创建、停止、重置
	aTimer := time.NewTimer(time.Second * 10) //创建为10秒
	aTimer.Reset(3)                           //重置为3秒
	aTimer.Stop()                             //停止

	//3种定时方法
	//1.time.sleep
	time.Sleep(time.Second * 3)

	//2.定时器time.NewTimer
	myTimer := time.NewTimer(time.Second * 2)
	nowTime := <-myTimer.C //读取指针里的里面的C变量
	fmt.Println(nowTime)

	//3.time.After
	myTimer2 := time.After(time.Second * 3) //相比time.NewTimer，这个直接返回Time对象，更快一步
	nowTime2 := <-myTimer2
	fmt.Println(nowTime2)
}
