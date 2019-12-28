package main

import (
	"fmt"
	"strings"
)

func main() {
	TestStr()
}

func TestStr() {
	var enterDate, sentence string

	//1.输入一段日期，把横杠-换成年月日
	fmt.Println("Enter a Date:")
	fmt.Scan(&enterDate)
	str := strings.Split(enterDate, "-")
	fmt.Println(str[0] + "年" + str[1] + "月" + str[2] + "日")

	//2.将一段话中的“邪恶”替换成“**”
	fmt.Println("Enter a sentence:")
	fmt.Scan(&sentence)
	if strings.Contains(sentence, "邪恶") {
		sentence = strings.Replace(sentence, "邪恶", "**", -1)
	}
	fmt.Println(sentence)

}
