package main

import "fmt"

func main() {
	//const PI float64= 3.14		//定义常量方法1
	const PI = 3.14 //定义常量方法2（注意，等号前面没有冒号）
	fmt.Println(PI)
	//fmt.Printf("%p\n",&PI)		//无法打印常量的地址，会报错

}
