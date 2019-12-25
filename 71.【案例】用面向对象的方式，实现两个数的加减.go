package main

import "fmt"

type Object struct {
	num1 int
	num2 int
}

//加操作
type Add struct {
	Object
}

func (p *Add) GetResult(num1 int, num2 int) int {
	p.num1 = num1
	p.num2 = num2
	return p.num1 + p.num2
}

//减操作
type Sub struct {
	Object
}

func (p *Sub) GetResult(num1 int, num2 int) int {
	p.num1 = num1
	p.num2 = num2
	return p.num1 - p.num2
}

func main() {
	var add Add
	var sub Sub
	a := add.GetResult(2, 6)
	s := sub.GetResult(3, 1)
	fmt.Println(a)
	fmt.Println(s)
}
