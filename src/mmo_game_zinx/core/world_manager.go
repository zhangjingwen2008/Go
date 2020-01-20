package core

import "sync"

/*
	当前游戏的世界总管理模块
*/

type WorldManager struct {
	//AOIManager 当前世界地图AOI的管理模块
	AoiMgr *AOIManager
	//当前全部在线的Players集合
	Players map[int32]*Player
	//保护Players集合的锁
	pLock sync.RWMutex
}

//提供一个对外的世界管理模块的句柄（全局）
var WorldMgrObj *WorldManager

//初始化方法(在调用此包时就会自动调用，而且仅调用一次)
func init() {
	WorldMgrObj = &WorldManager{
		AoiMgr:  NewAOIManager(AOI_MIN_X, AOI_MAX_X, AOI_CNTS_X, AOI_MIN_Y, AOI_MAX_Y, AOI_CNTS_Y), //创建世界AOI地图规划
		Players: make(map[int32]*Player),                                                           //初始化Player集合
	}
}

//添加一个玩家
func (wm *WorldManager) AddPlayer(player *Player) {
	wm.pLock.Lock()
	wm.Players[player.Pid] = player
	wm.pLock.Unlock()

	//将player添加到 AOIManager中
	wm.AoiMgr.AddToGridByPos(int(player.Pid), player.X, player.Z)
}

//删除玩家
func (wm *WorldManager) RemovePlayerByPid(Pid int32) {
	player := wm.Players[Pid]                                   //得到当前玩家
	wm.AoiMgr.RemoveFromGridByPos(int(Pid), player.X, player.Z) //将玩家从AOIManager删除

	wm.pLock.Lock()
	delete(wm.Players, Pid) //将玩家从世界管理器中删除
	wm.pLock.Unlock()
}

//通过玩家ID查询Player对象
func (wm *WorldManager) GetPlayerByPid(Pid int32) (player *Player) {
	wm.pLock.RLock()
	defer wm.pLock.RUnlock()

	return wm.Players[Pid]
}

//获取全部的在线玩家
func (wm *WorldManager) GetAllPlayers() []*Player {
	wm.pLock.RLock()
	defer wm.pLock.RUnlock()

	players := make([]*Player, 0)

	//添加到切片中
	for _, p := range wm.Players {
		players = append(players, p)
	}

	return players
}