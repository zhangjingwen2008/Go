package main

import "fmt"

func main() {
	var name string = "itcast" //字符串中隐藏了一个结束标志\0，打印结果是\0前面的内容
	fmt.Printf("%s", name)

	var name2 string = "靓仔"
	fmt.Println(len(name2)) //结果为6。Go语言中，一个汉字占3个字符

}
