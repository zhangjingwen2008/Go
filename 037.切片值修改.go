package main

import "fmt"

//对新切片的值修改，会影响到原切片
//原因：截取后的新切片其实并没有分配新空间，仅仅是将新切片指向原切片
func main() {
	Slice := []int{3, 2, 6, 7}
	Slice1 := Slice[1:3]

	Slice1[0] = 22        //对新切片进行修改为22
	fmt.Println(Slice[1]) //结果为22，就是原切片也被修改
}
