package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

/*
	数字签名和验证步骤
		签名
			1.对原始数据进行哈希运算，得到散列值
			2.使用非对称加密的 私钥 对 散列值 进行加密，得到散列值即签名
			3.将原始数据和签名一并发送给对方
		验证
			1.接收数据，包括原始数据和数字签名
			2.对数字签名使用 公钥 进行解密，得到散列值
			3.对原始数据进行哈希运算得到散列值，与解密的散列值进行比对验证
*/

func main() {
	src := []byte("RSA签名测试数据")
	sig := SignatureRSA(src, "402.private.pem")
	ver := VerifyRSA(src, sig, "402.public.pem")

	fmt.Println("签名：", sig)
	fmt.Println("签名验证后的结果： ", ver)
}

//RSA签名
func SignatureRSA(plainText []byte, fileName string) []byte {
	//1. 打开磁盘的私钥文件
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	//2. 将私钥文件中的内容读出
	info, _ := file.Stat()
	buf := make([]byte, info.Size())
	file.Read(buf)
	file.Close()

	//3. 使用pem对数据解码, 得到了pem.Block结构体变量
	block, _ := pem.Decode(buf)

	//4. x509将数据解析成私钥结构体 -> 得到了私钥
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}

	//5. 创建一个哈希对象 -> md5/sha1 -> sha512
	myHash := sha512.New()

	//6. 给哈希对象添加数据
	myHash.Write(plainText)

	//7. 计算哈希值
	hashText := myHash.Sum(nil)

	//8. 使用rsa中的函数对散列值签名
	sigText, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA512, hashText)
	if err != nil {
		panic(err)
	}

	return sigText
}

//RSA签名验证
func VerifyRSA(plainText, sigText []byte, fileName string) bool {
	//1. 打开公钥文件, 将文件内容读出 - []byte
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

	//2. 使用pem解码 -> 得到pem.Block结构体变量
	block, _ := pem.Decode(buf)

	//3. 使用x509对pem.Block中的Bytes变量中的数据进行解析 ->  得到一接口
	pubInterface, _ := x509.ParsePKIXPublicKey(block.Bytes)

	//4. 进行类型断言 -> 得到了公钥结构体
	publicKey := pubInterface.(*rsa.PublicKey)

	//5. 对原始消息进行哈希运算(和签名使用的哈希算法一致) -> 散列值
	hashText := sha512.Sum512(plainText)

	//6. 签名认证 - rsa中的函数
	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA512, hashText[:], sigText)
	if err != nil {
		return false
	}
	return true
}
