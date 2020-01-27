package main

import (
	"fmt"
	"strings"
)

func main() {
	str1:=[]string{"hello","world","!"}
	res:=strings.Join(str1, "+")
	fmt.Println(res)
}