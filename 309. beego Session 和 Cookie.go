package main

/*
	Session：在服务器端存储
	Cookie：在客户端存储

	Beego中使用Cookie：
		 - 存Cookie
			this.Ctx.SetCookie(key, value, time)
			参数1：Cookie的key值
			参数2：Cookie的value值
			参数3：Cookie的有效时间
		- 取Cookie
			this.Ctx.GetCookie(key)
			参数：Cookie的key值
			返回：对应的value值。没有对应或已失效的Cookie，返回空字符串
		- 删Cookie
			this.Ctx.SetCookie(key, value, 0)
			参数1：Cookie的key值
			参数2：任意值
			参数3：设置为0，就立马失效

	Beego中使用Session：
		 - 要先在配置文件中进行设置：sessionon = true
		 - Beego设置Session一般不设置时间，浏览器关闭，Session失效
		 - 存Session
			this.SetSession(key, value)
			参数1：Session的key值
			参数2：Session的value值
		- 取Session
			this.GetSession(key)
			参数：Session的key值
			返回：对应的value值，类型是interface{}。
			interface{}返回的若为空时，值为nil，而不是空字符串""
		- 删Session
			this.DelSession(key)
			参数1：Session的key值
*/
