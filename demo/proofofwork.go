package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

//定义一个工作量证明的结果ProofOfWork
type ProofOfWork struct {
	//a.block
	block *Block
	//b.目标值
	target *big.Int
}

//2.提供创建POW的函数
func NewProofOfWork(block *Block) *ProofOfWork {
	pow := ProofOfWork{
		block: block,
	}
	//指定难度值，现在是一个string类型，需要进行转换
	targetStr := "0000f00000000000000000000000000000000000000000000000000000000000"
	//引入的辅助变量，目的是将上面的难度值转成big.Int
	tmpInt := big.Int{}
	//将难度值赋值给big.Int，指定16进制的格式
	tmpInt.SetString(targetStr, 16)
	pow.target = &tmpInt
	return &pow
}

//3.提供计算不断计算hash的函数
func (pow *ProofOfWork) Run() ([]byte, uint64) {
	//1.拼装数据（区块的数据，还有不断变化的随机数）
	//2.做哈希运算
	//3.与pow中的target进行比较
	//a.找到了，退出返回
	//b.没找到，继续找，随机数加1

	var nonce uint64
	block := pow.block
	var hash [32]byte

	fmt.Println("开始挖矿……")
	for {
		//1.拼装数据（区块的数据，还有不断变化的随机数）
		tmp := [][]byte{
			Uint64ToByte(block.Version),
			block.PrevHash,
			block.MerkleRoot,
			Uint64ToByte(block.TimeStamp),
			Uint64ToByte(block.Difficulty),
			Uint64ToByte(nonce),
			//block.Data,			//只对区块头做哈希值，区块体通过MerkleRoot产生影响
		}
		//将二维的切片数组连接起来，返回一个一维的切片
		blockInfo := bytes.Join(tmp, []byte{})
		//2.做哈希运算
		hash = sha256.Sum256(blockInfo)
		//3.与pow中的target进行比较
		tmpInt := big.Int{}
		//将我们得到的hash数组转换成一个big.int
		tmpInt.SetBytes(hash[:])
		//比较当前的哈希与目标哈希值，如果当前的哈希值小于目标哈希值，就说明找到了，否则继续找
		if tmpInt.Cmp(pow.target) == -1 { //两个BigInt进行比较
			//a.找到了，退出返回
			fmt.Printf("挖矿成功，hash：%x，nonce：%d\n", hash, nonce)
			return hash[:], nonce
		} else {
			//b.没找到，继续找，随机数加1
			nonce++
		}
	}

}

//4.提供一个校验函数

//	//指定难度值，现在是一个string类型，需要进行转换
//	targetStr:="0000100000000000000000000000000000000000000000000"
//	//引入的辅助变量，目的是将上面的难度值转成big.Int
//	tmpInt:=big.Int{}
//	//将难度值赋值给big.Int，指定16进制的格式
//	tmpInt.SetString(targetStr, 16)
//	fmt.Println(tmpInt)
//}
