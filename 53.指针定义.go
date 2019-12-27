package main

import "fmt"

func main() {
	var a int = 10
	var p *int = &a //将指针p指向变量a。称为“解引用”、“间接引用”

	fmt.Println(p)  //变量a的内存地址
	fmt.Println(&p) //变量p的内存地址
	fmt.Println(*p) //变量p所指向地址的值

	*p = 222
	fmt.Println(a) //修改p，原地址a被影响

}
