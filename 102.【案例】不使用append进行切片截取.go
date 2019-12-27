package main

import "fmt"

func main() {
	str := []string{"red", "", "blue", "", "", "green"}
	af := Test(str)
	fmt.Println(af)
}

func Test(str []string) []string {
	i := 0
	for _, value := range str {
		if value != "" {
			str[i] = value
			i++
		}
	}
	return str[:i]
}
