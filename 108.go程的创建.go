package main

import (
	"fmt"
	"time"
)

//关键字go：即可实现并发
func main() {
	go Dance()
	go Sing()

	for { //需要这个for循环，否则并发提前结束。因为主go程退出后，子go程也会退出

	}
}

func Dance() {
	for i := 0; i < 5; i++ {
		fmt.Println("------Dancing------")
		time.Sleep(100 * time.Millisecond)
	}
}

func Sing() {
	for i := 0; i < 5; i++ {
		fmt.Println("------Sing------")
		time.Sleep(100 * time.Millisecond)
	}
}
