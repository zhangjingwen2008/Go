package main

import "fmt"

func main() {
	var score int
	fmt.Println("Enter Score:")
	fmt.Scan(&score)

	switch {
	case score >= 90:
		fallthrough //fallthrough：继续往下执行
	case score >= 80:
		fallthrough
	case score >= 70:
		fmt.Println("A")
	case score >= 60:
		fmt.Println("D")
	default:
		fmt.Println("E")
	}
}
