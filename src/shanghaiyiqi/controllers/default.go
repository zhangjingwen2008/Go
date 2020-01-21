package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller //继承了beego的控制器
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["Hunter"] = "Sally" //Data[]：可以用于传递数据给视图，在视图层通过Hunter获取Sally这个值
	c.TplName = "test.html"    //TplName：指定视图文件
}

func (c *MainController) Post() {
	c.Data["data"] = "靓仔"
	c.TplName = "test.html"
}

func (c *MainController) ShowGet() {
	c.Data["liantzai"] = "Hunter"
	c.TplName = "test.html"
}
