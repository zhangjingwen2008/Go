package main

import "fmt"

func main() {
	nums := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	//创建数组指针
	var p *[10]int
	p = &nums

	fmt.Println(*p)      //获取整个数组中的内容
	fmt.Println(p[3])    //等价于(*p)[3]
	fmt.Println((*p)[3]) //*p需要加括号的原因：方括号[]优先级比星号*还高

	//在函数内修改数组指针内容
	UpdateArray(p)
	fmt.Println(*p)

}

func UpdateArray(p *[10]int) {
	p[0] = 111
}
