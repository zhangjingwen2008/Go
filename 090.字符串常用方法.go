package main

import (
	"fmt"
	"strings"
)

/*
常用字符串方法：
Contains()：判断一个字符串是否包含在另一个字符串中
Join()：字符串连接
Index()：在字符串中查找某个字串的位置。若不存在则返回-1
Repeat()：某个字符串重复多少次，返回的是重复后的字符串
Replace()：在字符串中，把旧字符串替换为新字符串，n表示替换次数，n<0表示全部替换
Split()：把字符串按照sep分隔，返回切片
*/
func main() {
	//1.Contains()
	var str1 string = "hellogo"
	b1 := strings.Contains(str1, "go")
	fmt.Println(b1)

	//2.Join()
	str2 := []string{"abd", "hello", "world"}
	s := strings.Join(str2, "-")
	fmt.Println(s)

	//3.Index()
	str3 := "hellogo"
	n := strings.Index(str3, "go4")
	fmt.Println(n)

	//4.Repeat()
	str4 := strings.Repeat("go", 3)
	fmt.Println(str4)

	//5.Replace
	str5 := strings.Replace("hello-world-go", "-", "|", -1)
	fmt.Println(str5)

	//6.Split()
	str6 := strings.Split("hello-world-go", "-")
	for _, value := range str6 {
		fmt.Println(value)
	}

}
