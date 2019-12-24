package main

import "fmt"

type Student struct {
	id   int
	name string
	age  int
	addr string
}

func main() {
	stu := Student{101, "张三", 18, "北京"}

	var p *Student
	p = &stu

	fmt.Println(*p)        //结果为结构体的所有内容
	fmt.Println((*p).name) //结果为指定字段内容
	fmt.Println(p.name)    //结果为指定字段内容

	p.name = "Update" //对结构体指针进行修改
	fmt.Println(*p)

	Update(p) //将结构体指针作为参数
	fmt.Println(*p)
}

func Update(p *Student) {
	p.name = "AfterUpdate"
}
