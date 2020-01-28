package main

import (
	"fmt"
)

func (cli *CLI) AddBlock(data string) {
	//cli.bc.AddBlock(data)		TODO
	fmt.Println("添加区块成功！")
}

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
		fmt.Printf("区块数据：%s\n\n", block.Transactions[0].TXInput[0].Sig)

		if len(block.PrevHash) == 0 {
			fmt.Println("区块链遍历结束")
			break
		}
	}
}

func (cli *CLI) GetBalance(address string)  {
	utxos:=cli.bc.FindUTXOs(address)
	total:=0.0
	for _,utxo:=range utxos{
		total+=utxo.value
	}
	fmt.Println(address,"的余额为：",total)
}