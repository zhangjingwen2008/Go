package main

import (
	"fmt"
	"os"
)

/*
函数OpenFil()
作用：打开文件，进行操作

参数1：打开文件路径
参数2：打开模式
	O_RDONLY：只读模式
	O_WRONLY：只写模式
	O_RDWR：可读可写模式（光标在最开头）
	O_APPEND：追加模式（光标在内容最后）
参数3：权限，取值范围（0-7）
	0：无权限
	1：执行权限
	2：写权限
	3：写权限与执行权限
	4：读权限
	5：读权限与执行权限
	6：读权限与写权限
	7：读权限+写权限，与执行权限
*/
func main() {
	//打开文件
	file, err := os.OpenFile("D:/a.txt", os.O_APPEND, 6) //追加模式，不会覆盖
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	//写入/读取数据
	n, err := file.WriteString("fff")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)

}
