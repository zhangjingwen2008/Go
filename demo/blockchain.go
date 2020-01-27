package main

import (
	"bolt"
	"log"
)

//4.引入区块链
//使用数据库代替数组
type BlockChain struct {
	//blocks []*Block //定义一个区块链数组
	db *bolt.DB

	tail []byte //存储最后一个区块的哈希
}

const blockChainDb = "blockChain.db"
const blockBucket = "blockBucket"

//5.定义一个区块链
func NewBlockChain() *BlockChain {
	//return &BlockChain{
	//	blocks: []*Block{genesisBlock},
	//}

	//最后一个区块的哈希，从数据库中读出来的
	var lastHash []byte

	//1.打开数据库
	db, err := bolt.Open(blockChainDb, 0600, nil)
	defer db.Close()
	if err != nil {
		log.Panic("打开数据库失败,", err)
	}

	//将要操作数据库（改写）
	db.Update(func(tx *bolt.Tx) error {
		//2.找到抽屉bucket(如果没有，就创建)
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			//没有抽屉，我们需要创建一个
			bucket, err = tx.CreateBucket([]byte(blockBucket))
			if err != nil {
				log.Panic("创建bucket(b1)失败")
			}

			genesisBlock := GenesisBlock() //创建一个创世块，并作为第一个区块添加到区块链中

			//3.写数据
			//hash作为key，block的字节流作为value
			bucket.Put(genesisBlock.Hash, genesisBlock.Serialize())
			bucket.Put([]byte("LastHashKey"), []byte(genesisBlock.Hash))
			lastHash = genesisBlock.Hash

			//测试
			//blockBytes:=bucket.Get(genesisBlock.Hash)
			//block:=Deserialize(blockBytes)
			//fmt.Printf("block info:%s",block)

		} else {
			lastHash = bucket.Get([]byte("LastHashKey"))
		}

		return nil
	})

	return &BlockChain{db, lastHash}

}

//定义一个创世区块
func GenesisBlock() *Block {
	return NewBlock("这是创世区块哇", []byte{})
}

//5.添加区块
func (bc *BlockChain) AddBlock(data string) {
	//获取前一区块的哈希值
	db:=bc.db				//区块链数据库
	lastHash:=bc.tail		//最后一个区块的Hash

	db.Update(func(tx *bolt.Tx) error {
		//完成数据添加
		bucket:=tx.Bucket([]byte(blockBucket))
		if bucket==nil{
			log.Panic("bucket 不应为空，请检查！")
		}
		//a.创建新的区块
		block:=NewBlock(data,lastHash)

		//b.添加到区块链数组中
		bucket.Put(block.Hash,block.Serialize())
		bucket.Put([]byte("LastHashKey"),block.Hash)
		lastHash=block.Hash
		//还需更新一下内存中的区块链，指的是把最后的tail更新一下
		bc.tail=lastHash

		return nil
	})

}
