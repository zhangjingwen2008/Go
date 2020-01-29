package main

import (
	"fmt"
	"os"
	"strconv"
)

//这是一个用来接收命令行参数并且控制区块链操作的文件
type CLI struct {
	bc *BlockChain
}

const Usage = `
	printChain					"print all blockchain data"
	printChainR					"反向打印区块链"
	getBalance --address ADDRESS	"获取指定地址的余额"
	send FROM TO AMOUNT MINER DATA 	"由FROM转AMOUNT给TO，由MINER挖矿，同时写入DATA"
	newWallet					"创建一个新的钱包（私钥公钥对）"
`

//接收参数的动作，我们放到一个函数中
func (cli *CLI) Run() {
	//1.得到所有的命令
	args := os.Args
	if len(args) <2 {
		fmt.Println(Usage)
		return
	}

	//2.分析命令
	cmd:=args[1]
	switch cmd {
		case "printChain":
			//打印区块
			fmt.Println("打印区块")
			cli.PrintBlockChain()
		case "getBalance":
			//打印区块
			fmt.Println("获取余额")
			if len(args)!=4 && args[2]=="--address" { //确保命令有效
				address:=args[3]
				cli.GetBalance(address)
			}
		case "send":
			if len(args)!=7{
				fmt.Println("参数个数错误")
				fmt.Println(Usage)
				return
			}
			from:=args[2]
			to:=args[3]
			amount,_:=strconv.ParseFloat(args[4],64)		//字符串转float64
			miner:=args[5]
			data:=args[6]
			cli.Send(from,to,amount,miner,data)
		case "newWallet":
			fmt.Println("创建新的钱包...")
			cli.NewWallet()
		default:
			fmt.Println("无效命令")
			fmt.Println(Usage)
	}
}

