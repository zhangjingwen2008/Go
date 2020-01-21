package main

import (
	"github.com/astaxie/beego"
	_ "shanghaiyiqi/models"
	_ "shanghaiyiqi/routers" //报名前加下划线的意义：在main函数执行前，调用该包下的init()函数
)

func main() {

}
