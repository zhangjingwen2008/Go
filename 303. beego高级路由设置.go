package main

/*
	beego高级路由设置

		- 设置路由方法 beego.Router(参数1, 参数2, 参数3)
			- 参数1：路由路径
			- 参数2：访问的Controller
			- 参数3：指定访问方式与对应调用方法

		- 参数3的多种设置方法：
			1.给请求指定自定义方法 一个请求指定一个方法
				beego.Router("/login",&controllers.LoginController{}, "get:ShowLogin;post:PostFunc")
			2.给多个请求指定一个方法
				beego.Router("/index",&controllers.IndexController{}, "get,post:HandleFunc")
			3.给所有请求指定一个方法
				beego.Router("/index",&controllers.IndexController{}, "*:HandleFunc")
			4.当两种指定方法冲突的时候：范围越小优先级越高。在这里访问post的时候，结果是访问PostFunc方法
				beego.Router("/index",&controllers.IndexController{}, "*:HandleFunc;post:PostFunc")

		- 使用例子:
			func init() {
				beego.Router("/login",&controllers.LoginController{}, "get:ShowLogin;post:PostFunc")
			}

*/
