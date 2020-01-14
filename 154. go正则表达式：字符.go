package main

import (
	"fmt"
	"regexp"
)

/*
go语言使用正则表达式步骤：
	1.编译正则表达式：regexp.MustCompile([正则表达式])
	2.提取信息：ret.FindAllStringSubmatch([需要匹配的原内容], [匹配次数，-1表示全局匹配])
*/
func main() {
	str := "abc a7c mfc cat 8ca azc cba"

	//1.解析、编译正则表达式
	ret := regexp.MustCompile(`a.c`) //反引号`：表示使用原生字符串；句号.：匹配一个任意字符

	//2.提取需要的信息
	all := ret.FindAllStringSubmatch(str, -1)

	fmt.Println("alls:", all)
	//运行结果：
	//alls: [[abc] [a7c] [azc]]
}
