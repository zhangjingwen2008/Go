package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

//表的设计

//定义一个结构体
type User struct {
	Id       int
	Name     string
	Password string
	//若定义为Pass_Word，则映射到数据库时会变成双下划线pass__word，而在ORM里双下划线__是由特殊含义的，所以尽量不要使用下划线的命名方式

	Articles []*Article `orm:"reverse(many)"` //多对多：用reverse(many)标识
}

type Article struct {
	Id       int       `orm:"pk;auto"`
	ArtiName string    `orm:"size(20)"`
	Atime    time.Time `orm:"auto_now"`
	Acount   int       `orm:"default(0);null"`
	Acontent string    `orm:"size(500)"`
	Aimg     string    `orm:"size(100)"`

	ArticleType *ArticleType `orm:"rel(fk)"`  //一对多：多的那一端，用rel(fk)标志位外键
	Users       []*User      `orm:"rel(m2m)"` //多对多：用rel(m2m)标识
}

//类型表
type ArticleType struct {
	Id       int
	TypeName string `orm:"size(20)"`

	Articles []*Article `orm:"reverse(many)"` //一对多，一的那一端，用reverse(many)标志和切片来接收
}

func init() {
	//ORM操作数据库
	//获取连接对象
	orm.RegisterDataBase("default", "mysql", "root:admin@tcp(127.0.0.1:3306)/test?charset=utf8")
	//创建表，绑定
	orm.RegisterModel(new(User), new(Article), new(ArticleType))

	//生成表
	//参数1：表别名；参数2：是否强制更新；参数3：是否显示详细信息
	orm.RunSyncdb("default", false, true)

}
