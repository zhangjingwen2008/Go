package main

import "fmt"

/*
append 切片内追加数据
*/
func main() {
	var Slice []int
	Slice = append(Slice, 1, 2, 6) //append追加
	fmt.Println(Slice)
}
