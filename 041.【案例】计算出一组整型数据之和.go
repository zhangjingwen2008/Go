package main

import "fmt"

func main() {
	var count int
	fmt.Print("请输入要求和的数据个数：")
	fmt.Scan(&count)

	Slice := make([]int, count)
	Initial(Slice)
	sum := Sum(Slice)

	fmt.Println("总和为:", sum)

}

func Initial(Slice []int) {
	var Num int
	for i := 0; i < len(Slice); i++ {
		fmt.Printf("请输入第%d个数字：", i+1)
		fmt.Scan(&Num)
		Slice[i] = Num
	}
}

func Sum(Slice []int) (sum int) {
	for i := 0; i < len(Slice); i++ {
		sum += Slice[i]
	}
	return
}
