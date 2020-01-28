package main

import (
	"fmt"
	"os"
)

//这是一个用来接收命令行参数并且控制区块链操作的文件
type CLI struct {
	bc *BlockChain
}

const Usage = `
	addBlock --data DATA		"add data to blockchain"
	printChain					"print all blockchain data"
	printChainR					"反向打印区块链"
	getBalance --address ADDRESS	"获取指定地址的余额"
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
		case "addBlock":
			//增加区块
			fmt.Println("添加区块")
			if len(args)!=4 && args[2]=="--data"{		//确保命令有效
				//获取命令的数据
				//a.获取数据
				data:=args[3]
				//b.使用bc添加区块AddBlock
				cli.AddBlock(data)
			}else{
				fmt.Println("添加区块参数使用不当，请检查")
			}
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
		default:
			fmt.Println("无效命令")
			fmt.Println(Usage)
	}
}


