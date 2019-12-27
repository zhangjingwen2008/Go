package main

import "fmt"

func main() {
	str := []string{"red", "", "blue", "", "", "green", "red", "", "blue"}
	out := Test(str)
	fmt.Println(out)
}

func Test(str []string) []string {
	out := str[:1] //存入第1个字符串
	for _, value := range str {
		i := 0
		for ; i < len(out); i++ {
			if out[i] == value {
				break
			}
			if i == len(out)-1 { //当新切片全部检索完后，再执行插入
				out = append(out, value)
			}
		}
	}
	return out
}
