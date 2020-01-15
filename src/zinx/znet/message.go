package znet

type Message struct {
	Id      uint32 //消息ID
	DataLen uint32 //消息长度
	Data    []byte //消息内容
}

//创建Message消息包的方法
func NewMsgPackage(id uint32, data []byte) *Message {
	return &Message{
		Id:      id,
		DataLen: uint32(len(data)),
		Data:    data,
	}
}

//獲取消息的ID
func (m *Message) GetMsgId() uint32 {
	return m.Id
}

//获取消息的长度
func (m *Message) GetMsgLen() uint32 {
	return m.DataLen
}

//获取消息的内容
func (m *Message) GetData() []byte {
	return m.Data
}

//设置消息的ID
func (m *Message) SetMsgId(id uint32) {
	m.SetMsgId(id)
}

//设置消息的内容
func (m *Message) SetData(data []byte) {
	m.SetData(data)
}

//设置消息的长度
func (m *Message) SetDataLen(len uint32) {
	m.SetDataLen(len)
}
