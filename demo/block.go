package main

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"log"
	"time"
)

//0.定义结构
type Block struct {
	//1.版本号
	Version uint64
	//2.前区块哈希
	PrevHash []byte
	//3.Merkle根
	MerkleRoot []byte
	//4.时间戳
	TimeStamp uint64
	//5.难度值
	Difficulty uint64
	//6.随机数 即挖矿要找的数据
	Nonce uint64

	//a.当前区块哈希
	Hash []byte
	//b.数据
	//Data []byte
	//真实的交易数组
	Transactions []*Transaction
}

//2.创建区块
func NewBlock(txs []*Transaction, prevBlockHash []byte) *Block {
	block := Block{
		Version:      00,
		PrevHash:     prevBlockHash,
		MerkleRoot:   []byte{},
		TimeStamp:    uint64(time.Now().Unix()),
		Difficulty:   0,
		Nonce:        0,
		Hash:         []byte{},
		Transactions: txs,
	}

	block.MerkleRoot=block.MakeMerkleRoot()

	//block.SetHash()
	//创建一个pow对象
	pow := NewProofOfWork(&block)
	//查找随机数，不停地进行哈希运算
	hash, nonce := pow.Run()
	//根据挖矿结果对区块数据进行更新（补充）
	block.Hash = hash
	block.Nonce = nonce

	return &block
}

func (block *Block) Serialize() []byte {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(&block)
	if err != nil {
		log.Panic("编码失败")
	}

	return buffer.Bytes()
}

func Deserialize(data []byte) Block {
	decoder := gob.NewDecoder(bytes.NewReader(data))
	var block Block
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic("解码出错！")
	}
	return block
}

/*
//3.生成哈希
func (block *Block) SetHash() {
	//1.拼装数据
		var blockInfo []byte
		blockInfo = append(blockInfo, Uint64ToByte(block.Version)...)
		blockInfo = append(blockInfo, block.PrevHash...)
		blockInfo = append(blockInfo, block.MerkleRoot...)
		blockInfo = append(blockInfo, Uint64ToByte(block.TimeStamp)...)
		blockInfo = append(blockInfo, Uint64ToByte(block.Difficulty)...)
		blockInfo = append(blockInfo, Uint64ToByte(block.Nonce)...)
		blockInfo = append(blockInfo, block.Data...)
	tmp := [][]byte{
		Uint64ToByte(block.Version),
		block.PrevHash,
		block.MerkleRoot,
		Uint64ToByte(block.TimeStamp),
		Uint64ToByte(block.Difficulty),
		Uint64ToByte(block.Nonce),
		block.Data,
	}
	//将二维的切片数组连接起来，返回一个一维的切片
	blockInfo := bytes.Join(tmp, []byte{})

	//2.sha256
	//func Sum256(data []byte) [Size]byte
	hash := sha256.Sum256(blockInfo)
	block.Hash = hash[:]
}
*/

func (block *Block) toByte() []byte {
	return []byte{}
}

//将Uint64转成byte
func Uint64ToByte(num uint64) []byte {
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buffer.Bytes()
}

//模拟MerkleRoot，只是对交易的数据做简单的拼接
func (block *Block) MakeMerkleRoot() []byte {
	//TODO
	return []byte{}
}
