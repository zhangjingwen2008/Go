package main

import "fmt"

func main() {
	var a int = 10

	var p *int //一级指针
	p = &a
	fmt.Println(*p)

	var pp **int //二级指针
	pp = &p
	fmt.Println(**pp)

	**pp = 222 //对二级指针修改
	fmt.Println(a)

}
