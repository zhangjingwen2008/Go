package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"shanghaiyiqi/models"
)

type UserController struct {
	beego.Controller
}

//显示注册页面
func (this *UserController) ShowRegister() {
	this.TplName = "register.html"
}

//处理注册数据
func (this *UserController) HandlePost() {
	//1.获取数据
	userName := this.GetString("userName")
	pwd := this.GetString("password")
	beego.Info("验证成功")

	//2.校验数据
	if userName == "" || pwd == "" {
		beego.Info("注册数据不完整，请重新注册")
		this.Data["errmsg"] = "注册数据不完整，请重新注册"
		this.TplName = "register.html"
		return
	}
	beego.Info("校验成功")

	//3.操作数据
	//获取ORM对象
	o := orm.NewOrm()
	//获取插入对象
	var user models.User
	//给插入对象赋值
	user.Name = userName
	user.Password = pwd
	o.Insert(&user)
	beego.Info("插入成功")

	//4.返回页面
	//this.Ctx.WriteString("注册成功")
	//this.TplName="login.html"				//用此方法转发出现的问题，就是当跳转的时候虽然内容已经变了，但URL仍然是原来的地址
	this.Redirect("/login", 302) //重新发起请求，但无法用c.Data来传递数据
}

//展示登录页面
func (this *UserController) ShowLogin() {
	this.TplName = "login.html"
}

func (this *UserController) HandleLogin() {
	//1.获取数据
	userName := this.GetString("userName")
	pwd := this.GetString("password")

	//2.校验数据
	if userName == "" || pwd == "" {
		this.Data["errmsg"] = "登录数据不完整"
		this.TplName = "login.html"
		return
	}

	//3.操作数据
	//3.1.获取ORM对象
	o := orm.NewOrm()
	var user models.User
	user.Name = userName
	err := o.Read(&user, "Name")
	if err != nil {
		this.Data["errmsg"] = "用户不存在"
		this.TplName = "login.html"
		return
	}
	if user.Password != pwd {
		this.Data["errmsg"] = "密码错误"
		this.TplName = "login.html"
		return
	}

	//4.返回页面
	//this.Ctx.WriteString("登陆成功")
	this.Redirect("/showArticleList", 302)
}
