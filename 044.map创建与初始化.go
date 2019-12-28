package main

import "fmt"

//注：在函数中修改map的值，会影响到原map
func main() {
	//创建方式1
	var map1 map[int]string = map[int]string{1: "test", 3: "ttttt"}

	//创建方式2
	map2 := map[int]string{1: "test", 3: "ttttt"}

	//创建方式3
	map3 := make(map[int]string, 10)
	map3[1] = "test"
	map3[4] = "TTTT"

	fmt.Println(map1)
	fmt.Println(map2)
	fmt.Println(map3)
}
