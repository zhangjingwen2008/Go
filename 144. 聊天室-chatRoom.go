package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

//创建用户结构体类型
type Client struct {
	C    chan string
	Name string
	Addr string
}

//创建全局map，存储在线用户
var onlineMap map[string]Client

//创建全局channel传递用户信息
var message = make(chan string)

func WriteMsgToClient(clnt Client, conn net.Conn) {
	//监听用户自带channel上是否有消息
	for msg := range clnt.C {
		conn.Write([]byte(msg + "\n"))
	}
}

func MakeMsg(clnt Client, msg string) (buf string) {
	buf = "[" + clnt.Addr + "]" + clnt.Name + ": " + msg
	return
}

func HandlerConnect(conn net.Conn) {
	defer conn.Close()
	//创建channel，判断用户是否活跃（用于超时退出）
	hasData := make(chan bool)

	//获取用户网络地址 IP+port
	netAddr := conn.RemoteAddr().String()
	//创建新连接用户的结构体, 默认用户名是IP+Port
	clnt := Client{make(chan string), netAddr, netAddr}

	//将新连接用户，添加到在线用户map中
	onlineMap[netAddr] = clnt

	//创建专门用来给当前用户发送消息的go程
	go WriteMsgToClient(clnt, conn)

	//发送用户上线消息到全局channel中
	//message <- "[" + netAddr + "]" + clnt.Name + "login"
	message <- MakeMsg(clnt, "login")

	//创建一个channel，用来判断用户退出状态
	isQuit := make(chan bool)

	//创建匿名go程，专门处理用户发送的消息
	go func() {
		buf := make([]byte, 1024*4)
		for {
			n, err := conn.Read(buf)
			if n == 0 { //用户退出
				isQuit <- true
				fmt.Printf("检测到%s退出\n", netAddr)
				return
			}
			if err != nil {
				fmt.Println("conn.Read ERROR:", err)
				return
			}

			//将读到的用户信息，保存到msg中，string类型
			msg := string(buf[:n-1]) //去除输入消息结尾的\n字符

			//提取在线用户列表
			if msg == "who" && len(msg) == 3 {
				conn.Write([]byte("online user list:\n"))
				//便利当前map，获取在线用户
				for _, user := range onlineMap {
					userInfo := user.Addr + ":" + user.Name
					conn.Write([]byte(userInfo + "\n"))
				}
			} else if len(msg) >= 8 && msg[:6] == "rename" { //判断用户发送了【改名】命令
				newName := strings.Split(msg, "|")[1] //提取新名字
				clnt.Name = newName
				onlineMap[netAddr] = clnt //更新map
				conn.Write([]byte("rename success!!\n"))
			} else {
				//将接收到的消息广播给所有用户
				message <- MakeMsg(clnt, msg)
			}
			hasData <- true //表示用户仍是活跃状态
		}
	}()

	//用户退出
	for {
		//监听channel上的数据流动
		select {
		case <-isQuit: //主动退出
			close(clnt.C)                      //当前匿名go程退出，其内的嵌套go程仍然存在，所以仍需要需要手动退出
			delete(onlineMap, clnt.Addr)       //先删除用户，再进行广播，这样就可以避免广播给已删除用户
			message <- MakeMsg(clnt, "logout") //写入用户退出消息到全局channel
			return
		case <-hasData: //用户活跃判断：什么都不做，仅是为了重置下面的计时器，避免超时退出

		case <-time.After(time.Second * 10): //超时退出：计时10秒若没活动，则判定为超时退出
			delete(onlineMap, clnt.Addr)
			message <- MakeMsg(clnt, "logout")
			return
		}
	}
}

func Manager() {
	//初始化map
	onlineMap = make(map[string]Client)

	//监听全局channel是否有数据，有数据存储至msg，无数据阻塞
	for {
		msg := <-message
		//循环发送消息给所有在线用户
		for _, clnt := range onlineMap {
			clnt.C <- msg
		}
	}
}

func main() {
	//创建监听套接字
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.listen ERROR", err)
		return
	}
	defer listener.Close()

	//创建管理者go程，管理map和全局channel
	go Manager()

	//循环监听客户端连接请求
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept ERROR", err)
			return
		}
		//启动go程监听客户端数据请求
		go HandlerConnect(conn)
	}
}
