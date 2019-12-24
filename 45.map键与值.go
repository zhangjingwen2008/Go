package main

import "fmt"

func main() {
	//1.通过key获取值
	var map1 map[int]string = map[int]string{1: "test", 3: "ttttt"}
	fmt.Println(map1[3])

	//2.通过key判断是否存在。存在，则返回value和true，反之
	value, ok := map1[1]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("ERROR")
	}

	//3.使用for遍历
	for key, value := range map1 {
		fmt.Println(key)
		fmt.Println(value)
	}

	//4.删除键值
	delete(map1, 3)
	fmt.Println(map1)
}
