package main

import (
	"fmt"
	"io"
	"os"
)

/*
作用：在内容指定位置写入
语法：file.WriteAt([]byte, int)   	参数1：插入内容；参数2：插入位置（插入位置在内容中间时，会产生覆盖）
指定写入位置：file.Seek(int, int)		参数1：光标往后面移动n个位置，即n个空格；参数2：设置为io.SeekEnd时，意为内容的最后位置
	io.SeekStart：文件起始位置
	io.SeekCurrent：文件当前位置
	io.SeekEnd：文件结尾位置
*/
func main() {
	//创建文件
	file, err := os.Create("D:/a.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	file.WriteString("Hello World")

	//获得写入位置
	num, _ := file.Seek(0, io.SeekEnd) //将光标定位到文件中原有内容的后面，返回文件中原有数据的长度
	fmt.Println("num=", num)

	//指定位置写入
	var str string = "嘎"
	n, err := file.WriteAt([]byte(str), num) //将string插入内容最后的位置
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)
}
