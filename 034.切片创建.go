package main

import "fmt"

func main() {
	//创建方法1
	var Slice1 []int
	fmt.Println(Slice1) //结果为[]

	//创建方法2
	Slice2 := []int{}
	fmt.Println(Slice2) //结果为[]

	//创建方法3:make()函数
	Slice3 := make([]int, 3, 5) //数值3:初始化空间; 数值5:容量(可省略，省略后 容量==长度)
	fmt.Println(Slice3)
	fmt.Println(len(Slice3)) //结果为3，内容为[0,0,0]
	fmt.Println(cap(Slice3)) //结果为5，cap函数为查看容量
}
