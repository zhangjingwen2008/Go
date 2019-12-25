package main

import "fmt"

type Object struct {
	id int
}
type Person struct {
	Object //继承Object
	name   string
	age    int
}

type Student struct {
	Person //继承Person
	score  float64
}

func main() {
	var stu Student
	fmt.Println(stu) //结果为：{{{0}  0} 0}

	stu.id = 33
	fmt.Println(stu) //结果为：{{{33}  0} 0}
}
