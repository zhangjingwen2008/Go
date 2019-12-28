package main

import "fmt"

func main() {
	Slice := []int{22, 33, 44, 55, 66}
	Init(Slice)
}

func Init(Slice []int) {
	for i := 0; i < len(Slice); i++ {
		fmt.Println(Slice[i])
	}
}
