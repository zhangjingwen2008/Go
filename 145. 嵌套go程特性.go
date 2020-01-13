package main

import (
	"fmt"
	"runtime"
	"time"
)

/*
结论：嵌套外层go程退出，其内层创建的go程仍然正常运行
应用：在文件144聊天室的退出功能中，return退出了当前go程后，仍需要手动关闭其内的子go程
本程序运行结果：打印完“----2”且触发return退出了外层go程后，其内层子go程AA()仍然在继续打印，验证了上面的结论
*/

func main() {
	go func() {
		fmt.Println("------------------ 1")
		go AA()
		fmt.Println("------------------ 2")
		return
	}()

	for {
		runtime.GC()
	}
}

func AA() {
	for {
		time.Sleep(time.Millisecond * 200)
		fmt.Println("-------")
	}
}
