package ziface

import "net"

//定义连接模块的抽象层
type IConnection interface {
	//启动连接 让当前的连接准备开始工作
	Start()

	//停止连接 结束当前连接的工作
	Stop()

	//获取当前连接的绑定socket conn
	GetTCPConnection() *net.TCPConn

	//获取当前连接模块的连接ID
	GetConnID() uint32

	//获取远程客户端的TCP状态 IP port
	RemoteAddr() net.Addr

	//发送数据给远程的客户端
	Send(data []byte) error
}

//定义一个处理连接业务的方法
type HandleFunc func(*net.TCPConn, []byte, int) error
