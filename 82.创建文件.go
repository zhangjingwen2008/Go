package main

import (
	"fmt"
	"os"
)

//创建文件语句：os.Create()
func main() {
	//创建文件
	file, err := os.Create("D:/a.txt")

	//判断是否出现异常
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close() //操作完文件后需要关闭，无论途中是否出错，因为会占用内存

	//对文件的操作

	//关闭

}
