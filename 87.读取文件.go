package main

import (
	"fmt"
	"os"
)

/*
file.Read([]byte)：读取文件，参数为接收读取内容的容器

os.OpenFile()：可指定打开文件模式与权限
os.Open()：不可指定，固定模式为O_RDONLY，权限为0
*/
func main() {
	//1.读取文件
	file, err := os.Open("D:/a.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	//2：进行文件内容读取
	buffer := make([]byte, 1024*2) //定义切片，用于存储从文件中读取的数据。大小为2KB
	n, err := file.Read(buffer)
	if err != nil {
		fmt.Println(err)
	}

	//3.打印内容
	fmt.Println(n)
	fmt.Println(buffer[:n])         //读取获得的原字节码，全是ASCII值
	fmt.Println(string(buffer[:n])) //读取获得的原内容，使用string转换
}
