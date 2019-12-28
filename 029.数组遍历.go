package main

import "fmt"

/*
for...len()
for range
*/

func main() {
	var Numbers [5]int = [5]int{2, 6, 4, 7, 1}

	//方法1：for...len()
	for i := 0; i < len(Numbers); i++ {
		fmt.Println(Numbers[i])
	}

	//方法2：for range
	for index, value := range Numbers {
		fmt.Printf("下标：%d,值为%d\n", index, value)
	}
}
