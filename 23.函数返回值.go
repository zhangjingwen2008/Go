package main

import "fmt"

/*
return
*/
func main() {
	var s int

	s = ReturnData(2) //接收单值
	s1, s2 := ReturnMultiData()

	fmt.Println(s)
	fmt.Println(s1)
	fmt.Println(s2)
}

//返回单值写法1
func ReturnData(num int) int { //最后一个int：指定返回的数据类型
	var sum int
	sum = num + 1
	return sum
}

//返回单值写法2
func ReturnData2(num int) (sum int) { //sum写在这里，等于提前声明了，后面就不用再声明
	//var sum int
	sum = num + 1
	return sum
}

//返回单值写法3
func ReturnData3(num int) (sum int) { //若已指定返回变量名称，则return可以不用再加
	//var sum int
	sum = num + 1
	return
}

//返回多值写法1
func ReturnMultiData() (int, int) {
	//var sum int
	num1 := 10
	num2 := 20
	return num1, num2
}
