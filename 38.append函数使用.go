package main

import "fmt"

//切片扩容：当元素数量超出设定值时，会按容量*2的方式扩容
//但是，若容量超过1024字节，则每次扩容量为上一次的1/4
func main() {
	Slice := make([]int, 2, 3)
	Slice[0] = 22
	Slice[1] = 33
	Slice = append(Slice, 44)
	Slice = append(Slice, 55)
	Slice = append(Slice, 66)

	fmt.Println(Slice)      //结果[22 33 44 55 66]
	fmt.Println(len(Slice)) //结果5
	fmt.Println(cap(Slice)) //结果6（超过容量后，会X2扩容）
}
