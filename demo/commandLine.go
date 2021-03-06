package main

import (
	"fmt"
	"time"
)

//正向打印
func (cli *CLI) PrintBChain() {
	bc := cli.bc
	//创建迭代器
	it := bc.NewIterator()
	//调用迭代器，返回每一个区块数据
	for {
		block := it.Next() //返回区块，左移

		fmt.Printf("\n\n=====================================当前区块哈希值：%x\n", block.Hash)
		fmt.Printf("版本号：%d\n", block.Version)
		fmt.Printf("前区块哈希值：%x\n", block.PrevHash)
		fmt.Printf("MerkelRoot：%x\n", block.MerkleRoot)
		timeFormat := time.Unix(int64(block.TimeStamp), 0).Format("2006-01-02 15:04:05")
		fmt.Printf("时间戳：%s\n", timeFormat)
		fmt.Printf("难度值：%d\n", block.Difficulty)
		fmt.Printf("Nonce：%d\n", block.Nonce)
		fmt.Printf("区块数据：%s\n\n", block.Transactions[0].TXInput[0].PubKey)

		if len(block.PrevHash) == 0 {
			fmt.Println("区块链遍历结束")
			break
		}
	}
}

func (cli *CLI) GetBalance(address string) {

	//1.校验地址
	if !IsValidAddress(address) {
		fmt.Println("地址无效：", address)
		return
	}

	//2.生成公钥哈希
	pubKeyHash := GetPubKeyFromAddress(address)

	utxos := cli.bc.FindUTXOs(pubKeyHash)
	total := 0.0
	for _, utxo := range utxos {
		total += utxo.Value
	}
	fmt.Println(address, "的余额为：", total)
}

func (cli *CLI) Send(from, to string, amount float64, miner, data string) {
	//fmt.Println("from:", from)
	//fmt.Println("to:", to)
	//fmt.Println("amount:", amount)
	//fmt.Println("mincer:", miner)
	//fmt.Println("data:", data)

	//校验地址
	if !IsValidAddress(from){
		fmt.Println("地址无效：",from)
		return
	}
	if !IsValidAddress(to){
		fmt.Println("地址无效：",to)
		return
	}
	if !IsValidAddress(miner){
		fmt.Println("地址无效：",miner)
		return
	}

	//1.创建挖矿交易
	coinbase := NewCoinbaseTX(miner, data)
	//2.创建一个普通交易
	tx := NewTransaction(from, to, amount, cli.bc)
	if tx == nil {
		return
	}
	//3.添加到区块
	cli.bc.AddBlock([]*Transaction{coinbase, tx})
	fmt.Println("转账成功")
}

func (cli *CLI) NewWallet() {
	ws := NewWallets()
	address := ws.CreateWallet()
	fmt.Println("钱包地址：", address)
	//wallet:=NewWallet()
	//address:=wallet.NewAddress()
	//ws:=NewWallets()
	//for address:=range ws.WalletsMap{
	//	fmt.Println("地址：",address)
	//}

}

func (cli *CLI) ListAddresses() {
	ws := NewWallets()
	addresses := ws.ListAllAddresses()
	for _, address := range addresses {
		fmt.Println("地址：", address)
	}
}
