package main

func main() {

	bc := NewBlockChain()
	cli:=CLI{bc}
	cli.Run()
	//bc.AddBlock("111111111")
	//bc.AddBlock("222222")
	//


}
