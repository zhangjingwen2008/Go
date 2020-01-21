package main

/*
	beego调用post请求，执行顺序：
		1.外部访问
			- method指定为post
			- action内容，要与指定router的rootpath一致
			<form method="post" action="/">
				  <input type="submit"/>
			</form>
		2.路由
			- 第一个参数为路由路径
			- 第二个参数为调用的控制层
			func init() {
				beego.Router("/", &controllers.MainController{})
			}
		3.控制层C
			- 定义post方法
			- 从模型数据层M调用数据
			- Data：设置传输的数据
			- TplName：设置数据输出的视图层位置
			func (c *MainController) Post() {
				c.Data["data"]="靓仔"
				c.TplName= "test.html"
			}
		4.视图层V
			- 使用双大括号来渲染传输来的数据
			- 调用的名称，对应控制层C里的Data设置的名称
			<body>
				{{ .data }}
			</body>
*/
