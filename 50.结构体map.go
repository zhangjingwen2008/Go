package main

import "fmt"

type Student struct {
	id   int
	name string
	age  int
	addr string
}

func main() {
	m := make(map[int]Student)
	m[1] = Student{101, "张三", 18, "北京"}
	m[2] = Student{102, "李四", 18, "北京"}

	fmt.Println(m)
	fmt.Println(m[1])
	fmt.Println(m[2].id)

	delete(m, 1)
	fmt.Println(m)

}
