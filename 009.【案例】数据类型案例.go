package main

import "fmt"

func main() {
	var name string
	var age int
	fmt.Println("姓名:")
	fmt.Scan(&name)
	fmt.Println("年龄:")
	fmt.Scan(&age)

	fmt.Printf("您好：%s您的年龄是%d", name, age)
}
