package main

import "fmt"

type Student struct {
	id   int
	name string
	age  int
}

//结构体方法1：打印
func (s Student) PringShow() {
	fmt.Println(s)
}

//结构体方法2：修改
func (s Student) EditInfo() { //此修改不会影响到原对象
	s.age = 111
}

//结构体方法3：指针修改
func (s *Student) EditInfoWithPointer() { //会对原对象造成修改
	s.age = 222
}

func main() {
	stu := Student{1, "Hunter", 10}
	stu.PringShow() //结果：{1 Hunter 10}
	stu.EditInfo()  //修改
	stu.PringShow() //结果：{1 Hunter 222}
}
