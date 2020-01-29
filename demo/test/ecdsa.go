package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"log"
	"math/big"
)

//1.演示如何使用ecdsa生成公钥私钥
//2.签名校验

func main() {
	//===================签名===================
	//创建曲线
	curve:=elliptic.P256()
	//生成私钥
	privateKey,err:=ecdsa.GenerateKey(curve,rand.Reader)
	if err!=nil{
		log.Panic(err)
	}

	//生成公钥
	publicKey:=privateKey.PublicKey

	data:="hello world"
	hash:=sha256.Sum256([]byte(data))

	//签名
	r,s,err:=ecdsa.Sign(rand.Reader,privateKey,hash[:])
	if err!=nil{
		log.Panic(err)
	}

	//把r,s进行序列化传输
	signature:=append(r.Bytes(),s.Bytes()...)

	fmt.Println("公钥：",publicKey)
	fmt.Println("签名：",signature)

	//===================验证===================
	//1.定义两个辅助的big.int
	r1:=big.Int{}
	s1:=big.Int{}
	//2.拆分signature，前半部分给r，后半部分给s
	r1.SetBytes(signature[0:len(signature)/2])
	s1.SetBytes(signature[len(signature)/2:])
	//3.校验
	res:=ecdsa.Verify(&publicKey,hash[:],&r1,&s1)
	fmt.Println("校验结果：",res)

}
