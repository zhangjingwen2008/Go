package main

import "fmt"

func main() {

	chinese, math, english := 90, 89, 69

	total := chinese + math + english
	average := float64(chinese+math+english) / 3 //注意，需要在相除前把类型转换为浮点型，后面才不会报错

	fmt.Printf("总分:%d\n", total)
	fmt.Printf("品均分:%.1f", average)
}
