package main

import (
	"fmt"
	"time"
)

/*
channel：是一种数据类型，对应一个FIFO的“管道”
定义语法：make(chan 在channel中传递的数据类型，容量)
	容量=0：无缓冲channel
	容量>0：有缓冲channel
例如：make(chan int) 或 make(chan string, 0)
	- 此时长度len：channel中剩余未读取数据个数，为0
	- 此时长度cap：channel通道的容量，为0

本程序结果
未使用channel前：字符串打印乱序
使用channel后：字符串按顺序打印

【补充知识点】
	1.每当有一个进程启动时，系统会自动打开三个文件：
		- 标准输入、标准输出、标准错误
		- 对应3个文件：stdin（键盘）、stdout（屏幕）、stderr（屏幕）
	2.当进程运行结束，操作系统自动关闭三个文件
*/

//全局定义channel，用来完成数据同步
var channel = make(chan int)

func main() {
	go person1()
	go person2()

	for {

	}
}

//打印机
func printer(s string) {
	for _, ch := range s {
		fmt.Printf("%c", ch)
		time.Sleep(300 * time.Millisecond)
	}
}

//定义两个人使用打印机
func person1() {
	printer("hello")
	channel <- 8 //通道入口
}
func person2() {
	<-channel //通道出口。出口已建立但入口未建立时，通道会在这里卡住；一旦通道连通，才会继续往下执行，从而达到顺序打印的结果
	printer("world")
}
