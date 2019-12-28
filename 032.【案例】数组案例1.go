package main

import "fmt"

func main() {
	var Numbers [5]int = [5]int{2, -6, 4, 22, 1}

	max, min, total, average := Calculage(Numbers)

	fmt.Printf("Max=%d,Min=%d,Total=%d,Average=%.1f", max, min, total, average)
}

func Calculage(n [5]int) (max int, min int, total int, average float64) {
	var current int
	max = n[0]
	min = n[0]

	for i := 0; i < len(n); i++ {
		current = n[i]
		if max < current { //最大值
			max = current
		}
		if min > current { //最小值
			min = current
		}
		total += current //总和
	}
	average = float64(total) / 5 //平均值
	return
}
