package main

import (
	"fmt"
	"runtime"
)

/*
Gosched作用：出让当前go程所占用的CPU时间片。当再次获得CPU时，从出让位置继续恢复执行
语法：runtime.Gosched()
*/
func main() {

	//子go程：匿名函数
	go func() {
		for {
			fmt.Println("----------子")
		}
	}()

	//主go程
	for {
		runtime.Gosched() //出让所占用时间片
		fmt.Println("----------主go程----------")
	}
}
