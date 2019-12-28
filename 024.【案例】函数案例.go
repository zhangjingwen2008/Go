package main

import "fmt"

func main() {
	var username, password, email string
	fmt.Println("Enter Username:")
	fmt.Scan(&username)
	fmt.Println("Enter Password:")
	fmt.Scan(&password)
	fmt.Println("Enter Email:")
	fmt.Scan(&email)

	msg := Validate(username, password, email)
	fmt.Println(msg)
}

func Validate(username string, password string, email string) (msg string) {
	if username == "" || password == "" || email == "" {
		msg = "信息不能为空，用户注册失败"
	} else {
		msg = "用户注册成功"
	}
	return
}
