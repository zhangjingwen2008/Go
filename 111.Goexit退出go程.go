package main

import (
	"fmt"
	"runtime"
)

/*
return：返回当前函数调用到调用者那里去，return之前的defer都生效
runtime.Goexit()：结束调用该函数的当前go程。Goexit()之前注册的defer都生效。
os.Exit(1)：结束当前进程。参数0：正常退出；参数非0：异常退出
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
