package main

import (
	"github.com/astaxie/beego"
	"github.com/gomodule/redigo/redis"
)

/*

	GoLang操作Redis数据库

		连接数据库
			- conn, err := redis.Dial("tcp", ":6379")

		执行操作数据库语句————写入
			- conn.Send("set", "Liangzai", "hahaha")
			- conn.Flush()
			- rep,_:=conn.Receive()

		执行操作数据库语句————读取
			读取1个
				- rep,err:=conn.Do("get","liangzai")
				- str,err:=redis.String(rep, err)
				- beego.Info(str)
			读取多个，且数据类型不同
				- rep,err:=conn.Do("get","liangzai", "c1")			//一个为字符串，一个为整数
				- res,err:=redis.Values(rep,err)					//使用回复函数助手的Values方法接收多个不同类型数据
				- var str string
				- var num int
				- redis.Scan(res,&str,&num)							//使用redis.Scan保存接收到的数据。有限制，只能扫描常见类型，不可扫描结构体。若遇到结构体，则需要使用序列化与反序列化。
				- beego.Info(str,num)

*/

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//连接数据库
	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		c.Ctx.WriteString("连接redis数据库失败")
		return
	}
	defer conn.Close()

	//执行操作数据库语句
	//conn.Send("set","liangzai","哇哈哈")
	//conn.Flush()
	//rep,_:=conn.Receive()
	//rep, err:=conn.Do("get","liangzai")									//获取1个redis值，返回1个值
	rep, err := conn.Do("mget", "liangzai", "c1") //获取多个redis值，返回切片
	//回复助手函数
	res, err := redis.String(rep, err)
	if err != nil {
		c.Ctx.WriteString("获取内容出错")
		return
	}
	c.Ctx.WriteString(res)

}
