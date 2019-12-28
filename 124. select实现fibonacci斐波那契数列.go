package main

import (
	"fmt"
	"runtime"
)

/*
斐波那契数列：1 1 2 3 5 8，即后一项等于前两项的和
当x=1,y=1时，第3项就是x+y=2
其中的关系：
	x=y（新的x=上一轮的y）
	y=x+y（新的y=上一轮的x加y）
*/
func main() {
	ch := make(chan int)
	quit := make(chan bool)

	go fibonacci(ch, quit)

	x, y := 1, 1
	for i := 0; i < 20; i++ {
		ch <- x
		x, y = y, x+y //斐波那契数列计算
	}

	quit <- true //上面计算结束后，往channel传入停止标志

}

func fibonacci(ch <-chan int, quit <-chan bool) {
	for {
		select {
		case num := <-ch: //有斐波那契数进入channel时，打印出来
			fmt.Println(num, " ")
		case <-quit: //计算结束，则结束该子go程
			runtime.Goexit()
		}
	}
}
