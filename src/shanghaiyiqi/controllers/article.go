package controllers

import (
	"bytes"
	"encoding/gob"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/gomodule/redigo/redis"
	"math"
	"path"
	"shanghaiyiqi/models"
	"time"
)

type ArticleController struct {
	beego.Controller
}

//展示文章列表页
func (this *ArticleController) ShowArticleList() {
	//session判断是否登录
	userName := this.GetSession("userName")
	if userName == nil {
		this.Redirect("/login", 302)
		return
	}

	//获取数据
	//高级查询
	//指定表
	o := orm.NewOrm()
	qs := o.QueryTable("Article") //o.QueryTable()查询表所有数据
	var articles []models.Article
	//_,err :=qs.All(&articles)
	//if err != nil{
	//	beego.Info("查询数据错误")
	//}

	//查询总记录数
	typeName := this.GetString("select")
	var count int64

	//获取总页数
	pageSize := 2

	//获取页码
	pageIndex, err := this.GetInt("pageIndex") //获得页码
	if err != nil {
		pageIndex = 1
	}
	//获取数据
	start := (pageIndex - 1) * pageSize //起始位置
	//获取数据库部数据。参数1：获取几条；参数2：从哪条数据开始获取。返回值还是querySetter
	//RelatedSel()：一对多时绑定一的那一端的对象，这样才能从“多”端调用“一”端
	qs.Limit(pageSize, start).RelatedSel("ArticleType").All(&articles)

	if typeName == "" {
		count, _ = qs.Count()
	} else {
		count, _ = qs.Limit(pageSize, start).RelatedSel("ArticleType").Filter("ArticleType__TypeName", typeName).Count()
	}
	pageCount := math.Ceil(float64(count) / float64(pageSize)) //天花板函数：两个浮点数相除，向上取整；地板函数，向下取整

	//获取文章类型
	var types []models.ArticleType

	conn, err := redis.Dial("tcp", ":6379")
	//判断：若redis中有则直接取；若没有则从数据库取，取完放入redis
	rep, err := conn.Do("get", "types")
	if rep == nil {
		o.QueryTable("ArticleType").All(&types)
		//redis存
		//针对从redis中无法存取结构体的问题，使用序列化和反序列化的方案
		var buffer bytes.Buffer
		enc := gob.NewEncoder(&buffer) //获取编码器
		enc.Encode(types)              //编码
		conn.Do("set", "types", buffer.Bytes())
	} else {
		//redis取
		data, _ := redis.Bytes(rep, err)
		dec := gob.NewDecoder(bytes.NewReader(data)) //解码
		dec.Decode(types)
	}

	this.Data["types"] = types

	//根据选中的类型查询相应类型文章
	//Filter()：过滤器，使用双下划线指定表下的特定字段
	if typeName == "" {
		qs.Limit(pageSize, start).RelatedSel("ArticleType").All(&articles)
	} else {
		qs.Limit(pageSize, start).RelatedSel("ArticleType").Filter("ArticleType__TypeName", typeName).All(&articles)
	}

	//把数据存入redis中
	if err != nil {
		beego.Info("redis连接失败")
		return
	}
	defer conn.Close()

	//操作数据
	conn.Do("set", "types", types)

	//展示数据
	this.Data["typeName"] = typeName
	this.Data["pageIndex"] = pageIndex
	this.Data["pageCount"] = int(pageCount)
	this.Data["count"] = count
	this.Data["articles"] = articles

	//指定试图布局
	//this.TplName = "index.html"
	this.Layout = "layout.html"
}

//展示添加文章页面
func (this *ArticleController) ShowAddArticle() {
	o := orm.NewOrm()
	var articles []models.Article
	o.QueryTable("Article").All(&articles)

	this.Data["articles"] = articles
	this.TplName = "add.html"
}

//获取添加文章数据
func (this *ArticleController) HandleAddArticle() {
	//1.获取数据
	articleName := this.GetString("articleName")
	content := this.GetString("content")
	beego.Info("获取数据完成")

	//2.校验数据
	if articleName == "" || content == "" {
		this.Data["errmsg"] = "添加数据不完整"
		this.TplName = "add.html"
		return
	}
	beego.Info("校验数据完成")

	//处理文件上传
	file, head, err := this.GetFile("uploadname") //.GetFile() 上传文件
	if err != nil {
		this.Data["errmsg"] = "文件上传失败"
		this.TplName = "add.html"
		return
	}
	defer file.Close()
	//this.SaveToFile("uploadname", "./static/img/"+head.Filename)
	//判断：文件大小
	if head.Size > 5000000 {
		this.Data["errmsg"] = "文件太大，重新上传"
		this.TplName = "add.html"
		return
	}
	//判断：文件格式
	ext := path.Ext(head.Filename) //path.Ext()：返回文件的后缀名。例如a.jpg，则返回.jpg
	if ext != ".jpg" && ext != ".png" && ext != ".jpeg" {
		this.Data["errmsg"] = "文件格式错误"
		this.TplName = "add.html"
		return
	}
	//判断：防止重名
	fileName := time.Now().Format("2006-01-02-15:04:05") + ext
	//存储
	this.SaveToFile("uploadname", "./static/img/"+fileName) //保存文件到指定路径
	beego.Info("上传文件完成 ", fileName)

	//3.处理数据
	//插入操作
	o := orm.NewOrm()
	var article models.Article
	article.ArtiName = articleName
	article.Acontent = content
	article.Aimg = "/static/img/" + fileName //这里的路径最前面不能有点. 与前面存储文件时不同
	//给文章添加类型
	typeName := this.GetString("select")
	var articleType models.ArticleType
	articleType.TypeName = typeName
	o.Read(&articleType, "TypeName")
	article.ArticleType = &articleType

	_, err1 := o.Insert(&article)
	if err != nil {
		this.Data["errmsg"] = err1
		this.TplName = "add.html"
		return
	}
	beego.Info("处理数据完成")

	//4.返回页面
	this.Redirect("/showArticleList", 302)

}

//展示文章详情页面
func (this *ArticleController) ShowArticleDetail() {
	id, er := this.GetInt("articleId")
	if er != nil {
		beego.Info("传递的连接错误")
	}
	o := orm.NewOrm()
	var article models.Article
	article.Id = id
	o.Read(&article)

	//Filter:过滤器，过滤条件为本表，Filter的第1个参数则不用指定表名，直接字段名即可
	//One:由于指定id进行查询，返回的肯定只有1条数据，所有不用All()而用one()
	o.QueryTable("Article").RelatedSel("ArticleType").Filter("Id", id).One(&article)

	//修改阅读量
	article.Acount += 1
	o.Update(&article)

	//多对多插入浏览记录
	m2m := o.QueryM2M(&article, "Users")
	userName := this.GetSession("userName")
	if userName == nil {
		this.Redirect("/login", 302)
		return
	}
	var user models.User
	user.Name = userName.(string)
	o.Read(&user, "Name")

	//多对多插入操作，Insert()不同
	m2m.Add(&user, "Name")

	//多对多查询
	//o.LoadRelated(&article, "Users")
	var users []models.User
	//Distinct()：去除重复
	o.QueryTable("User").Filter("Articles__Article__Id").Distinct().All(&users)

	this.Data["users"] = users
	this.Data["article"] = article
	//this.TplName="content.html"
	userLayout := this.GetSession("userName")
	this.Data["userName"] = userLayout.(string)
	//指定试图布局
	this.Layout = "layout.html"
}

//显示编辑頁面
func (this *ArticleController) ShowUpdateArticle() {
	//获取数据
	id, err := this.GetInt("articleId")

	//校验数据
	if err != nil {
		beego.Info("请求文章错误")
		return
	}

	//数据处理
	//查询相应文章
	o := orm.NewOrm()
	var article models.Article
	article.Id = id
	o.Read(&article)

	//返回视图
	this.Data["article"] = article
	this.TplName = "update.html"
}

//处理编辑页面
func (this *ArticleController) HandleUpdateArticle() {
	id, err := this.GetInt("articleId")
	articleName := this.GetString("articleName")
	content := this.GetString("content")
	filePath := UploadFile(&this.Controller, "uploadname")

	//数据校验
	if err != nil || articleName == "" || content == "" || filePath == "" {
		beego.Error("请求错误,", err)
		return
	}

	//数据处理
	o := orm.NewOrm()
	var article models.Article
	article.Id = id
	err = o.Read(&article)
	if err != nil {
		beego.Info("更新的文章不存在")
		return
	}
	article.ArtiName = articleName
	article.Acontent = content
	if filePath != "NoImg" {
		article.Aimg = filePath
	}
	o.Update(&article)

	//返回视图
	this.Redirect("/showArticleList", 302)

}

//删除文章处理
func (this *ArticleController) DeleteArticle() {
	id, err := this.GetInt("articleId")
	if err != nil {
		beego.Error("获取文章id失败")
		return
	}

	o := orm.NewOrm()
	var article models.Article
	article.Id = id
	o.Delete(&article)

	this.Redirect("/showArticleList", 302)
}

//展示添加类型页面
func (this *ArticleController) ShowAddType() {
	//查询
	o := orm.NewOrm()
	var types []models.ArticleType
	o.QueryTable("ArticleType").All(&types)

	this.Data["types"] = types
	this.TplName = "addType.html"
}

//处理添加类型数据
func (this *ArticleController) HandleAddType() {
	typeName := this.GetString("typeName")
	if typeName == "" {
		beego.Info("信息不完整，请重新输入")
		return
	}

	o := orm.NewOrm()
	var articleType models.ArticleType
	articleType.TypeName = typeName
	o.Insert(&articleType)

	//this.TplName="add.html"
	this.Redirect("/addType", 302)
}

//删除文章类型
func (this *ArticleController) DeleteType() {
	id, err := this.GetInt("articleId")
	if err != nil {
		beego.Info("删除类型错误", err)
		return
	}

	o := orm.NewOrm()
	var article models.Article
	article.Id = id
	o.Delete(&article)

	this.Redirect("/article/articleAddType", 302)
}

//封装上传文件函数
//形参：为了通用，一般定义原参数的父类
func UploadFile(this *beego.Controller, filePath string) string {
	//处理文件上传
	file, head, err := this.GetFile(filePath) //.GetFile() 上传文件
	if err != nil {
		this.Data["errmsg"] = "文件上传失败"
		this.TplName = "add.html"
		return ""
	}
	defer file.Close()

	//判断：文件大小
	if head.Size > 5000000 {
		this.Data["errmsg"] = "文件太大，重新上传"
		this.TplName = "add.html"
		return ""
	}

	//判断：文件格式
	ext := path.Ext(head.Filename) //path.Ext()：返回文件的后缀名。例如a.jpg，则返回.jpg
	if ext != ".jpg" && ext != ".png" && ext != ".jpeg" {
		this.Data["errmsg"] = "文件格式错误"
		this.TplName = "add.html"
		return ""
	}

	//判断：防止重名
	fileName := time.Now().Format("2006-01-02-15:04:05") + ext
	//存储
	this.SaveToFile(filePath, "./static/img/"+fileName) //保存文件到指定路径

	return "/static/img/" + fileName
}
