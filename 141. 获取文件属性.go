package main

import (
	"fmt"
	"os"
)

/*
使用函数：os.Stat(fileName string)
*/

func main() {
	//通过参数获取文件
	list := os.Args
	if len(list) < 2 {
		fmt.Println("输入格式错误")
		return
	}

	//提取文件信息
	fileName := list[1]
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		fmt.Println("ERROR:", err)
	}
	fmt.Println("文件名称为:", fileInfo.Name())
	fmt.Println("文件大小为:", fileInfo.Size())

}
