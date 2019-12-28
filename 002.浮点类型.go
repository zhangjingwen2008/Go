package main

import "fmt"

func main() {
	var num float64
	var num2 float64
	num = 32.111
	num2 = 532.11
	fmt.Printf("%f\n", num)
	fmt.Printf("%.2f\n", num2) //保留2位小数

	num1 := 23.111
	fmt.Printf("%T", num1) //获得数据类型
}
