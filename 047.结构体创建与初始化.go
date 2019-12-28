package main

import "fmt"

func main() {
	//创建
	type Student struct {
		id   int
		name string
		age  int
		addr string
	}

	//初始化方式1：全部初始化
	var s Student = Student{101, "张三", 23, "深圳"}
	fmt.Println(s)

	//初始化方式2：部分初始化
	var s1 Student = Student{name: "initial01", age: 22}
	fmt.Println(s1)

	//初始化方式2：部分初始化
	var Stu Student
	Stu.age = 21
	Stu.addr = "Test"
	fmt.Println(Stu)
}
