package main

import (
	"github.com/astaxie/beego"
	_ "shanghaiyiqi/routers"
)

func main() {
	beego.AddFuncMap("prepage", ShowPrePage)
	beego.AddFuncMap("nextpage", ShowNextPage)
	beego.Run()
}

//后台定义一个函数
func ShowPrePage(pageIndex int) int {
	if pageIndex == 1 {
		return pageIndex
	}
	return pageIndex - 1
}

func ShowNextPage(pageIndex int, pageCount int) int {
	if pageIndex == pageCount {
		return pageIndex
	}
	return pageIndex + 1
}

/*
	视图函数
		- 作用：处理视图中简单业务逻辑
		- 步骤：
			1.创建后台函数
			2.在视图中定义函数名
			3.在beego.Run之前关联起来
*/
