package main

import "fmt"

func main() {
	var result int
	for i := 1; i <= 100; i++ { //for循环
		if i%2 == 0 {
			result += i
		}
	}
	fmt.Printf("1-100偶数之和为:%d", result)
}
