package main

import (
	"fmt"
	"runtime"
)

/*
GO MAX PROCS作用：设置当前进程使用的最大CPU核数。返回结果，为上一次调用成功的设置值。首次返回，为默认值
语法：runtime.GOMAXPROCS(n)
*/
func main() {

	n := runtime.GOMAXPROCS(2)
	fmt.Println(n)

	for {
		fmt.Print(1) //主go程

		go fmt.Print(0) //字go程
	}

}
