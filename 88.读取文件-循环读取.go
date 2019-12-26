package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//1.读取文件
	file, err := os.Open("D:/a.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	//2.循环读取
	buffer := make([]byte, 10) //固定每次读取10个字节的内容
	for {
		n, err := file.Read(buffer)
		if err == io.EOF { //读取到内容末尾时，就会有io.EOF的异常，可据此进行判断退出循环
			break
		}
		fmt.Print(string(buffer[:n]))
	}
}
