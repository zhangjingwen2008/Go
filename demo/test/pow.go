package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	data := "helloworld"
	for i := 0; i < 100000; i++ {
		result := sha256.Sum256([]byte(data + string(i)))
		fmt.Printf("result:%x %d\n", result[:], i)
	}
}
