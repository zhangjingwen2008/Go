package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

/*
	使用之前需要先安装数据库驱动：go get github.com/go-sql-driver/mysql
*/
func main() {

	//1.连接数据库
	//"用户名:密码@[连接方式](主机名:端口名)/数据库名"
	db, _ := sql.Open("mysql", "root:admin@(127.0.0.1:3306)/test") //设置连接数据库的参数
	defer db.Close()                                               //关闭数据库

	//2.正式连接
	err := db.Ping()
	if err != nil {
		fmt.Println("数据库连接失败")
		return
	}

	/*
		//操作一：执行数据操作语句 db.Exec()
		sql := "insert into user values (4,'Hunter')"
		result, _ := db.Exec(sql)     //执行SQL语句
		n, _ := result.RowsAffected() //获取受影响的记录数
		fmt.Println("受影响的记录数是：", n)

		//操作二：执行预处理 db.Prepare()
		user := [2][2]string{{"5", "liangzai"}, {"6", "wasai"}}
		stmt, _ := db.Prepare("insert into user values (?,?)")		//获取预处理语句对象
		for _,u :=range user{
			stmt.Exec(u[0],u[1])		//调用预处理语句
		}

		//操作三：单行查询 db.QueryRow()
		var id,name string
		rows:=db.QueryRow("select * from user")		//获取1行数据
		rows.Scan(&id,&name)		//将rows中的数据存到id和name中
		fmt.Println(id,"--",name)

		//操作四：多行查询 db.Query()
		rows, _ := db.Query("select * from user") //获取所有数据
		var id, name string
		for rows.Next() {			//循环显示所有数据
			rows.Scan(&id, &name)
			fmt.Println(id, "--", name)
		}
	*/
}
