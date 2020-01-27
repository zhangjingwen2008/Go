package main

//4.引入区块链
type BlockChain struct {
	blocks []*Block //定义一个区块链数组
}

//5.定义一个区块链
func NewBlockChain() *BlockChain {
	genesisBlock := GenesisBlock() //创建一个创世块，并作为第一个区块添加到区块链中
	return &BlockChain{
		blocks: []*Block{genesisBlock},
	}
}

//定义一个创世区块
func GenesisBlock() *Block {
	return NewBlock("这是创世区块哇", []byte{})
}

//6.添加区块
func (bc *BlockChain) AddBlock(data string) {
	//获取前一区块的哈希值
	preBlock:=bc.blocks[len(bc.blocks)-1]
	preHash:=preBlock.Hash

	//a.创建新的区块
	newBlock:=NewBlock(data,preHash)
	//b.添加到区块链数组中
	bc.blocks=append(bc.blocks, newBlock)
}
