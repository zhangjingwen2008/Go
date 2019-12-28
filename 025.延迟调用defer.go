package main

import "fmt"

/*
defer
*/
func main() {
	//1个defer
	defer fmt.Println("fun01") //后执行
	fmt.Println("fun02")       //先执行

	//多个defer（多个defer时，按LIFO后进先出的顺序执行）
	defer fmt.Println("fun03") //第3个执行
	defer fmt.Println("fun04") //第2个执行
	defer fmt.Println("fun05") //第1个打印

}
