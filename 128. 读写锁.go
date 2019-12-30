package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
读写锁：sync.RWMutex
定义：var rwMutex sync.RWMutex
	1.加写锁：rwMutex.Lock()
	2.解写锁：rwMutex.Unlock
	3.加读锁：rwMutex.RLock()
	4.解读锁：rwMutex.RUnlock()
读时共享，写时独占
优先级：写锁>读锁（大量读取时，若有写操作，则会优先插入进行）
*/
var rwMutex sync.RWMutex
var value int //使用全句变量作为共享变量，来避免channel与读写锁混用造成的隐性死锁

func main() {

	rand.Seed(time.Now().UnixNano())
	//ch:= make(chan int)

	for i := 0; i < 5; i++ { //创建5个写go程
		go writeGO(i + 1)
	}

	for i := 0; i < 5; i++ { //创建5个读go程
		go readGo(i + 1)
	}

	for {

	}
}

func writeGO(n int) {
	for {
		num := rand.Intn(1000)
		rwMutex.Lock() //加写锁
		//out<-num			//channel与读写锁的混用，会造成隐性死锁
		value = num
		fmt.Println("--------第", n, "子go程写入", num)
		time.Sleep(time.Millisecond * 500)
		rwMutex.Unlock() //解写锁
	}
}

func readGo(n int) {
	for {
		rwMutex.RLock() //加读锁
		//num:=<-in
		num := value
		fmt.Println("第", n, "子go程读取：", num)
		rwMutex.RUnlock() //解读锁
	}
}
