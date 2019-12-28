package main

import (
	"fmt"
	"runtime"
	"time"
)

/*
select
作用：用来监听channel上数据流动
用法：与switch case语句基本一致，但case后面必须是IO操作，不可任意写判别表达式
注意：
	1.监听的case，中没有满足监听条件，就阻塞
	2.监听的case中，有多个满足监听条件，就任选一个执行
	3.可以使用default来处理所有case都不满足监听条件的状况（通常不使用default，因为会产生忙轮训）
	4.select自身不带有循环机制，需要借助外层for来循环监听
	5.break只能跳出select中的一个case，而无法跳出外层的for循环。若要跳出就要使用runtime.Goexit()或return
*/
func main() {
	ch := make(chan int)    //用来数据通信的channel
	quit := make(chan bool) //用来判断是否退出的channel

	go func() { //子go程写数据
		for i := 0; i < 5; i++ {
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
		quit <- true //通知主go程 退出
		runtime.Goexit()
	}()

	for { //主go程读数据
		select {
		case num := <-ch:
			fmt.Println("读到：", num)
		case <-quit:
			//break					//break跳出select
			//runtime.Goexit()		//退出主协程
			return
		}
		fmt.Println("----------")
	}
}
