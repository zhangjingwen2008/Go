package main

import "fmt"

func main() {
	t_shirt, trouser := 35, 120
	total := float64(t_shirt*3+trouser*2) * 0.88
	fmt.Printf("小明应该付%.1f元，整钱为%d", total, int(total)) //使用%.0f会四舍五入，转成int才不会
}
