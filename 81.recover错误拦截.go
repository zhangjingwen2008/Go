package main

import "fmt"

//recover()：错误拦截
func main() {
	Test(11)
}

func Test(n int) {
	defer TestRecover() //recover()必须配合defer才可生效
	sort := [5]int{}
	num := 6
	sort[num] = 2
	fmt.Println("aaa")
}

func TestRecover() {
	recover()
	//fmt.Println(recover())		也可对recover()进行打印，打印结果为异常内容
}
