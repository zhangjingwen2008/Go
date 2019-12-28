package main

import "fmt"

/*
无缓冲channel定义语法：channel:=make(chan int) 即不指定第2个参数即可
	- len=0，通道容量为0，所以不能存储数据
	- channel应用于两个go程中，一个读，一个写
	- 具备同步能力（类似打电话，需要两端人同时在线）

本程序调用结果：
	子go程： 0
	子go程： 1
	-----------主go程： 0
	-----------主go程： 1
	子go程： 2
	子go程： 3
	-----------主go程： 2
	-----------主go程： 3
	子go程： 4
	-----------主go程： 4
使用了channel仍然不是顺序打印的原因：fmt.Println()是一个io操作，多任务同时访问同一个硬件设备屏幕时，仍需要排队一个个访问。在排队的时候下一轮go程仍在继续运行
*/

func main() {

	channel := make(chan int) //定义无缓冲channel

	//子go程
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("子go程：", i)
			channel <- i
		}
	}()

	//主go程
	for i := 0; i < 5; i++ {
		j := <-channel
		fmt.Println("-----------主go程：", j) //问题出在这：打印到屏幕是io操作，会对流程造成缓慢堵塞
	}

}
