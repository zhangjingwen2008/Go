package main

import (
	"errors"
	"fmt"
)

func main() {
	num, err := TestError(10, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(num)
	}
}

func TestError(num1 int, num2 int) (result int, err error) {
	if num2 == 0 {
		err = errors.New("除数不能为0") //定义指定异常
		return
	}
	result = num1 / num2
	return
}
