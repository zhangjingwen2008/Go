package main

/*
	ORM建表步骤：
		1.在模型层M，定义结构体和初始化方法
			//定义一个结构体
			type User struct{
				Id int
				Name string
				PassWord string
				//若定义为Pass_Word，则映射到数据库时会变成双下划线pass__word，而在ORM里双下划线__是由特殊含义的，所以尽量不要使用下划线的命名方式
			}

			func init() {
				//ORM操作数据库
				//获取连接对象
				orm.RegisterDataBase("default","mysql", "root:admin@tcp(127.0.0.1:3306)/test?charset=utf8")
				orm.RegisterModel(new(User))		//需要注册此User结构体

				//创建表
				orm.RegisterModel(new(User))

				//生成表
				//参数1：表别名；参数2：是否强制更新；参数3：是否显示详细信息
				orm.RunSyncdb("default",false, true)
			}

*/
