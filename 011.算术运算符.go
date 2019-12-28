package main

import "fmt"

func main() {
	var num1 int = 20
	var num2 int = 10

	fmt.Println("num1+num2=", num1+num2)
	fmt.Println("num1-num2=", num1-num2) //除数不能为0，否则报错
	fmt.Println("num1*num2=", num1*num2)
	fmt.Println("num1/num2=", num1/num2)
	fmt.Println("num1%num2=", num1%num2) //除数不能为0，否则报错

}
