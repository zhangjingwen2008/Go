package User

import "fmt"

func Login() {
	fmt.Println("User Login")
}

//方法名称第1个字母为小写，则只能在本文件中调用。为大写，其他文件才可调用
func GetUser() {
	fmt.Println("Get User Info")
}
