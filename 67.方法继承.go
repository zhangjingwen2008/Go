package main

import "fmt"

//父类和父函数
type Person struct {
	name string
	age  int
}

func (p *Person) PrintInfo() { //要被继承的父函数
	fmt.Println(*p)
}

type Student struct {
	Person
	score float64
}

func main() {
	stu := Student{Person{"Hunter", 20}, 99}
	stu.PrintInfo()
}
