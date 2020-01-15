package main

import (
	"fmt"
	"zinx/ziface"
	"zinx/znet"
)

/*
	基于Zinx框架来开发的 服务器端应用程序
*/

//ping test 自定义路由
type PingRouter struct {
	znet.BaseRouter
}

//Test Handle
func (this *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call Router Handle...")

	//先读取客户端的数据
	fmt.Println("recv from client: msgID= ", request.GetMsgID(),
		", data= ", string(request.GetData()))

	//再回写数据
	err := request.GetConnection().SendMsg(request.GetMsgID(), []byte("ping...ping...ping"))
	if err != nil {
		fmt.Println(err)
	}

}

func main() {
	//1.创建一个server句柄，使用Zinx的api
	s := znet.NewServer("[zinx V0.5]")

	//2.给当前zinx框架添加一个自定义的router
	s.AddRouter(&PingRouter{})

	//3.启动server
	s.Serve()
}
