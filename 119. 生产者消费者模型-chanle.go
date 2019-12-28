package main

import "fmt"

/*
生产者消费者模型
	- 生产者：发送数据端
	- 消费者：接收数据段
	- 缓冲区
		1.解耦（降低生产者和消费者之间的耦合度）
		2.并发（生产着消费者数量不对等时，能保持正常通信）
		3.缓存（生产者和消费者数据处理速度不一致时，暂存数据）
*/
func main() {
	ch := make(chan int, 3)

	go producer(ch)
	consumer(ch)

}

//生产者
func producer(ch chan<- int) {
	for i := 0; i < 10; i++ {
		fmt.Println("生产：", i*i)
		ch <- i * i
	}
	close(ch)
}

//消费者
func consumer(ch <-chan int) {
	for in := range ch {
		fmt.Println("消费了：", in)
	}
}
