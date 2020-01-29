package main


//定一个Wallets结构，保存所有wallet与地址
type Wallets struct{
	WalletsMap map[string]*Wallet			//map[地址]钱包
}

//创建方法
func  NewWallets() *Wallets {
	wallet:=NewWallet()
	address:=wallet.NewAddress()

	var wallets Wallets
	wallets.WalletsMap=make(map[string]*Wallet)
	wallets.WalletsMap[address]=wallet

	return &wallets
}

//保存方法，把新建wallet添加进去

//读取文件方法，把所有wallet读出来
