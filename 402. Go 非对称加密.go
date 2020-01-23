package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"os"
)

//生成RSA的密钥对，并且保存到磁盘文件中
func GenerateRsaKey(keySize int) {
	//=================生成私钥=================
	//1.使用RSA中的GenerateKey方法生成私钥
	privateKey, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		panic(err)
	}
	//2.通过x509标准将得到的RSA私钥序列化成为ASN.1的DER编码字符串
	derText := x509.MarshalPKCS1PrivateKey(privateKey) //使用PKCS1算法
	//3.要组织一个pem.Block
	block := pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derText,
	}
	//4.pem编码
	file, err := os.Create("402.private.pem")
	if err != nil {
		panic(err)
	}
	pem.Encode(file, &block)
	file.Close()

	//=================生成公钥=================
	//1.从私钥中取出公钥
	publicKey := privateKey.PublicKey
	//2.使用x509标准序列化
	derStream, err := x509.MarshalPKIXPublicKey(&publicKey) //使用PKIX算法
	if err != nil {
		panic(err)
	}
	//3.将得到的数据放到pem.Block中
	block = pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: derStream,
	}
	//4.pem编码
	file2, err := os.Create("402.public.pem")
	if err != nil {
		panic(err)
	}
	pem.Encode(file2, &block)
	file2.Close()
}

func main() {
	GenerateRsaKey(1024)
	src := []byte("测试RSA公私钥的加解密")
	cipherText := RSAEncrypt(src, "402.public.pem")
	plainText := RSADecrypt(src, "402.private.pem")
	fmt.Printf("原文数据：%s\n加密后的数据：%s\n解密后的数据：%s\n\n", string(src), hex.EncodeToString(cipherText), string(plainText))

}

//RSA加密，公钥加密
func RSAEncrypt(plainText []byte, fileName string) []byte {
	//1.打开文件，并且读出文件内容
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	fileInfo, err := file.Stat()
	if err != nil {
		panic(err)
	}
	buf := make([]byte, fileInfo.Size())
	file.Read(buf)
	file.Close()
	//2.pem解码
	block, _ := pem.Decode(buf)
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes) //使用PKIX算法
	pubKey := pubInterface.(*rsa.PublicKey)
	//3.使用公钥加密
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, pubKey, plainText)
	if err != nil {
		panic(err)
	}

	return cipherText
}

//RSA解密
func RSADecrypt(cipherText []byte, fileName string) []byte {
	//1.打开文件，并且读出文件内容
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	fileInfo, err := file.Stat()
	if err != nil {
		panic(err)
	}
	buf := make([]byte, fileInfo.Size())
	file.Read(buf)
	file.Close()
	//2.pem解码
	block, _ := pem.Decode(buf)
	privKey, err := x509.ParsePKCS1PrivateKey(block.Bytes) //使用PKCS1算法
	//3.使用公钥加密
	plainText, err := rsa.DecryptPKCS1v15(rand.Reader, privKey, cipherText)
	if err != nil {
		panic(err)
	}

	return plainText
}
