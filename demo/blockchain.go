package main

import (
	"bolt"
	"bytes"
	"crypto/ecdsa"
	"errors"
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

	for _,tx:=range txs{
		if !bc.VerifyTransaction(tx){
			fmt.Println("矿工发现无效交易！")
			return
		}
	}

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

func (bc *BlockChain) FindUTXOs(pubKeyHash []byte) []TXOutput {
	var UTXO []TXOutput

	txs := bc.FindUTXOTransactions(pubKeyHash)
	for _, tx := range txs {
		for _, output := range tx.TXOutput {
			if bytes.Equal(pubKeyHash,output.PubKeyHash){
				UTXO = append(UTXO, output)
			}
		}
	}
	return UTXO
}


func (bc *BlockChain) FindNeedUTXOs(senderPubKeyHash []byte, amount float64) (map[string][]uint64, float64) {
	utxos := make(map[string][]uint64) //找到的合理的utxos集合
	var calc float64                   //找到的utxos里面包含的钱的总数

	txs := bc.FindUTXOTransactions(senderPubKeyHash)

	for _, tx := range txs {
		for i, output := range tx.TXOutput {
			if bytes.Equal(senderPubKeyHash,output.PubKeyHash) {		//两个[]byte的比较
				//这个output和我们目标地址相同，满足条件，加到返回UTXO数组中
					//UTXO = append(UTXO, output)
					//找到自己需要的最少UTXO

					if calc < amount {
						//1.把UTXO加进来
						utxos[string(tx.TXID)] = append(utxos[string(tx.TXID)], uint64(i))
						//2.统计当前UTXO总额
						calc += output.Value

						//加完之后若满足条件
						//3.比较是否满足转账需求
						//a.满足，直接返货 utxos，calc
						//b.不满足，继续统计
						if calc >= amount {
							//break
							fmt.Println("找到了满足的金额：", calc)
							return utxos, calc
						}

					} else {
						fmt.Println("不满足转账金额，当前总额：", calc, " 目标金额：", amount)
					}

			}
		}
	}

	return utxos, calc
}

func (bc *BlockChain) FindUTXOTransactions(senderPubKeyHash []byte) []*Transaction {
	var txs []*Transaction //存储所有包含utxo交易集合
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

		OUTPUT:
			//3.遍历output, 找到和自己相关的UTXO（在添加output之前检查一下是否已经消耗过）
			for i, output := range tx.TXOutput {
				fmt.Println("current index:", i)
				//在这里做一个过滤，将所有消耗过的output和当前的所即将添加的output对比一下
				//如果相同，则跳过，否则添加
				//如果当前的交易id存在于我们已经表示的map，那么说明这个交易里面有消耗过的output
				if spentOutputs[string(tx.TXID)] != nil {
					for _, j := range spentOutputs[string(tx.TXID)] {
						if int64(i) == j {
							//当前准备添加output已经消耗过了，无需添加
							continue OUTPUT
						}
					}
				}

				if bytes.Equal(output.PubKeyHash,senderPubKeyHash) { //这个output和我们目标地址相同，满足条件，加到返回UTXO数组中
					txs = append(txs, tx) //返回所有包含我的utxo的交易的集合
				}
			}

			//如果当前交易是挖矿交易的话，那么不做遍历，直接跳过
			if !tx.IsCoinbase() {
				//4.遍历input，找到自己花费过的UTXO的集合（把自己消耗过的标示出来）
				for _, input := range tx.TXInput {
					//判断一下当前这个input和目标（李四）是否一致，如果相同，说明这个是李四消耗过的output，就加进来
					pubKeyHash:=HashPubKey(input.PubKey)
					if bytes.Equal(pubKeyHash,senderPubKeyHash) {
						//indexArray := spentOutputs[string(input.TXid)]
						//indexArray = append(indexArray, input.index)
						spentOutputs[string(input.TXid)] = append(spentOutputs[string(input.TXid)], input.index)
					}
				}
			} else {
				fmt.Println("这是coinBase，不做input输出")
			}

		}

		if len(block.PrevHash) == 0 {
			fmt.Println("区块遍历完成退出")
			break
		}
	}

	return txs
}

//根据id查找交易本身，需要遍历整个区块链
func (bc *BlockChain) FindTransactionByTXid(id []byte) (Transaction,error) {
	//1.遍历区块链
	it:=bc.NewIterator()
	for{
		//2.遍历交易
		block:=it.Next()
		for _,tx :=range block.Transactions {
			//3.比较交易，找到了直接退出
			if bytes.Equal(tx.TXID, id) {
				return *tx,nil
			}
		}

		//TODO

		//4.如果没找到，返回空Transaction,同时返回错误状态
		if len(block.PrevHash)==0{
			fmt.Println("区块链遍历结束")
			break
		}
	}

	return Transaction{},errors.New("无效的交易id，请检查！j")
}

func (bc *BlockChain) SignTransaction(tx *Transaction, privateKey *ecdsa.PrivateKey)  {
	//签名，交易创建的最后进行签名
	prevTXs:=make(map[string]Transaction)

	//找到所有引用的交易
	//1.根据inputs来找，有多少input，就遍历多少次
	//2.找到目标交易，（根据TXid来找）
	//3.添加到prevTXs里面
	for _,input :=range tx.TXInput{
		//根据id查找交易本身，需要遍历整个区块链
		tx,err:=bc.FindTransactionByTXid(input.TXid)
		if err!=nil{
			log.Panic(err)
		}

		prevTXs[string(input.TXid)]=tx
	}

	tx.Sign(privateKey,prevTXs)
}

func (bc *BlockChain) VerifyTransaction(tx *Transaction) bool {
	//签名，交易创建的最后进行签名
	prevTXs:=make(map[string]Transaction)

	//找到所有引用的交易
	//1.根据inputs来找，有多少input，就遍历多少次
	//2.找到目标交易，（根据TXid来找）
	//3.添加到prevTXs里面
	for _,input :=range tx.TXInput{
		//根据id查找交易本身，需要遍历整个区块链
		tx,err:=bc.FindTransactionByTXid(input.TXid)
		if err!=nil{
			log.Panic(err)
		}

		prevTXs[string(input.TXid)]=tx

	}

	return tx.Verify(prevTXs)
}