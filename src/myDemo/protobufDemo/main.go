package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"myDemo/protobufDemo/pb"
)

//protocol buffer/protobuf 的编码和解码使用
func main() {
	//定义一个Person结构体对象
	person := &pb.Person{
		Name:   "Hunter景文",
		Age:    16,
		Emails: []string{"zhangjingwen2008@163.com", "562045804@qq.com"},
		Phones: []*pb.PhoneNumber{
			&pb.PhoneNumber{
				Number: "17300000000",
				Type:   pb.PhoneType_HOME,
			},
			&pb.PhoneNumber{
				Number: "153XXXXXXXX",
				Type:   pb.PhoneType_WORK,
			},
		},
	}

	//编码
	//将person对象 就是将protobuf的message进行序列化，得到一个二进制文件
	data, err := proto.Marshal(person)
	//data 就是我们要进行网络传输的数据，对端需要按照Message Person格式进行解析
	if err != nil {
		fmt.Println("Marshal err:", err)
	}

	//解码
	newdata := &pb.Person{}
	err = proto.Unmarshal(data, newdata)
	if err != nil {
		fmt.Println("Unmarshal err:", err)
	}
	fmt.Println("原数据： ", person)
	fmt.Println("解码数据： ", newdata)
}
