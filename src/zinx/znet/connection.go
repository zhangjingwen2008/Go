package znet

import (
	"errors"
	"fmt"
	"io"
	"net"
	"sync"
	"zinx/utils"
	"zinx/ziface"
)

/*
	连接模块
*/
type Connection struct {
	//当前Conn隶属于那个Server
	TcpServer ziface.IServer

	//当前连接的socket TCP套接字
	Conn *net.TCPConn

	//连接的ID
	ConnID uint32

	//当前的连接状态
	isClosed bool

	//消息的管理MsgID 和对应的处理业务API关系
	MsgHandler ziface.IMsgHandle

	//告知当前连接已经退出的/停止 channel
	ExitChan chan bool

	//无缓冲的管道，用于读、写Goroutine之间的消息通信
	msgChan chan []byte

	//连接属性集合
	property map[string]interface{}

	//保护连接属性修改的锁
	propertyLock sync.RWMutex

	//【ZinxV0.2】
	////当前连接所绑定的处理业务方法API
	//handleAPI ziface.HandleFunc

	//【ZinxV0.3】
	//该链接处理的方法Router
	//Router ziface.IRouter
}

//初始化连接模块的方法
func NewConnection(server ziface.IServer, conn *net.TCPConn, connID uint32, msgHandler ziface.IMsgHandle) *Connection {
	c := &Connection{
		TcpServer:  server,
		Conn:       conn,
		ConnID:     connID,
		MsgHandler: msgHandler,
		//Router: router,				//【ZinxV0.3】
		//handleAPI: callback_api,		//【ZinxV0.2】
		isClosed: false,
		ExitChan: make(chan bool, 1),
		msgChan:  make(chan []byte),
		property: make(map[string]interface{}),
	}

	//将conn加入到ConnManager中
	c.TcpServer.GetConnMgr().Add(c)

	return c
}

//链接的读业务方法
func (c *Connection) StartReader() {
	fmt.Println("[Reader Goroutine is running...]")
	defer fmt.Println("connID = ", c.ConnID, " [Reader is exit!], remote addr is ", c.RemoteAddr().String())
	defer c.Stop()

	for {
		//【ZinxV0.4】
		//读取客户端数据到buf中
		//buf := make([]byte, utils.GlobalObject.MaxPackageSize)
		//_, err := c.Conn.Read(buf)
		//if err != nil {
		//	fmt.Println("recv buf err", err)
		//	continue //当前包读失败，不代表下一个也失败
		//}
		//创建一个拆包解包对象
		dp := NewDataPack()
		//读取客户端的Msg Head 二进制流 8个字节，
		headData := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(c.GetTCPConnection(), headData); err != nil {
			fmt.Println("read msg head error: ", err)
			break
		}

		//拆包，得到msgID 和 msgDataLen 放在msg消息中
		msg, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("unpack error ", err)
			break
		}

		//根据dataLen 再次读取Data，放在msg.Data中
		var data []byte
		if msg.GetMsgLen() > 0 {
			data = make([]byte, msg.GetMsgLen())
			if _, err := io.ReadFull(c.GetTCPConnection(), data); err != nil {
				fmt.Println("read msg data error ", err)
				break
			}
		}
		ms := msg.(*Message)
		ms.Data = data
		//msg.SetData(data)

		//得到当前conn数据的Request请求数据
		req := Request{
			conn: c,
			//data: buf,	//【ZinxV0.4】
			msg: ms,
		}

		if utils.GlobalObject.WorkerPoolSize > 0 {
			//已经开启了工作池机制，将消息发送给Worker工作池处理即可
			c.MsgHandler.SendMsgToTaskQueue(&req)
		} else {
			//从路由中，找到注册绑定的Conn对应的router调用
			//根据绑定好的MsgID 找到对应处理api业务 执行
			go c.MsgHandler.DoMsgHandler(&req)
		}

		//【ZinxV0.3】
		//执行注册的路由方法
		//go func(request ziface.IRequest) {
		//	c.Router.PreHandle(request)
		//	c.Router.Handle(request)
		//	c.Router.PostHandle(request)
		//}(&req)
		////从路由中，找到注册绑定的Conn对应的route调用

		//【ZinxV0.2】
		//调用当前连接所绑定的API
		//if err := c.handleAPI(c.Conn, buf, cnt); err != nil {
		//	fmt.Println("ConnID ", c.ConnID, " handle is error", err)
		//	break
		//}
	}
}

/*
	写消息Goroutine，专门发送给客户端消息的模块
*/
func (c *Connection) StartWriter() {
	fmt.Println("[Writer Goroutine is running...]")
	defer fmt.Println(c.RemoteAddr().String(), " [conn Writer exit!]")

	//不断循环阻塞等待channel的消息，进行写给客户
	for {
		select {
		case data := <-c.msgChan:
			//有数据要写给客户端
			if _, err := c.Conn.Write(data); err != nil {
				fmt.Println("Send data error, ", err)
				return
			}
		case <-c.ExitChan:
			//代表Reader已经退出，此时Writer也要退出
			return
		}
	}
}

//启动连接 让当前的连接准备开始工作
func (c *Connection) Start() {
	fmt.Println("Conn Start()... ConnId = ", c.ConnID)

	//启动从当前连接的读数据的业务
	go c.StartReader()
	//启动从当前连接写数据的业务
	go c.StartWriter()

	//按照开发者传递进来的 创建连接之后需要调用的处理业务，执行对应Hook函数
	c.TcpServer.CallOnConnStart(c)
}

//停止连接 结束当前连接的工作
func (c *Connection) Stop() {
	fmt.Println("Conn Stop().. ConnId=", c.ConnID)

	//若当前链接已经关闭
	if c.isClosed == true {
		return
	}
	c.isClosed = true

	//调用开发者注册的 在销毁连接之前 需要执行的业务Hook函数
	c.TcpServer.CallOnConnStop(c)

	//关闭socket连接
	c.Conn.Close()

	//告知Writer关闭
	c.ExitChan <- true

	//将当前连接从ConnMgr中摘除掉
	c.TcpServer.GetConnMgr().Remove(c)

	//回收资源
	close(c.ExitChan)
	close(c.msgChan)
}

//获取当前连接的绑定socket conn
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

//获取当前连接模块的连接ID
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

//获取远程客户端的TCP状态 IP port
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

//提供一个SendMsg方法 将我们要发送给客户端的数据，先进行封包，再发送
func (c *Connection) SendMsg(msgId uint32, data []byte) error {
	if c.isClosed == true {
		return errors.New("Connection closed when send msg")
	}

	//将data进行封包 MsgDataLen|MsgID|Data
	db := NewDataPack()
	binaryMsg, err := db.Pack(NewMsgPackage(msgId, data))
	if err != nil {
		fmt.Println("Pack error msg id = ", msgId)
		return errors.New("Pack error msg")
	}

	//发送封包
	c.msgChan <- binaryMsg
	//【ZinxV0.5】
	//if _, err := c.Conn.Write(binaryMsg); err != nil {
	//	fmt.Println("Write msg id ", msgId, " error : ", err)
	//	return errors.New("conn Write err")
	//}

	return nil
}

//设置连接属性
func (c *Connection) SetProperty(key string, value interface{}) {
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()

	//添加一个连接属性
	c.property[key] = value
}

//获取连接属性
func (c *Connection) GetProperty(key string) (interface{}, error) {
	c.propertyLock.RLock()
	defer c.propertyLock.RUnlock()

	//读取属性
	if value, ok := c.property[key]; ok {
		return value, nil
	}
	return nil, errors.New("no property found")
}

//移除连接属性
func (c *Connection) RemoveProperty(key string) {
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()

	//删除属性
	delete(c.property, key)
}
