package main

import "fmt"

var s int = 1

func main() {
	Cal(6)
	fmt.Println(s)
}

func Cal(n int) {
	if n == 1 {
		return
	}
	s = s * n
	Cal(n - 1)
}
