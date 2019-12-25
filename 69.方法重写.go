package main

import "fmt"

type Person struct { //父类
	name   string
	age    int
	gender string
}

func (p *Person) PrintInfo() {
	fmt.Println("这是父类方法")
}

type Student struct { //子类
	Person
	score float64
}

func (s *Student) PrintInfo() {
	fmt.Println("这是子类方法")
}

func main() {
	var stu Student
	stu.PrintInfo()        //结果调用的是子类方法
	stu.Person.PrintInfo() //结果调用父类方法
}
