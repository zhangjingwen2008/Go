package main

import (
	"bolt"
	"fmt"
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
func NewBlockChain(address string) *BlockChain {
	//return &BlockChain{
	//	blocks: []*Block{genesisBlock},
	//}

	//最后一个区块的哈希，从数据库中读出来的
	var lastHash []byte

	//1.打开数据库
	db, err := bolt.Open(blockChainDb, 0600, nil)
	//defer db.Close()
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

			genesisBlock := GenesisBlock(address) //创建一个创世块，并作为第一个区块添加到区块链中

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
func GenesisBlock(address string) *Block {
	coinbase := NewCoinbaseTX(address, "这是创世区块哇")
	return NewBlock([]*Transaction{coinbase}, []byte{})
}

//5.添加区块
func (bc *BlockChain) AddBlock(txs []*Transaction) {
	//获取前一区块的哈希值
	db := bc.db         //区块链数据库
	lastHash := bc.tail //最后一个区块的Hash

	db.Update(func(tx *bolt.Tx) error {
		//完成数据添加
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			log.Panic("bucket 不应为空，请检查！")
		}
		//a.创建新的区块
		block := NewBlock(txs, lastHash)

		//b.添加到区块链数组中
		bucket.Put(block.Hash, block.Serialize())
		bucket.Put([]byte("LastHashKey"), block.Hash)
		lastHash = block.Hash
		//还需更新一下内存中的区块链，指的是把最后的tail更新一下
		bc.tail = lastHash

		return nil
	})

}

func (bc *BlockChain) FindUTXOs(address string) []TXOutput {
	var UTXO []TXOutput
	//我们定义一个map来保存消费过的output，key是这个output的交易id，value是这个交易中索引的数组
	//map[交易id][]int64
	spentOutputs := make(map[string][]int64)

	//创建迭代器
	it := bc.NewIterator()
	for {
		//1.遍历区块
		block := it.Next()

		//2.遍历交易
		for _, tx := range block.Transactions {
			fmt.Println("current txid:", tx.TXID)
			//3.遍历output, 找到和自己相关的UTXO（在添加output之前检查一下是否已经消耗过）
			for i, output := range tx.TXOutput {
				fmt.Println("current index:", i)
				if output.PubKeyHash == address { //这个output和我们目标地址相同，满足条件，加到返回UTXO数组中
					UTXO = append(UTXO, output)
				}
			}
			//4.遍历input，找到自己花费过的UTXO的集合（把自己消耗过的标示出来）
			for _, input := range tx.TXInput {
				//判断一下当前这个input和目标（李四）是否一致，如果相同，说明这个是李四消耗过的output，就加进来
				if input.Sig == address {
					indexArray := spentOutputs[string(input.TXid)]
					indexArray = append(indexArray, input.index)
				}
			}

		}

		if len(block.PrevHash) == 0 {
			fmt.Println("区块遍历完成退出")
			break
		}
	}

	return UTXO
}
