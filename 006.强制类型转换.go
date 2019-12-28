package main

import "fmt"

func main() {
	var int_num int = 1
	var float_num float64 = 23.11
	fmt.Printf("%f\n", float64(int_num)) //int转float64
	fmt.Printf("%d\n", int(float_num))   //float64转int

	var num1 float32 = 3.22
	var num2 float64 = 32.66
	fmt.Printf("%.4f\n", float64(num1)+num2) //需要将float32转为同为float64才可相加，否则报错

	var num3 int = 1234
	fmt.Printf("%d", int8(num3)) //结果为-46. int8的范围是-128~127,所以高精度往低精度转换导致结果溢出
}
