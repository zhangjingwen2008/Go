package main

import "fmt"

type Person struct { //父类
	name   string
	age    int
	gender string
}

func (p *Person) PrintInfo() {
	fmt.Println(*p)
}

func main() {
	var person Person = Person{"Test", 23, "man"}

	//方法值
	f := person.PrintInfo
	fmt.Printf("%T\n", f) //结果为func()，说明类型为函数
	f()                   //结果为{Test 23 man}

	//方法表达式
	ff := (*Person).PrintInfo
	ff(&person)

}
