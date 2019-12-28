package main

import "fmt"

//大小写字母之间的ASCII编码相差32
func main() {
	var ch byte = 'a'
	fmt.Println(ch)      //输出结果97（ASCII编码）
	fmt.Printf("%c", ch) //输出结果a

}
