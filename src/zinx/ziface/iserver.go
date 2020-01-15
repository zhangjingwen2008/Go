package ziface

/*
抽象层
*/

//定义一个服务器接口
type IServer interface {
	//启动服务器
	Start()
	//停止服务器
	Stop()
	//运行服务器
	Serve()
	//路由功能：给当前的服务注册一个路由方法，供客户端的连接处理使用
	AddRouter(msgID uint32, router IRouter)
	//获取当前server的连接管理器
	GetConnMgr() IConnManager

	//注册OnConnStart 钩子函数方法
	SetOnConnStart(func(connection IConnection))
	//注册OnConnStop 钩子函数方法
	SetOnConnStop(func(connection IConnection))
	//调用OnConnStart 钩子函数方法
	CallOnConnStart(connection IConnection)
	//调用OnConnStop 钩子函数方法
	CallOnConnStop(connection IConnection)
}
