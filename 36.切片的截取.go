package main

import "fmt"

func main() {
	Slice := []int{3, 2, 6, 7}

	//截取操作1
	//第一个值0：截取的起始位置
	//第二个值3：截取的终止位置（不包含该值）
	//第三个值4：用来计算容量，指的是新切片中最多能容纳多少元素
	//容量=原切片容量-第1个值
	//长度=第2个值-第1个值
	//注：截取后返回的是新的切片
	Slice1 := Slice[1:3:4]
	fmt.Println(Slice1)      //结果[2 6]
	fmt.Println(cap(Slice1)) //结果3
	fmt.Println(len(Slice1)) //结果2

	//截取操作2
	Sliece2 := Slice[:]       //一个冒号代表将内容全部截取
	fmt.Println(Sliece2)      //结果[3 2 6 7]
	fmt.Println(cap(Sliece2)) //结果4
	fmt.Println(len(Sliece2)) //结果4

	//截取操作3
	Sliece3 := Slice[2:]      //从下标2开始往后面截取
	fmt.Println(Sliece3)      //结果[6 7]
	fmt.Println(cap(Sliece3)) //结果2
	fmt.Println(len(Sliece3)) //结果2

	//截取操作4
	Sliece4 := Slice[:3]      //从下标3开始往前面截取
	fmt.Println(Sliece4)      //结果[3 2 6]
	fmt.Println(cap(Sliece4)) //结果4，相当于截取操作为[0:3]，则容量=4-0=4
	fmt.Println(len(Sliece4)) //结果3

	//截取操作5
	Sliece5 := Slice[1:3]     //截取范围1-3
	fmt.Println(Sliece5)      //结果[2 6]
	fmt.Println(cap(Sliece5)) //结果3
	fmt.Println(len(Sliece5)) //结果2

}
