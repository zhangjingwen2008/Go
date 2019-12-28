package main

import "fmt"

//订单对象
type OrderInfo struct {
	orderId int
}

func main() {
	ch := make(chan OrderInfo, 4) //创建一个传递订单的channel

	go producer(ch)
	consumer(ch)

}

//生产者
func producer(out chan<- OrderInfo) {
	for i := 0; i < 10; i++ { //模拟产生10份订单
		fmt.Println("-------生产订单：", i+1)
		order := OrderInfo{i + 1} //创建订单对象
		out <- order              //将订单放入channel用于传输
	}
	close(out) //channel用完要关闭
}

//消费者
func consumer(in <-chan OrderInfo) {
	for order := range in {
		fmt.Println("接收订单：", order.orderId)
	}
}
