package main

/*
	//这是一个用来接收命令行参数并且控制区块链操作的文件
	type CLI struct {
		bc *BlockChain
	}

	const Usage = `
		addBlock --data DATA		"add data to blockchain"
		printChain					"print all blockchain data"
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
					cli.bc.AddBlock(data)
				}else{
					fmt.Println("添加区块参数使用不当，请检查")
				}
			case "printChain":
				//打印区块
				fmt.Println("打印区块")
				cli.PrintBlockChain()
			default:
				fmt.Println("无效命令")
				fmt.Println(Usage)
		}
	}
 */



/*
	=====================使用=====================

	func (cli *CLI) AddBlock(data string) {
		cli.bc.AddBlock(data)
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
			fmt.Printf("区块数据：%s\n\n", block.Data)

			if len(block.PrevHash) == 0 {
				fmt.Println("区块链遍历结束")
				break
			}
		}
	}
 */