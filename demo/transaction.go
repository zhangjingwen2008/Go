package main

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/rand"
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
	Index int64  //引用的output的索引值
	//Sig   string //解锁脚本，这里用地址来模拟
	Signature []byte //真正的数字签名，由r，s拼成的[]byte
	PubKey    []byte //公钥，不是哈希或地址
}

//定义交易输出
type TXOutput struct {
	Value float64 //转账金额
	//PubKeyHash string  //锁定脚本，我们用地址模拟
	PubKeyHash []byte //收款方的公钥的哈希，不是公钥也不是地址
}

//由于现在存储的字段是地址的公钥哈希，所以无法直接创建TXOutput
//为了能够得到公钥哈希，需要写一个Lock函数处理一下
func (output *TXOutput) Lock(address string) {
	output.PubKeyHash = GetPubKeyFromAddress(address) //真正的锁定动作
}

//给TXOutput提供一个创建的方法，否则无法调用Lock
func NewTXOutput(value float64, address string) *TXOutput {
	output := TXOutput{
		Value: value,
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
	if len(tx.TXInput) == 1 && len(tx.TXInput[0].TXid) == 0 && tx.TXInput[0].Index == -1 {
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
	input := TXInput{[]byte{}, -1, nil, []byte(data)}
	//output := TXOutput{reward,[]byte(address)}
	output := NewTXOutput(reward, address) //新的创建方法

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
	ws := NewWallets()

	//2.找到自己的钱包，根据地址返回自己的wallet
	wallet := ws.WalletsMap[from]
	if wallet == nil {
		fmt.Println("没有找到给地址的私钥，交易创建失败!\n")
		return nil
	}

	//3.得到对应的公钥，私钥
	pubKey := wallet.PubKey
	privateKey := wallet.Private

	pubKeyHash := HashPubKey(pubKey) //传递公钥的哈希，而不是传递地址

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
			input := TXInput{[]byte(id), int64(i), nil, pubKey}
			inputs = append(inputs, input)
		}
	}

	//3.创建交易输出
	//output := TXOutput{amount, to}
	output := NewTXOutput(amount, to)
	outputs = append(outputs, *output)

	//4.如果有零钱，要找零
	if resValue > amount {
		//找零
		output = NewTXOutput(resValue-amount, from)
		outputs = append(outputs, *output)
	}

	tx := Transaction{
		TXID:     []byte{},
		TXInput:  inputs,
		TXOutput: outputs,
	}
	tx.SetHash()

	bc.SignTransaction(&tx, privateKey)

	return &tx
}

//签名的具体实现
//参数为：私钥、inputs里面所有引用的交易结构map[string]Transaction
func (tx *Transaction) Sign(privateKey *ecdsa.PrivateKey, prevTXs map[string]Transaction) {
	//具体签名动作// TODO

	//1.创建一个当前交易的副本：txCopy，使用函数：TrimmedCopy：要把Signature和PubKey字段设置为nil
	txCopy := tx.TrimmedCopy()
	txCopy.SetHash()
	//2.循环遍历txCopy的inputs，得到这个input索引的output的公钥哈希
	for i, input := range txCopy.TXInput {
		prevTX := prevTXs[string(input.TXid)]
		if len(prevTX.TXID) == 0 {
			log.Panic("引用的交易无效")
		}

		//不要对input进行赋值，这是一个副本，要对txCoput.TXInputs[xx]进行操作，否则无法把pubKeyHash传进来
		txCopy.TXInput[i].PubKey = prevTX.TXOutput[input.Index].PubKeyHash

		//具备所需3个数据，开始做哈希处理
		//3.生成要签名的数据。要签名的数据一定是哈希值
		//a.对每一个input都要签名一次，签名的数据是由当前input引用的output的哈希+当前的outputs（都承载在当前这个txCopy里面）
		//b.对拼接好的txCopy进行哈希处理，SetHash得到TXID，即所需的签名数据
		txCopy.SetHash()
		//还原，以免影响后面input的签名
		txCopy.TXInput[i].PubKey = nil
		signDataHash := txCopy.TXID
		//4.执行签名动作得到r,s字节流
		r,s,err:=ecdsa.Sign(rand.Reader, privateKey, signDataHash)
		if err!=nil{
			log.Panic(err)
		}
		//5.放到我们所签名的input的Signature中
		signature:=append(r.Bytes(),s.Bytes()...)
		tx.TXInput[i].Signature=signature

	}

}

func (tx *Transaction) TrimmedCopy() Transaction {
	var inputs []TXInput
	var outputs []TXOutput

	for _, input := range tx.TXInput {
		inputs = append(inputs, TXInput{input.TXid, input.Index, nil, nil})
	}
	for _, output := range tx.TXOutput {
		outputs = append(outputs, output)
	}

	return Transaction{[]byte{}, inputs, outputs}
}
