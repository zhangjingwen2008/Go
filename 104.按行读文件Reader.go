package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/*
NewReader(io.Reader)：为指定文件创建一个缓冲区
ReadBytes(byte)：读数据，以输入的byte为截止返回一行
*/
func main() {
	file, err := os.Open("D:/a.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	file.Close()

	//创建一个带有缓冲区的Reader
	reader := bufio.NewReader(file)
	for {
		buf, err := reader.ReadBytes('\n') //读一行数据
		if err != nil || err == io.EOF {
			fmt.Println("文件读取完毕")
			return
		} else if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(buf))
	}

}
