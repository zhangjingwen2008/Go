package main

import "fmt"

/*

单向channel：
	- 默认channel为双向的：var ch chan int | ch=make(chan int)
	- 单向写channel：var ch chan <- int	|	ch=make(chan <- int)
	- 单向读channel：var ch <- chan int	|	ch=make(<- chan int)

转换：
	1.双向channel可以隐式转换为任意一种单向channel
		sendChannel = ch
		readChannel = ch
	2.单向channel不能转换为双向channel
		ch = sendChannel/readChannel 	会报错

【注意】两端channel需要在不同的go程中才可正常运作
*/
func main() {

	//1.双向默认channel
	ch := make(chan int)
	fmt.Println(ch)

	//2.单向写channel
	var sendChannel chan<- int
	sendChannel <- 890

	//3.单向读channel
	var readChannel <-chan int
	num := <-readChannel
	fmt.Println(num)

}
