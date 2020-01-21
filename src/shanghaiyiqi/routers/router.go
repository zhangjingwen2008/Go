package routers

import (
	"github.com/astaxie/beego"
	"shanghaiyiqi/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/register", &controllers.UserController{}, "get:ShowRegister;post:HandlePost")

	beego.Router("/login", &controllers.UserController{}, "get:ShowLogin;post:HandleLogin")

	//文章列表访问
	beego.Router("/showArticleList", &controllers.ArticleController{}, "get:ShowArticleList")
	//添加文章
	beego.Router("/addArticle", &controllers.ArticleController{}, "get:ShowAddArticle;post:HandleAddArticle")
	//编辑文章
	beego.Router("/updateArticle", &controllers.ArticleController{}, "get:ShowUpdateArticle;post:HandleUpdateArticle")
	//删除文章
	beego.Router("/deleteArticle", &controllers.ArticleController{}, "get:DeleteArticle")
	//添加分类
	beego.Router("/addType", &controllers.ArticleController{}, "get:ShowAddType;post:HandleAddType")

}
