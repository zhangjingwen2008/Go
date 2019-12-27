package main

import "fmt"

func main() {
	//1.空指针：未被初始化的指针，会报错
	var p *int
	fmt.Println(*p)

	//2.野指针：被一片无效的地址空间初始化，无法被编译
	var pp *int = 0xff0004c080
	fmt.Println(*pp)

}
