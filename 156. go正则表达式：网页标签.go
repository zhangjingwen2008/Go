package main

import (
	"fmt"
	"regexp"
)

/*
正则表达式(.*)：表示匹配随意次数的所有字符
	- 注意：随意字符表达式.的意外情况，不匹配换行符。即匹配内容出现换行的话，则无法捕获
【*?】：匹配多次，越少越好
【?s:】：单行模式，原来不能匹配换行符\n的点.符号，现在也可进行升级匹配了
【(?s:(.*?))】：可用于精准提取某一特征字串内的内容。例如可精确提取<link ref="wahaha"/>中ref属性的内容，正则表达式为`<link ref="(?s:(.*?))"`
*/

func main() {
	str := "<div>wahaha</div>"

	//1.编译正则表达式，这里以提取html中的<div>标签为例
	//ret:=regexp.MustCompile(`<div>(.*)</div>`)				//方法一，虽能匹配，但遇到换行符则无法识别
	ret := regexp.MustCompile(`<div>(?s:(.*?))</div>`) //方法二（推荐），使用单行模式?s:的方式，就能将换行符也包括在内

	//2.提取需要的信息
	all := ret.FindAllStringSubmatch(str, -1)
	//因为刚提取出来的数据是二维数组，所以还需要进行数据处理，才能获得真正匹配的数据
	for _, one := range all {
		fmt.Println(one[1])
	}

}
