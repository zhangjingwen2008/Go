package main

import "fmt"

func main() {

	bc := NewBlockChain()
	bc.AddBlock("111111111")
	bc.AddBlock("222222")

	//创建迭代器
	it := bc.NewIterator()
	//调用迭代器，返回每一个区块数据
	for {
		block := it.Next() //返回区块，左移

		fmt.Println("=====================================")
		fmt.Printf("前区块哈希值：%x\n", block.PrevHash)
		fmt.Printf("当前区块哈希值：%x\n", block.Hash)
		fmt.Printf("区块数据：%s\n\n", block.Data)

		if len(block.PrevHash) == 0 {
			fmt.Println("区块链遍历结束")
			break
		}
	}

}
