package main

import "fmt"

type Student struct {
	id   int
	name string
	age  int
	addr string
}

func main() {
	//定义结构体切片
	var s []Student = []Student{
		Student{101, "张三", 18, "北京"},
		Student{102, "李四", 18, "北京"},
	}

	fmt.Println(s)
	fmt.Println(s[1])
	fmt.Println(s[1].addr)

	//循环遍历结构体切片
	for k, v := range s {
		fmt.Println(k)
		fmt.Println(v)
	}

	//使用append追加
	s = append(s, Student{103, "Test", 19, "append"})
	fmt.Println(s)
}
