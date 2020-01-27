package main

import (
	"bolt"
	"log"
)

func main() {
	//1.打开数据库
	db,err:=bolt.Open("test.db",0600,nil)
	if err!=nil{
		log.Panic("打开数据库失败,",err)
	}

	//将要操作数据库（改写）
	db.Update(func(tx *bolt.Tx) error {
		//2.找到抽屉bucket(如果没有，就创建)
		bucket:=tx.Bucket([]byte("b1"))
		if bucket==nil{
			//没有抽屉，我们需要创建一个
			bucket,err=tx.CreateBucket([]byte("b1"))
			if err!=nil{
				log.Panic("创建bucket(b1)失败")
			}
		}

		bucket.Put([]byte("1"),[]byte("hello"))
		bucket.Put([]byte("2"),[]byte("world"))

		return nil
	})
	//3.写数据
	//4.读数据
}
