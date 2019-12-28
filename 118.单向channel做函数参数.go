package main

import "fmt"

func main() {
	ch := make(chan int) //定义一个双向channel作为参数

	go func() {
		send(ch)
	}()

	read(ch)
}

//写入channel函数
func send(out chan<- int) {
	out <- 889 //向写channel中写入数据
	close(out)
}

//读取channel函数
func read(in <-chan int) {
	n := <-in //最终在读channel中获得数据
	fmt.Println(n)
}
