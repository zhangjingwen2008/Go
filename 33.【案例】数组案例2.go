package main

import "fmt"

//输出数组中长度最长的字符串
func main() {
	str := [...]string{"马龙", "迈克尔乔丹", "雷吉米勒aaaaaa", "邓肯哇", "科比布莱恩特"}
	var MaxStr string = str[0]

	for i := 1; i < len(str); i++ {
		if len(str[i]) > len(MaxStr) {
			MaxStr = str[i]
		}
	}

	fmt.Printf(MaxStr)
}
