package main

import "fmt"

//最好将结构体定义在函数外部
type Student struct {
	id   int
	name string
	age  int
	addr string
}

func main() {

	//定义结构体数组
	var arr [3]Student = [3]Student{
		Student{101, "张三", 18, "北京"},
		Student{102, "李四", 18, "北京"},
		Student{103, "王五", 18, "北京"}, //最后一条仍然要保留逗号
	}

	//修改结构体数组的数据
	arr[0].age = 23

	fmt.Println(arr)
	fmt.Println(arr[1])
	fmt.Println(arr[2].name)

}
