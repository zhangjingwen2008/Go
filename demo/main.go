package main

func main() {

	//var a uint64
	//a=1580113409
	//fmt.Println(a)
	//
	//b:=Uint64ToByte(a)
	//fmt.Println(b)
	//fmt.Println("string:",string(b))

	bc := NewBlockChain()
	bc.AddBlock("A给B了10BTC")
	bc.AddBlock("A给C了10BTC")

	/*
		for i,block:=range bc.blocks{
			fmt.Printf("===================块高度：%d====================\n", i)
			fmt.Printf("前区块哈希值：%x\n", block.PrevHash)
			fmt.Printf("当前区块哈希值：%x\n", block.Hash)
			fmt.Printf("区块数据：%s\n\n", block.Data)
		}

	*/
}
