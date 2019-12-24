package main

import "fmt"

type Student struct {
	id   int
	name string
	age  int
	addr string
}

func main() {
	Stu := Student{101, "张三", 18, "北京"}
	fmt.Println(Stu)
}

func PrintDemo(stu Student) {
	stu.name = "update" //尝试对结构体修改，其结果不会影响到原结构体
}
