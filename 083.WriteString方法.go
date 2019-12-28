package main

import (
	"fmt"
	"os"
)

//语法：file.WriteString(str)，参数为string类型
func main() {
	file, err := os.Create("D:/a.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	//写入数据
	n, err := file.WriteString("Hello World") //n:内容长度（空格也算）
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n) //结果为11
}
