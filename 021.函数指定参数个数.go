package main

import "fmt"

func main() {
	show(2, "ff")
}

//函数
func show(num1 int, num2 string) {
	fmt.Printf("输入的内容为：%s，数值为:%d", num2, num1)
}
