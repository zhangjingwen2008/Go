package main

import "fmt"

func main() {
	const PI = 3.14
	var radius float64

	fmt.Println("请输入半径")
	fmt.Scan(&radius)
	zhouchang := 2 * PI * radius
	area := 2 * PI * radius * radius

	fmt.Printf("圆的周长为：%.1f\n", zhouchang)
	fmt.Printf("圆的面积为：%.1f", area)

}
