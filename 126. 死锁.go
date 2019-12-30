package main

import "fmt"

/*
死锁：不是一种锁，而是一种错误使用锁导致的情况
	1.单go程自己死锁
		- channel应该在至少 2 个以上的go程中进行通信，否则死锁
	2.go程间channel访问顺序错误导致死锁
		- 使用channel一端读（写）时，要保证另一端
	3.多go程，多channel交叉死锁
	4.隐性死锁（不报错，仅卡住）：在go语言中，尽量不要将互斥锁、读写锁与channel混用
		无缓冲channel时，若在任意一端加上锁，另一端则会因无法取到锁而永远卡住，从而导致锁中的channel也同样被卡住
*/

//死锁1
func main() {
	ch := make(chan int)
	ch <- 890 //在这里就会堵塞死锁，因为没有接收者
	num := <-ch
	fmt.Println(num)
}

//死锁2
func main() {
	ch := make(chan int)
	num := <-ch //这里就会堵塞死锁，因为没有写入者。应该将这句放到本函数的写入端之后
	fmt.Println(num)

	go func() {
		ch <- 888
	}()
}

//死锁3
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	//子go程
	go func() {
		for {
			select {
			case num := <-ch1: //互相不释放资源，导致死锁
				ch2 <- num
			}
		}
	}()

	//主go程
	for {
		select {
		case num := <-ch2:
			ch1 <- num
		}
	}

}
