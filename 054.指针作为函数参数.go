package main

import "fmt"

func main() {
	num := 10
	Update(&num) //传入函数的参数为变量地址
	fmt.Println(num)
}

func Update(p *int) { //使用指针来接收参数
	*p = 22 //指针的修改就可以影响到原变量
}
