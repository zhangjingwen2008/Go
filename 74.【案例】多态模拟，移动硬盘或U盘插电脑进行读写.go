package main

import "fmt"

//接口
type Storager interface {
	Read()
	Write()
}

//移动硬盘
type MDisk struct {
}

func (m *MDisk) Read() {
	fmt.Println("移动硬盘读操作")
}
func (m *MDisk) Write() {
	fmt.Println("移动硬盘写操作")
}

//U盘
type UDisk struct {
}

func (u *UDisk) Read() {
	fmt.Println("U盘读操作")
}
func (u *UDisk) Write() {
	fmt.Println("U盘写操作")
}

//公共类，用于多态
func Computer(s Storager) {
	s.Read()
	s.Write()
}

func main() {
	var m MDisk
	var u UDisk
	Computer(&m)
	Computer(&u)
}
