package main

import "fmt"

func main() {
	s := []int{1, 2, 4, 5, 6}

	var p *[]int //创建切片指针
	p = &s

	fmt.Println(*p)      //结果为切片内容
	fmt.Println((*p)[2]) //结果为切片下标为2的数据

	(*p)[0] = 222 //对切片指针的指定数据进行修改
	fmt.Println(*p)

	//循环遍历1
	for i := 0; i < len(*p); i++ {
		fmt.Println((*p)[i])
	}
	//循环遍历2
	for key, value := range *p {
		fmt.Println(key)
		fmt.Println(value)
	}
}
