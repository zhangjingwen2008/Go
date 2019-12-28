package main

import "fmt"

//多态：使用同一个接口，使用不同的实例而产生不同的结果

type Personer interface {
	SayHello()
}

type Student struct {
}
type Teacher struct {
}

func (s *Student) SayHello() {
	fmt.Println("老师好")
}
func (s *Teacher) SayHello() {
	fmt.Println("同学好")
}

//多态的实现：以接口为参数
func WhoSayHi(h Personer) {
	h.SayHello()

}

func main() {
	var student Student
	var teacher Teacher

	WhoSayHi(&student)
	WhoSayHi(&teacher)

}
