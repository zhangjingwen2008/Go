package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "3.14 123.123 haha 1.0 ab.3 123."

	//ret:=regexp.MustCompile(`[0-9]+\.[0-9]+`)
	ret := regexp.MustCompile(`\d+\.\d+`)
	all := ret.FindAllStringSubmatch(str, -1)

	fmt.Println(all)
	//匹配结果：[[3.14] [123.123] [1.0]]
}
