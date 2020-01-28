package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

//1.定义交易结构
type Transaction struct {
	TXID     []byte     //交易ID
	TXInput  []TXInput  //交易输入的数组
	TXOutput []TXOutput //交易输出的数组
}

//定义交易输入
type TXInput struct {
	TXid  []byte //引用的交易ID
	index int64  //引用的output的索引值
	Sig   string //解锁脚本，这里用地址来模拟
}

//定义交易输出
type TXOutput struct {
	value float64			//转账金额
	PubKeyHash string		//锁定脚本，我们用地址模拟
}

//设置交易ID
func (tx *Transaction) SetHash() {
	var buffer bytes.Buffer
	encoder:=gob.NewEncoder(&buffer)
	err:= encoder.Encode(tx)
	if err!=nil{
		log.Panic(err)
	}
	data:= buffer.Bytes()
	hash:= sha256.Sum256(data)
	tx.TXID=hash[:]
}

//2.提供创建交易方法
//3.创建挖矿交易
//4.根据交易调整程序
