package main

import "fmt"

func main() {
	//调用两个函数
	PlayGame()
	Play()
}

//函数1
func PlayGame() {
	fmt.Println("fun 1")
}

//函数2
func Play() {
	fmt.Println("fun 22")
}
