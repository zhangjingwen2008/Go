package main

import "fmt"

/*
空接口：没有任何方法类型的接口，可以接收任何数据
*/
func main() {
	//空接口1：变量
	var i interface{}
	i = 123
	i = "abc"
	fmt.Println(i)

	//空接口2：切片
	var s []interface{}
	s = append(s, 23, "Test", 0.2) //可以放任何类型的数据
	for _, value := range s {
		fmt.Println(value)
	}

}
