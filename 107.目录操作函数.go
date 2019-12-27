package main

import (
	"fmt"
	"os"
)

/*
打开目录的语句：os.OpenFile(path,os.O_RDONLY,os.ModeDir)
注意：第3个参数os.ModeDir是关键

file.Readdir(n)：读取目录中的目录项，n为-1表示读取全部
*/
func main() {

	//获取用户输入的目录路径
	fmt.Println("输入待查询的目录：")
	var path string
	fmt.Scan(&path)

	//打开目录
	f, err := os.OpenFile(path, os.O_RDONLY, os.ModeDir)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	//读取目录项
	info, err := f.Readdir(-1) //-1表示读取所有目录项

	//遍历返回的切片
	for _, fileInfo := range info {
		if fileInfo.IsDir() { //是目录
			fmt.Println(fileInfo.Name(), "是一个目录")
		} else {
			fmt.Println(fileInfo.Name(), "是一个文件")
		}
	}

}
