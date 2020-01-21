package main

/*

	beego操作数据库：
		- beego.Info()
		- beego.Error()
		- 其他均与普通数据库操作一致

		func init() {
			conn,err:=sql.Open("mysql","root:admin@tcp(127.0.0.1:3306)/test")
			if err!=nil{
				beego.Info("连接错误", err)
				beego.Error("连接错误", err)
				return
			}
			defer conn.Close()

			//创建表
			_,err :=conn.Exec("create table itcast(name VARCHAR(40),password VARCHAR(40));")
			if err!=nil{
				beego.Error("创建表失败", err)
				beego.Info("创建表失败", err)
				return
			}

			//增加
			conn.Exec("insert into itcast(name,password) values(?,?)","Hunter", "Sally")

			//查询
			res,err:=conn.Query("select name from itcast")
			var name string
			for res.Next() {
				res.Scan(&name)
				beego.Info(name)
			}
		}

*/
