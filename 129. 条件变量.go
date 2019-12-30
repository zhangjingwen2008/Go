package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
条件变量：本身不是锁，但经常与锁结合使用
定义：var cond sync.Cond
3个常用方法：
	1.func (c *Cond) Wait()
		- 阻塞等待条件变量满足的go程
		- 释放已掌握的互斥锁，相当于cond.L.Unlock()；注意这一步和上面一步，两者为1个原子操作
		- 当被唤醒，Wait()函数返回时，解除阻塞并重新获取互斥锁，相当于cond.L.Lock()
	2.func (c *Cond) Signal()
		- 单发通知，给一个正等待（阻塞）在该条件变量上的go程发送通知
	3.func (c *Cond) Broadcast()
		- 广播通知，给正在等待（阻塞）在该条件变量上的所有go程发送通知
使用流程：
	1.创建条件变量：var cond sync.Cond
	2.指定条件变量用的锁：cond.L = new(sync.Mutex)
	3.cond.L.Lock()	给公共区加锁（互斥量）
	4.判断是否达到阻塞条件（缓冲区满/空）————for循环判断
		for len(ch) == cap(ch){
			cond.Wait()				//1个语句做了3件事：1)阻塞 2)解锁 3)被唤醒后重新加锁
		}
	5.访问公共区————读、写数据、打印
	6.解锁条件变量用的锁 cond.L.Unlock()
	7.唤醒阻塞在条件变量上的对端
*/

var cond sync.Cond //步骤1：创建条件变量，这里设定为全局变量

func main() {
	product := make(chan int, 5)
	rand.Seed(time.Now().UnixNano())

	cond.L = new(sync.Mutex) //步骤2：指定条件变量用的锁，这里使用互斥锁

	for i := 0; i < 5; i++ {
		go producer(product, i+1)
	}

	for i := 0; i < 5; i++ {
		go consumer(product, i+1)
	}

	for {

	}

}

//生产者
func producer(out chan<- int, n int) {
	for {
		cond.L.Lock()       //步骤3：给公共区加锁
		for len(out) == 5 { //步骤4：判断是否达到阻塞条件，不能用if的原因是需要不断的去判断，而if进行1次后就结束了，无法再次触发
			cond.Wait()
		}

		//步骤5：访问公共区进行操作
		num := rand.Intn(900) + 100
		fmt.Printf("----------线程%d生产了%d\n", n, num)
		out <- num

		cond.L.Unlock() //步骤6：给公共区解锁
		cond.Signal()   //步骤7：唤醒阻塞对端
		time.Sleep(time.Millisecond * 300)
	}
}

//消费者
func consumer(in <-chan int, n int) {
	for {
		cond.L.Lock()
		for len(in) == 0 {
			cond.Wait()
		}

		num := <-in
		fmt.Printf("消费者%d获得了%d\n", n, num)

		cond.L.Unlock()
		cond.Signal()
		time.Sleep(time.Millisecond * 300)
	}
}
