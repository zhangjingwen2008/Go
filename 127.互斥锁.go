package main

import (
	"fmt"
	"sync"
	"time"
)

/*
sync.Mutex
创建一个互斥量（互斥锁）：var mutex sync.Mutex
	- 加锁：mutex.Lock()
	- 解锁：mutex.Unlock()

go程中，锁只有1把，但是可以有多种不同属性的锁（例如读写锁的读锁和写锁，虽然属性不同但始终还是1把锁）
互斥锁是“建议锁”，操作系统提供，作用是建议你在编程时使用
*/

var mutex sync.Mutex //定义一个全句锁

func main() {
	go person1()
	go person2()

	for {

	}
}

//打印机
func printer(n string) {
	mutex.Lock() //加锁
	for _, value := range n {
		time.Sleep(time.Millisecond * 300)
		fmt.Print(string(value))
	}
	mutex.Unlock() //解锁
}

func person1() {
	printer("hello")
}

func person2() {
	printer(" world")
}
