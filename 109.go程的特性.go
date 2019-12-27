package main

import "fmt"

/*
本程序调用结果：子go程打印次数不足10次就结束了
结论：主go程退出后，其子go程都会中断并退出
*/
func main() {

	//子go程：匿名函数
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("----------子")
		}
	}()

	//主go程
	for i := 0; i < 3; i++ {
		fmt.Println("----------主go程----------")
	}

}
