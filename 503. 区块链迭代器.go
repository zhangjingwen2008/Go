package main

/*
	import (
		"bolt"
		"log"
	)

	type BLockChainIterator struct {
		db *bolt.DB
		//游标，用于不断索引
		currentHashPointer []byte
	}

	func (bc *BlockChain) NewIterator() *BLockChainIterator {
		return &BLockChainIterator{
			db:                 bc.db,
			currentHashPointer: bc.tail, //最初指向区块链的最后一个区块，随着Next的调用，不断向创世块方向遍历
		}
	}

	//迭代器是属于区块链的
	//Next方式是属于迭代器的
	//1.返回当前区块
	//2.指针前移
	func (it *BLockChainIterator) Next() *Block {
		var block Block
		it.db.View(func(tx *bolt.Tx) error {
			bucket := tx.Bucket([]byte(blockBucket))
			if bucket == nil {
				log.Panic("迭代器遍历时，bucket不应该为空")
			}
			blockTmp := bucket.Get(it.currentHashPointer)
			//解码
			block = Deserialize(blockTmp)

			it.currentHashPointer = block.PrevHash 		//指针前移
			return nil
		})
		return &block
	}

*/








/*
	=====================使用=====================

	func (cli *CLI) PrintBlockChain()  {
		bc:=cli.bc
		//创建迭代器
		it := bc.NewIterator()
		//调用迭代器，返回每一个区块数据
		for {
			block := it.Next() //返回区块，左移

			fmt.Printf("\n\n=====================================当前区块哈希值：%x\n", block.Hash)
			fmt.Printf("版本号：%d\n", block.Version)
			fmt.Printf("前区块哈希值：%x\n", block.PrevHash)
			fmt.Printf("MerkelRoot：%x\n", block.MerkleRoot)
			fmt.Printf("时间戳：%d\n", block.TimeStamp)
			fmt.Printf("难度值：%d\n", block.Difficulty)
			fmt.Printf("Nonce：%d\n", block.Nonce)
			fmt.Printf("区块数据：%s\n\n", block.Data)

			if len(block.PrevHash) == 0 {
				fmt.Println("区块链遍历结束")
				break
			}
		}
	}
 */