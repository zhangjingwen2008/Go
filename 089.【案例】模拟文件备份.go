package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//1.打开原文件
	file1, err := os.Open("D:/a.txt")
	if err != nil {
		fmt.Print(err)
	}

	//2.创建新文件
	file2, err := os.Create("D:/b.txt")
	if err != nil {
		fmt.Print(err)
	}

	defer file1.Close()
	defer file2.Close()

	//3.将源文件内容写入到新文件
	buffer := make([]byte, 1024*2)
	for {
		n, err := file1.Read(buffer)
		if err == io.EOF {
			break
		}
		file2.Write(buffer[:n])
	}
}
