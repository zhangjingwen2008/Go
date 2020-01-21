package main

/*
	ORM查询数据步骤：

		1.在模型层M，定义结构体和数据库连接
			type User struct{
				Id int
				Name string
				PassWord string
			}
			func init() {
				orm.RegisterDataBase("default","mysql", "root:admin@tcp(127.0.0.1:3306)/test?charset=utf8")
				orm.RegisterModel(new(User))		//需要注册此User结构体

			}

		2.在控制层C，调用ORM相应方法
			func (c *MainController) ShowGet()  {

				//获取ORM对象
				o:=orm.NewOrm()

				//插入操作
				var user models.User
				user.Name= "Hunter"
				user.Password="liangzai"
				count,err :=o.Insert(&user)		//插入语句

				//查询操作
				var user models.User
				user.Id= 1
				err:=o.Read(&user, "Id")		//查询语句

				//更新操作
				var user models.User			//先查询数据
				user.Id= 1
				err:=o.Read(&user)
				user.Name= "hahaha"				//后更新数据
				n,err :=o.Update(&user, "name")

				//删除操作
				var user models.User
				user.Id= 1						//指定id才可删除，其他字段不可用于删除
				count,err :=o.Delete(&user)		//删除语句

			}

*/
