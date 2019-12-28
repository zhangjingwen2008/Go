package main

import "fmt"

func main() {
	var count int
	fmt.Print("请输入要排序的数字个数:")
	fmt.Scan(&count)

	Slice := make([]int, count)
	Initial(Slice) //初始化
	Sort(Slice)    //排序

	fmt.Print("排序结果:", Slice)

}

func Sort(Slice []int) {
	var temp int
	for i := 0; i < len(Slice)-1; i++ {
		for j := 0; j < len(Slice)-1-i; j++ {
			if Slice[j] > Slice[j+1] {
				temp = Slice[j]
				Slice[j] = Slice[j+1]
				Slice[j+1] = temp
			}
		}
	}
}

func Initial(Slice []int) {
	var temp int
	for i := 0; i < len(Slice); i++ {
		fmt.Printf("第%d个数字：", i+1)
		fmt.Scan(&temp)
		Slice[i] = temp
	}
}
