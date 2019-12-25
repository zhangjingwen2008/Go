package main

import "fmt"

//父类
type Person struct {
	id   int
	name string
	age  int
}

//子类：Studeng和Teacher
type Student struct {
	Person //继承Person：这是匿名字段的方式，只有类型，没有成员名字
	score  float64
}
type Teacher struct {
	Person
	salary float64
}

func main() {

	//全部初始化
	var stu1 Student = Student{Person{101, "Test", 32}, 96}
	fmt.Println(stu1)

	//部分初始化1（仅子类字段）
	var stu2 Student = Student{score: 22}
	fmt.Println(stu2)

	//部分初始化2（仅父类字段）
	var stu3 Student = Student{Person: Person{id: 102}}
	fmt.Println(stu3)
}
