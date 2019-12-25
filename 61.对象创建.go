package main

import "fmt"

//Go结构体==Java类
type Student struct {
	//Go成员==Java属性
	id   int
	name string
	age  int
}

//Go函数==Java方法
func main() {
	//对象
	stu1 := Student{1, "Hunter", 20}
	stu2 := Student{2, "Sally", 10}
	fmt.Println(stu1, stu2)
}
