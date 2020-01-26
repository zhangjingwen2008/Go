package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
)

/*
	Go中使用消息认证码
		- 所在包： crypto/hmac

		第1步：创建接口
			func New(h func() hash.Hash, key []byte) hash.Hash
			- 返回值: hash接口
			- 参数1: 函数函数的函数名
				sha1.new
				md5.new
				sha256.new
			- 参数2: 秘钥

		第2步: 添加数据
			type Hash interface {
				// 通过嵌入的匿名io.Writer接口的Write方法向hash中添加更多数据，永远不返回错误
				io.Writer
				// 返回添加b到当前的hash值后的新切片，不会改变底层的hash状态
				Sum(b []byte) []byte
				// 重设hash为无数据输入的状态
				Reset()
				// 返回Sum会返回的切片的长度
				Size() int
				// 返回hash底层的块大小；Write方法可以接受任何大小的数据，
				// 但提供的数据是块大小的倍数时效率更高
				BlockSize() int
			}
			type Writer interface {
				Write(p []byte) (n int, err error)
			}

		第3步: 计算散列值
*/

func main() {
	src := []byte("消息认证码的测试数据哇哇哇")
	key := []byte("1234abcd")
	hmacText := GenerateHmac(src, key)
	verify := VerifyHmac(src, key, hmacText)

	fmt.Println("消息认证码 ", hmacText)
	fmt.Println("测试结果 ", verify)
}

//生成消息验证码
func GenerateHmac(plainText, key []byte) []byte {
	//1.创建哈希接口，需要指定使用的哈希算法，和秘钥
	myHash := hmac.New(sha1.New, key)
	//2.给哈希对象添加数据
	myHash.Write(plainText)
	//3.计算散列值
	hashText := myHash.Sum(nil)

	return hashText
}

//验证消息认证码
func VerifyHmac(plainText, key, InHashTest []byte) bool {
	//1.创建哈希接口，需要指定使用的哈希算法，和秘钥
	myHash := hmac.New(sha1.New, key)
	//2.给哈希对象添加数据
	myHash.Write(plainText)
	//3.计算散列值
	hashText := myHash.Sum(nil)
	//4.比较两个哈希值
	return hmac.Equal(hashText, InHashTest)
}
