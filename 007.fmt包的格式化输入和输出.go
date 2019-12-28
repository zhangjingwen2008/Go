package main

import "fmt"

func main() {
	num := 11
	fmt.Printf("%b\n", num) //%b：10进制转成2进制
	fmt.Printf("%o\n", num) //%o：10进制转成8进制
	fmt.Printf("%x\n", num) //%x：10进制转成16进制
	fmt.Printf("%X\n", num) //%X：10进制转成16进制（大写）
}
