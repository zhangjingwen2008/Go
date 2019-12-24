package main

import "fmt"

func main() {
	var p [2]*int
	var i = 10
	var j = 20
	p[0] = &i
	p[1] = &j

	fmt.Println(p)     //结果为两个内存地址
	fmt.Println(*p[0]) //结果为10，这里不用给*p加括号，因为需要先计算p[0]以获得内存地址

	//循环遍历指针数组1
	for i := 0; i < len(p); i++ {
		fmt.Println(*p[i])
	}
	//循环遍历指针数组2
	for key, value := range p {
		fmt.Println(key)
		fmt.Println(*value)
	}
}
