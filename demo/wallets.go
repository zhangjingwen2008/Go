package main

import (
	"bytes"
	"crypto/elliptic"
	"encoding/gob"
	"io/ioutil"
	"log"
)

//定一个Wallets结构，保存所有wallet与地址
type Wallets struct {
	WalletsMap map[string]*Wallet //map[地址]钱包
}

//创建方法
func NewWallets() *Wallets {
	var ws Wallets
	//ws.WalletsMap=make(map[string]*Wallet)
	ws.loadFile()
	return &ws
}

func (ws *Wallets) CreateWallet() string {
	wallet := NewWallet()
	address := wallet.NewAddress()

	//wallets.WalletsMap=make(map[string]*Wallet)
	ws.WalletsMap[address] = wallet

	ws.saveToFile()
	return address
}

//保存方法，把新建wallet添加进去
func (ws *Wallets) saveToFile() {

	var buffer bytes.Buffer
	gob.Register(elliptic.P256()) //需要把类型为interface的在编码前进行注册，否则会报错
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(ws)
	if err != nil {
		log.Panic(err)
	}

	ioutil.WriteFile("wallet.dat", buffer.Bytes(), 0600)
}

//读取文件方法，把所有wallet读出来
func (ws *Wallets) loadFile() {

	//读取内容
	content, err := ioutil.ReadFile("wallet.dat")
	if err != nil {
		log.Panic(err)
	}

	//解码
	gob.Register(elliptic.P256())
	decoder := gob.NewDecoder(bytes.NewReader(content))
	var wsLocal Wallets
	err = decoder.Decode(wsLocal)
	if err != nil {
		log.Panic(err)
	}

	//对于结构来说，里面有map的，要指定赋值，不要在最外层直接赋值
	ws.WalletsMap = wsLocal.WalletsMap
}

func (ws *Wallets) ListAllAddresses() []string {
	var addresses []string
	for address := range ws.WalletsMap {
		addresses = append(addresses, address)
	}
	return addresses

}
