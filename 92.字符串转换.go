package main

import (
	"fmt"
	"strconv"
)

func main() {
	//1.bool-->string
	str := strconv.FormatBool(true)
	fmt.Println(str)

	//2.string-->bool
	boolStr, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(boolStr)

	//3.int-->string
	str2 := strconv.Itoa(123)
	fmt.Println(str2)

	//4.string-->int
	intStr, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(intStr)

}
