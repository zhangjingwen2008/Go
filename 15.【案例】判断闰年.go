package main

import "fmt"

func main() {
	var year int

	fmt.Println("请输入年份：")
	fmt.Scan(&year)

	isRun := (year%400 == 0) || (year%4 == 0 && year%100 != 0)

	fmt.Printf("是否为闰年：%t", isRun)
}
