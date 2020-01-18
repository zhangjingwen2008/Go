package main

import "fmt"

/*
通过类型断言，可以判断空接口中存储的数据类型

语法:
value, ok := m.(T)

m：空接口类型变量
T：断言类型
value：变量m中的值
ok：bool变量，若断言成功则为true
- 也可利用断言，将获得的空接口变量，转换为指定数据类型
*/

func main() {
	var i interface{}
	i = "fasf"

	value, ok := i.(string) //类型断言

	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("类型推断错误")
	}
}
