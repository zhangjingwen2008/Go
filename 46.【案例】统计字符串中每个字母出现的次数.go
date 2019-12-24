package main

import "fmt"

//使用char类型接收后，处理完的最终打印时需要用%c来格式化，否则会打印ASCII代码
func main() {
	var str string = "helloworld"
	map1 := make(map[byte]int)

	for i := 0; i < len(str); i++ {
		ch := str[i]
		map1[ch] = map1[ch] + 1
	}

	for key, value := range map1 {
		fmt.Printf("%c：%d\n", key, value)
	}

}
