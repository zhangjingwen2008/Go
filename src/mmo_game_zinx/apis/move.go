package apis

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"mmo_game_zinx/core"
	"mmo_game_zinx/pb"
	"zinx/ziface"
	"zinx/znet"
)

type MoveApi struct {
	znet.BaseRouter
}

func (m *MoveApi) Handle(request ziface.IRequest) {
	//解析客户端传递进来proto协议
	proto_msg := &pb.Position{}
	err := proto.Unmarshal(request.GetData(), proto_msg)
	if err != nil {
		fmt.Println("Move : Position Unmarshal error ", err)
		return
	}

	//得到当前发送位置的是哪个玩家
	pid, err := request.GetConnection().GetProperty("pid")
	if err != nil {
		fmt.Println("GetProperty pid error ", err)
		return
	}

	fmt.Printf("Player pid = %d, move(%f, %f, %f, %f)", pid, proto_msg.X, proto_msg.Y, proto_msg.Z, proto_msg.V)

	//给其他玩家进行当前玩家的位置信息广播
	player := core.WorldMgrObj.GetPlayerByPid(pid.(int32))
	//广播并更新当前玩家的坐标
	player.UpdatePos(proto_msg.X, proto_msg.Y, proto_msg.Z, proto_msg.V)
}
