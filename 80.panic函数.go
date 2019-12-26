package main

import "fmt"

//panic(str)：str为异常内容，触发后后面语句不再执行
func main() {
	Test(11)
}

func Test(n int) {
	//触发异常方式1：手动
	fmt.Println("hhh")
	panic("Touch Error!!!") //结果：在此处抛出异常，后面的语句不会执行。一般不会由程序员手动抛出，而是系统自动调用
	fmt.Println("After Panic!!")

	//触发异常方式2：自动
	sort := [5]int{}
	sort[n] = 11
	fmt.Println(sort)
}
