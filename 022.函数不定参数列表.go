package main

import "fmt"

/*
args...
for range
_（下划线）：匿名变量，不保存任何数据
*/
func main() {
	TestSum(6, 88, 234)
	TestSum2(10)
}

//固定实参放前面，不固定放后面（可以不传入，传入数量最少为0）
func TestSum2(num1 int, args ...int) {
}

func TestSum(args ...int) {
	fmt.Println("for循环：")
	for i := 0; i < len(args); i++ { //动态打印集合参数
		fmt.Println(args[i])
	}

	fmt.Println("for range循环：") //for range返回2种值：下标和值
	for index, num := range args {
		fmt.Println("index=", index)
		fmt.Println("num=", num)
	}

	fmt.Println("for range循环（不需要index）：") //对于不需要的数据，用下划线代替即可
	for _, num := range args {
		fmt.Println("num=", num)
	}
}
