package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"encoding/hex"
	"fmt"
)

/*
	加密流程：
		1.创建一个底层使用 DES/3DES/AES 的密码接口
			- func NewCipher(key []byte) (cipher.Block, error)					//DES加密 所在包：crypto/des
			- func NewTripleDESCipher(key []byte) (cipher.Block, error)			//3DES加密 所在包：crypto/des
			- func NewCipher(key []byte) (cipher.Block, error)					//AES加密 所在包：crypto/aes
		2.如果使用的是 ECB/CBC 分组模式需要对明文分组进行填充
		3.创建一个密码分组模式的接口对象
			- CBC
				func NewCBCEncrypter(b Block, iv []byte) BlockMode
			- CFB
				func NewCFBEncryter(block Block, iv []byte) Stream
			- OFB
			- CTR
		4.加密，得到密文
*/

func main() {
	//DES加解密
	fmt.Println("DES 加解密 CBC模式......")
	key := []byte("1234abcd")
	src := []byte("这是要加密的数据")
	cipherText := desEncrypt(src, key)
	plainText := desDecrypt(cipherText, key)
	fmt.Printf("原文数据：%s\n加密后的数据：%s\n解密后的数据：%s\n\n", string(src), hex.EncodeToString(cipherText), string(plainText))

	//AES加解密
	fmt.Println("aes 加解密 ctr模式......")
	key = []byte("1234abcd1234abcd")
	cipherText = aesEncrypt(src, key)
	plainText = aesDecrypt(cipherText, key)
	fmt.Printf("原文数据：%s\n加密后的数据：%s\n解密后的数据：%s\n\n", string(src), hex.EncodeToString(cipherText), string(plainText))

}

//DES的CBC加密
//编写填充函数，如果最后一个分组字节数不够，填充
//......字节数刚好合适，添加一个新的分组
//填充的字节的值 == 缺少的字节的数
func paddingLastGroup(plainText []byte, blockSize int) []byte {
	//1.求出最后一个组中剩余的需要填充的字节数
	padNum := blockSize - len(plainText)%blockSize
	//2.创建新的切片，长度==padNum，每个字节值byte(padNum)
	char := []byte{byte(padNum)} //长度1
	//切片创建，并初始化
	newPlain := bytes.Repeat(char, padNum)
	//4.newPlain数组追加到原始明文的后面
	newText := append(plainText, newPlain...)
	return newText
}

//去掉填充的数据
func unPaddingLastGroup(plainText []byte) []byte {
	//1.拿去切片中的最后一个字节
	length := len(plainText)
	lastChar := plainText[length-1]
	number := int(lastChar) //尾部填充的字节个数
	return plainText[:length-number]
}

//DES加密
func desEncrypt(plainText, key []byte) []byte {
	//1.建一个底层使用DES的密码接口
	block, err := des.NewCipher(key)
	if err != nil {
		panic(err)
	}

	//2.明文填充
	newText := paddingLastGroup(plainText, block.BlockSize())

	//3.创建一个使用CBC分组的接口
	iv := []byte("12345678") //初始化向量，要求与blockSize相同，即8字节
	blockMode := cipher.NewCBCEncrypter(block, iv)

	//4.加密
	cipherText := make([]byte, len(newText))
	blockMode.CryptBlocks(cipherText, newText)

	return cipherText
}

//DES解密
func desDecrypt(cipherText, key []byte) []byte {
	//1.建一个底层使用DES的密码接口
	block, err := des.NewCipher(key)
	if err != nil {
		panic(err)
	}

	//2.创建一个使用CBC分组的接口
	iv := []byte("12345678")
	blockMode := cipher.NewCBCDecrypter(block, iv)

	//3.解密
	blockMode.CryptBlocks(cipherText, cipherText)

	//4.明文删除尾部
	plainText := unPaddingLastGroup(cipherText)

	return plainText
}

//AES加密
//无需填充组
func aesEncrypt(plainText, key []byte) []byte {
	//1.建一个底层使用AES的密码接口
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	//2.创建一个使用CTR分组的接口
	//iv := []byte("12345678")
	iv := []byte("12345678abcdefgh") //不是初始化向量，而是作为随机数种子
	stream := cipher.NewCTR(block, iv)

	//3.加密
	cipherText := make([]byte, len(plainText))
	stream.XORKeyStream(cipherText, plainText)

	return cipherText
}

//AES解密
func aesDecrypt(cipherText, key []byte) []byte {
	//1.建一个底层使用AES的密码接口
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	//2.创建一个使用CTR分组的接口
	iv := []byte("12345678abcdefgh") //不是初始化向量，而是作为随机数种子
	stream := cipher.NewCTR(block, iv)

	//3.加密
	stream.XORKeyStream(cipherText, cipherText)

	return cipherText
}
