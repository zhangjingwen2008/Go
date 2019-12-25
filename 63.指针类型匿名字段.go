package main

import "fmt"

type Person struct {
	id   int
	name string
	age  int
}

type Student struct {
	*Person //父类变成指针类型
	score   float64
}

func main() {
	//初始化对象时，父类对象需要加上&符号，来将指针指向此对象
	var stu Student = Student{&Person{101, "张三", 32}, 96}

	fmt.Println(stu)      //结果为内存地址：{0xc000044400 96}
	fmt.Println(stu.name) //结果为字段数据：张三
}
