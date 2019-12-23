package main

import "fmt"

func main() {
	var username string
	var password string

	fmt.Println("Enter Username:")
	fmt.Scan(&username)
	fmt.Println("Enter Password:")
	fmt.Scan(&password)

	if username == "admin" && password == "88888" {
		fmt.Println("登录Success！")
	} else if username == "admin" {
		fmt.Println("Password Error")
	} else if password == "88888" {
		fmt.Println("Username Error")
	} else {
		fmt.Println("Login Failed!")
	}
}
