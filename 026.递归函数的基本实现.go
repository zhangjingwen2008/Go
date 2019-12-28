package main

import "fmt"

func main() {
	Test(5)
}

func Test(n int) int {
	if n == 1 {
		return 1
	}

	r := Test(n - 1)
	fmt.Println("Next Row:", r)

	return r + 1
}
