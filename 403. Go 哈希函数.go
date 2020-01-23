package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

/*
	go中使用单向散列函数

		第1种方式, 直接调用sum
			- 适用于数据量比较小的情况
			- func Sum(data []byte) [Size]byte

		第2种方式
		1. 创建哈希接口对象
			func New() hash.Hash
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
		2. 往创建出的哈希对象中添加数据
			hash.Hash.Write([]byte("添加的数据..."))
			hash.Hash.Write([]byte("添加的数据..."))
			hash.Hash.Write([]byte("添加的数据..."))
			hash.Hash.Write([]byte("添加的数据..."))
		3. 计算结果, md5就是散列值
			md5 := hash.Sum(nil);
			// 散列值一般是一个二进制的字符串, 有些字符不可见, 需要格式化
			// 格式化为16进制的数字串 - 0-9, a-f
			func EncodeToString(src []byte) string
			// 数据转换完成之后, 长度是原来的2倍
*/

//使用SHA256
func myHash() {
	//1.创建哈希接口对象
	myHash := sha256.New()
	//2.添加数据
	src := []byte("测试加密数据哈哈哈")
	myHash.Write(src)
	myHash.Write(src)
	myHash.Write(src)
	//3.计算结果
	res := myHash.Sum(nil)
	//4.格式化为16进制形式
	myStr := hex.EncodeToString(res)
	fmt.Printf("%s\n", myStr)

	//sha256.Sum256([]byte("测试加密数据哈哈哈"))			//上面四步操作统一成1句，即第1种方式
}

func main() {
	myHash()
}
