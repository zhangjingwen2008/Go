package main

import (
	"fmt"
	"os"
)

/*
命令行参数：在命令行进行go文件运行时，可在最后以空格为分隔加入参数

使用：os.Args
返回数据格式：切片
命令行输入方式：D:\Project\Go>go run "140. 命令行参数.go" [第一个参数] [第二个参数] [……]
	xxx.go 第0个参数
	[第一个参数] 第1个参数
	[第二个参数] 第2个参数
*/
func main() {
	list := os.Args
	fmt.Println("list的内容为：", list)
}
