package main

import "fmt"

//定义接口
type Personer interface { //对接口命名：用er结尾
	SayHello()
}

type Student struct {
}
type Teacher struct {
}

//实现接口
func (s *Student) SayHello() {
	fmt.Println("老师好")
}
func (s *Teacher) SayHello() {
	fmt.Println("同学好")
}

func main() {
	//调用方式1：
	var stu Student
	stu.SayHello()
	var tea Teacher
	tea.SayHello()

	//调用方式2：通过接口变量来调用
	//注意：此方式的前提，必须实现接口中所有声明的方法，否则会报错
	var student Personer
	student = &stu
	student.SayHello() //调用的是student的实现方法
	var teacher Personer
	teacher = &tea
	teacher.SayHello() //调用的是teacher的实现方法
}
