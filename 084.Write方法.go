package main

import (
	"fmt"
	"os"
)

//语法：file.Write([]byte])，参数为byte类型的切片
func main() {
	file, err := os.Create("D:/a.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	//使用Write写入数据
	var str string = "Hello World"
	n, err := file.Write([]byte(str)) //将string转换成byte的切片
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)
}
