package main

import "fmt"

func main() {
	var chinese int
	var math int

	fmt.Println("Enter chinese")
	fmt.Scan(&chinese)
	fmt.Println("Enter math")
	fmt.Scan(&math)

	if chinese > 70 && math == 100 { //相比Java，Go里if的判断区域没有括号
		fmt.Println("Prize 1000$")
	}
}
