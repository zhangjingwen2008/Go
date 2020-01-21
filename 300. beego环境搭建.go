package main

/*
	搭建环境：
		1.安装Beego源码
			- go get -u -v github.com/astaxie/beego
		2.安装Bee开发工具
			- go get -u -v github.com/beego/bee
		3.把bee的可执行文件bee.exe添加进环境变量即可

	创建、启动Beego项目：
		1.新建项目
			- bee new liangzai
		2.启动项目，先进入到要启动的项目目录下
			- bee run
		3.浏览即可，默认127.0.0.1:8080

	BeeGo项目目录：
		- conf：项目配置文件
		- controllers：MVC中的控制器
		- models：
		- routers：路由
		- static：静态文件，存放css、js、图片等资源
		- tests：
		- views：MVC中的视图

	其他命令：
		1.树的形式显示项目内容
			- tree
*/
