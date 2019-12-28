package main

import (
	"fmt"
)

/*

使用语法：close(channel)

对端判断channel是否关闭：
	num, ok := <- channel		//num：从channel中接收到的数据；ok：false表示已关闭
	if ok{
		……
	}

1.数据不发送完，不应该关闭
2.已经关闭的channel，无法再向其写数据。报错：panic: send on closed channel
3.写端已经关闭的channel，可以从中读取数据。若为int类型，则读到0，说明写端关闭

*/
func main() {
	ch := make(chan int)

	//子go程
	go func() {
		for i := 0; i < 8; i++ {
			ch <- i
		}
		close(ch) //写端，写完数据主动关闭channel
	}()

	//主go程
	for {
		num, ok := <-ch //判断channel是否关闭
		if ok == true {
			fmt.Println("主go程读到：", num)
		} else {
			break
		}
	}
}
