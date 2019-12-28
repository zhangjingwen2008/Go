package main

import "fmt"

type Humaner interface { //接口1
	SayHello()
}
type Personer interface { //接口2，继承接口1
	Humaner
	Say()
}

type Student struct {
}

func (s *Student) SayHello() {
	fmt.Println("SayHello!!!")
}
func (s *Student) Say() {
	fmt.Println("Say!!!")
}

func main() {
	var stu Student
	var per Personer
	per = &stu

	//可调用所继承接口的方法
	per.Say()
	per.SayHello()

	//转换
	var h Humaner
	h = per //反过来则会出错
	h.SayHello()
}
