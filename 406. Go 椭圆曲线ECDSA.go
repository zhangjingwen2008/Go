package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha1"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
)

func main() {
	GenerateEccKey()
	src := []byte("ECC签名测试数据")
	rText, sText := EccSignature(src, "406.EccPrivate.pem")
	ver := EccVerity(src, rText, sText, "406.EccPublic.pem")

	fmt.Println("rText: ", rText)
	fmt.Println("sText: ", sText)
	fmt.Println(ver)
}

//生成密钥对
func GenerateEccKey() {
	//1.使用ecdsa生成密钥对
	privateKey, err := ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	if err != nil {
		panic(err)
	}
	//2.将私钥写入磁盘
	derText, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		panic(err)
	}
	block := pem.Block{
		Type:  "Ecc Private Key",
		Bytes: derText,
	}
	file, err := os.Create("406.EccPrivate.pem")
	if err != nil {
		panic(err)
	}
	pem.Encode(file, &block)
	file.Close()

	//3.将公钥写入磁盘
	publicKey := privateKey.PublicKey
	derText, err = x509.MarshalPKIXPublicKey(&publicKey) //注意这里要传入的是地址，而不是值
	if err != nil {
		panic(err)
	}
	block = pem.Block{
		Type:  "Ecc Public Key",
		Bytes: derText,
	}
	file, err = os.Create("406.EccPublic.pem")
	if err != nil {
		panic(err)
	}
	pem.Encode(file, &block)
	file.Close()
}

//Ecc签名
func EccSignature(plainText []byte, fileName string) (rText, sText []byte) {
	//1. 打开私钥文件, 将内容读出来 ->[]byte
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	info, err := file.Stat()
	if err != nil {
		panic(err)
	}
	buf := make([]byte, info.Size())
	file.Read(buf)
	file.Close()

	//2. 使用pem进行数据解码 -> pem.Decode()
	block, _ := pem.Decode(buf)

	//3. 使用x509, 对私钥进行还原
	privateKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}

	//4. 对原始数据进行哈希运算 -> 散列值
	hashText := sha1.Sum(plainText)

	//5. 进行数字签名
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hashText[:])
	if err != nil {
		panic(err)
	}

	//6. 对r和s进行序列化
	rText, err = r.MarshalText()
	if err != nil {
		panic(err)
	}
	sText, err = s.MarshalText()
	if err != nil {
		panic(err)
	}

	return
}

//ECC签名认证
func EccVerity(plainText, rText, sText []byte, fileName string) bool {
	//1. 打开公钥文件, 将里边的内容读出 -> []byte
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	info, err := file.Stat()
	if err != nil {
		panic(err)
	}
	buf := make([]byte, info.Size())
	file.Read(buf)
	file.Close()

	//2. pem解码 -> pem.Decode()
	block, _ := pem.Decode(buf)

	//3. 使用x509对公钥还原
	publicInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}

	//4. 将接口 -> 公钥
	publicKey := publicInterface.(*ecdsa.PublicKey)

	//5. 对原始数据进行哈希运算 -> 得到散列值
	hashText := sha1.Sum(plainText)

	//对rText和sText进行反序列化
	var r, s big.Int
	r.UnmarshalText(rText)
	s.UnmarshalText(sText)

	//6. 签名的认证 - > ecdsa
	ver := ecdsa.Verify(publicKey, hashText[:], &r, &s)

	return ver
}
