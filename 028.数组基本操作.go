package main

import "fmt"

//Numbers[5] int

func main() {
	//1.全部赋值
	var Numbers1 [5]int = [5]int{2, 1, 6, 3, 4}
	fmt.Println(Numbers1[2]) //结果为6

	//2.部分赋值
	var Numbers2 [5]int = [5]int{32, 5} //未指定参数的数据，都为0
	fmt.Println(Numbers2[1])            //结果为5

	//3.指定元素初始化
	var Numbers3 [5]int = [5]int{2: 4, 4: 88} //指定位置为数组下标
	fmt.Println(Numbers3[4])                  //结果为88

	//4.动态初始化
	Numbers4 := [...]int{5, 2, 1, 6, 7, 3} //赋值三个点...可以赋任意数量的值
	fmt.Println(len(Numbers4))             //长度结果为6

	//5.外部初始化（循环赋值为例）
	var Numbers5 [5]int
	for i := 0; i < len(Numbers5); i++ {
		Numbers5[i] = i + 2
	}
	fmt.Println(Numbers5[3]) //结果为5
}
