package main

import "fmt"

func main() {
	Slice := []int{4, 8, 2, 0, 33}
	for i := 0; i < len(Slice)-1; i++ {
		min := Slice[i]
		minIndex := i
		for j := i + 1; j < len(Slice); j++ {
			if Slice[j] < min {
				min = Slice[j]
				minIndex = j
			}
		}
		if minIndex != i {
			Slice[i], Slice[minIndex] = Slice[minIndex], Slice[i]
		}
	}
	fmt.Println(Slice)
}
