package main

import (
	"fmt"
	"runtime"
)

/*
return：返回当前函数调用到调用者那里去，return之前的defer都生效
Goexit()：结束调用该函数的当前go程。Goexit()之前注册的defer都生效
语法：runtime.Goexit()
*/
func main() {

	go func() {
		fmt.Println("aaaaaaaaaaaa")
		test()
		defer fmt.Println("bbbbbbbb")
	}()

	for {

	}

}

func test() {
	defer fmt.Println("cccccc")
	//return
	runtime.Goexit()
	fmt.Println("dddddddddddd")
}
