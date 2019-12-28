package main

import (
	"fmt"
	"runtime"
	"time"
)

/*
超时处理的机制：利用了select其中的一个case。当其他case都不满足时，就进入以time.After(n)为条件的case，若n时间内其他case仍未被触发，则此case将作为超时处理而触发。
*/
func main() {
	ch := make(chan int)
	quit := make(chan bool)

	go func() {
		for {
			select {
			case num := <-ch:
				fmt.Println("num=", num)
			case <-time.After(time.Second * 3): //超时处理代码
				fmt.Println("超时！！！")
				quit <- true
				runtime.Goexit()
			}
		}
	}()

	for i := 0; i < 2; i++ {
		time.Sleep(time.Second * 2)
		ch <- i
	}

	<-quit
	fmt.Println("Finish!!!")

}
