package main

/*
	beego 级联删除：

		- beego默认执行级联删除 可设置

		- 设置对应 rel 关系删除 可选的关系字段
			1.cascade 级联删除（默认值）
			2.set_null 非级联删除
			3.set_default 设置为默认值
			4.do_nothing 什么也不做，忽略

		- 示例：
			type Article struct {
				Id       int       `orm:"pk;auto"`
				ArtiName string    `orm:"size(20)"`
				Atime    time.Time `orm:"auto_now"`
				Acount   int       `orm:"default(0);null"`
				Acontent string    `orm:"size(500)"`
				Aimg     string    `orm:"size(100)"`

				ArticleType *ArticleType `orm:"rel(fk);on_delete(set_null);null"`  //on_delete设置级联删除
				Users       []*User      `orm:"rel(m2m)"`
			}
*/
