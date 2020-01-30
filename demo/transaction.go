package main

import (
	"base58"
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
)

const reward = 50

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
	//Sig   string //解锁脚本，这里用地址来模拟
	Signature []byte //真正的数字签名，由r，s拼成的[]byte
	PubKey    []byte	//公钥，不是哈希或地址
}

//定义交易输出
type TXOutput struct {
	Value      float64 //转账金额
	//PubKeyHash string  //锁定脚本，我们用地址模拟
	PubKeyHash []byte	//收款方的公钥的哈希，不是公钥也不是地址
}

//由于现在存储的字段是地址的公钥哈希，所以无法直接创建TXOutput
//为了能够得到公钥哈希，需要写一个Lock函数处理一下
func (output *TXOutput) Lock(address string)  {
	//1.解码
	addressByte:=base58.Decode(address)		//25字节
	len:=len(addressByte)
	//2.截取出公钥哈希：去除Version（1字节），去除校验码（4字节）
	pubKeyHash:=addressByte[1:len-4]
	output.PubKeyHash=pubKeyHash			//真正的锁定动作
}

//给TXOutput提供一个创建的方法，否则无法调用Lock
func NewTXOutput(value float64,address string) *TXOutput {
	output:=TXOutput{
		Value:      value,
	}
	output.Lock(address)
	return &output
}

//设置交易ID
func (tx *Transaction) SetHash() {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(tx)
	if err != nil {
		log.Panic(err)
	}
	data := buffer.Bytes()
	hash := sha256.Sum256(data)
	tx.TXID = hash[:]
}

//实现一个函数，判断当前的交易是否为挖矿交易
func (tx *Transaction) IsCoinbase() bool {
	//1.交易input只有一个
	//2.交易id为空
	//3.交易的index为-1
	if len(tx.TXInput) == 1 && len(tx.TXInput[0].TXid) == 0 && tx.TXInput[0].index == -1 {
		return true
	}
	return false
}

//2.提供创建交易方法（挖矿交易）
func NewCoinbaseTX(address string, data string) *Transaction {
	//挖矿交易的特点
	//1.只有1个input
	//2.无需引用交易ID
	//3.无需引用index
	//矿工由于挖矿时无需指定签名，所以这个PubKey字段可以由矿工自由填写数据，一般是填矿池的名字
	input := TXInput{[]byte{},-1,nil,[]byte(data)}
	//output := TXOutput{reward,[]byte(address)}
	output:=NewTXOutput(reward,address)				//新的创建方法

	tx := Transaction{ //对于挖矿交易来说，只有一个input和一个output
		TXID:     []byte{},
		TXInput:  []TXInput{input},
		TXOutput: []TXOutput{*output},
	}
	tx.SetHash()
	return &tx
}

//创建普通的转账交易
func NewTransaction(from, to string, amount float64, bc *BlockChain) *Transaction {

	//1.创建交易之后要进行数字签名->所以需要私钥->打开钱包（打开钱包"NewWallets()"）
	ws:=NewWallets()

	//2.找到自己的钱包，根据地址返回自己的wallet
	wallet:=ws.WalletsMap[from]
	if wallet==nil{
		fmt.Println("没有找到给地址的私钥，交易创建失败!\n")
		return nil
	}

	//3.得到对应的公钥，私钥
	pubKey:=wallet.PubKey
	//privateKey:=wallet.Private

	pubKeyHash:=HashPubKey(pubKey)

	//1.找到最合理的UTXO集合 map[string][]uint64
	utxos, resValue := bc.FindNeedUTXOs(pubKeyHash, amount)

	if resValue < amount {
		fmt.Println("余额不足，交易失败")
		return nil
	}

	var inputs []TXInput
	var outputs []TXOutput

	//2.创建交易输入，将这些UTXO逐一转成inputs
	for id, indexArray := range utxos {
		for _, i := range indexArray {
			input := TXInput{[]byte(id), int64(i), nil,pubKey}
			inputs = append(inputs, input)
		}
	}

	//3.创建交易输出
	//output := TXOutput{amount, to}
	output:=NewTXOutput(amount,to)
	outputs = append(outputs, *output)

	//4.如果有零钱，要找零
	if resValue > amount {
		//找零
		output=NewTXOutput(resValue-amount,from)
		outputs = append(outputs, *output)
	}

	tx := Transaction{
		TXID:     []byte{},
		TXInput:  inputs,
		TXOutput: outputs,
	}
	tx.SetHash()
	return &tx
}

//3.创建挖矿交易
//4.根据交易调整程序
