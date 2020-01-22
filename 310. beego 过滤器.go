package main

/*
	路由过滤器
		- 作用：可以根据 指定的匹配规则 在 特定的项目运行阶段 去 执行自定义函数，函数一般放在beego.router()之前
		- 场景：可实现登录判断等功能，特定页面只有登陆了才可访问
		- 使用格式：
			beego.InsertFilter("/article/*", position int, filter FinterFunc)
			- 参数1：路由匹配规则，支持正则表达式
			- 参数2：要执行在项目的哪个阶段，例如beego.BeforeExec
			- 参数3：自定义函数
			- beego项目运行过程中，框架帮我们分了五个阶段，分别是：
				1.BeforeStatic 静态地址之前
				2.BeforeRouter 寻找路由之前
				3.BeforeExec 找到路由之后，开始执行相应的 Controller 之前
				4.AfterExec 执行完 Controller 逻辑值后执行的过滤器
				5.FinishRouter 执行完逻辑值后执行的过滤器
*/

/*
	登录过滤器的例子

	func init() {
		//构建登录过滤器
		beego.InsertFilter("/article/*",beego.BeforeExec,Filter)

		//不需要过滤器即可访问的地址
		beego.Router("/", &controllers.MainController{})
		beego.Router("/register", &controllers.UserController{}, "get:ShowRegister;post:HandlePost")
		beego.Router("/login", &controllers.UserController{}, "get:ShowLogin;post:HandleLogin")

		//需要登陆过滤的地址，使用article前缀作为标识
		beego.Router("/article/showArticleList", &controllers.ArticleController{}, "get:ShowArticleList")
		beego.Router("/article/addArticle", &controllers.ArticleController{}, "get:ShowAddArticle;post:HandleAddArticle")
		beego.Router("/article/showArticleDetail", &controllers.ArticleController{}, "get:ShowArticleDetail")
		beego.Router("/article/updateArticle", &controllers.ArticleController{}, "get:ShowUpdateArticle;post:HandleUpdateArticle")
		beego.Router("/article/deleteArticle", &controllers.ArticleController{}, "get:DeleteArticle")
		beego.Router("/article/addType", &controllers.ArticleController{}, "get:ShowAddType;post:HandleAddType")
		beego.Router("/article/logout",&controllers.UserController{},"get:Logout")

	}

	//过滤器函数，以上下文环境context.Context作为参数，注意导包是beego框架的，若是go框架的则无法运作
	var Filter = func(ctx *context.Context) {
		userName :=ctx.Input.Session("userName")
		if userName==nil{
			ctx.Redirect(302,"/login")
			return
		}
	}
*/
