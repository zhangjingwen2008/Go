package main

import "fmt"

func main() {
	var Numbers [5]int = [5]int{2, 6, 4, 7, 1}

	GetNumbers(Numbers)

	for _, value := range Numbers { //发现在函数中修改数组的值，不会影响到原函数
		fmt.Println(value) //结果仍为2,6,4,7,1，没有修改后的10
	}
}

func GetNumbers(n [5]int) {
	n[3] = 10 //尝试对数组进行修改
}
