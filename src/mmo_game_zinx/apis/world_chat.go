package apis

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"mmo_game_zinx/core"
	"mmo_game_zinx/pb"
	"zinx/ziface"
	"zinx/znet"
)

//世界聊天 路由业务

type WorldChatApi struct {
	znet.BaseRouter
}

func (wc *WorldChatApi) Handle(request ziface.IRequest) {
	//1.解析客户端传递进来的proto协议
	proto_msg := &pb.Talk{}
	err := proto.Unmarshal(request.GetData(), proto_msg)
	if err != nil {
		fmt.Println("Talk Unmarshal ERR:", err)
		return
	}

	//2.当前的聊天数据是属于哪个玩家
	Pid, err := request.GetConnection().GetProperty("pid")

	//3.根据Pid得到对应的Player对象
	player := core.WorldMgrObj.GetPlayerByPid(Pid.(int32)) //因为步骤2得到的Pid是空接口类型，所以需要断言来转换类型

	//4.将这个消息广播给全部在线玩家
	player.Talk(proto_msg.Content)

}
