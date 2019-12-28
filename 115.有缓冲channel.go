package main

import "fmt"

/*
有缓冲channel定义语法：make(chan int,n)	n表示channel容量cap，当元素个数达到n时才会阻塞
	- 通道容量非0，len为channel中剩余未读取数据的个数，cap为通道的容量
	- channel应用于两个go程中，一个读，一个写
	- 缓冲区可以进行数据存储，存储至容量上限才阻塞。
	- 具有异步能力，不需要同时操作channel缓冲区（就像发短信，发出去后，不需要接收人一定在线立马收到，可以延迟再收）
*/
func main() {

	ch := make(chan int, 3) //无缓冲channel定义：存满3个元素之前，不会阻塞

	go func() {
		for i := 0; i < 8; i++ {
			ch <- i
			fmt.Println("子go程：i=", i, " len(channel):", len(ch))
		}
	}()

	for i := 0; i < 8; i++ {
		num := <-ch
		fmt.Println("主go 程读到：", num, " len(channel):", len(ch))
	}

}
