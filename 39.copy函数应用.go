package main

import "fmt"

//copy(目的地，复制源)
//拷贝长度为两个切片中长度最小的值
func main() {
	var Slice1 []int = []int{2, 3}
	var Slice2 []int = []int{22, 33, 44, 55, 66}

	copy(Slice1, Slice2)

	fmt.Println(Slice1)
	fmt.Println(Slice2)

}
