package main

import (
	"fmt"
)

//联系人结构体
type Person struct {
	userName     string
	addressPhone map[string]string
}

//联系人列表切片
var personList = make([]Person, 0)

func main() {
	for {
		ScanNum()
	}
}

func ScanNum() {
	fmt.Println("添加1")
	fmt.Println("删除2")
	fmt.Println("查询3")
	fmt.Println("编辑4")

	var num int
	fmt.Scan(&num)
	SwitchType(num)
}

func SwitchType(n int) {
	switch n {
	case 1:
		AddPerson()
	case 2:
		RemovePerson()
	case 3:
		FindPerson()
	case 4:
		EditPerson()
	}
}

//增加联系人
func AddPerson() {
	var name, phoneType, phone string
	addressPhone := make(map[string]string)
	var AddEnter string

	fmt.Println("Enter Name:")
	fmt.Scan(&name)

	//循环输入联系方式
	for {
		fmt.Println("Enter PhoneType:")
		fmt.Scan(&phoneType)
		fmt.Println("Enter Phone:")
		fmt.Scan(&phone)
		addressPhone[phoneType] = phone //联系方式添加进map

		//退出操作
		fmt.Println("Press 'Q' to Exit!")
		fmt.Scan(&AddEnter)
		if AddEnter == "Q" {
			break
		} else {
			continue
		}
	}

	//联系人列表切片新增
	personList = append(personList, Person{userName: name, addressPhone: addressPhone})
	ShowPersonList()
}

//展示切片中联系人信息
func ShowPersonList() {
	if len(personList) == 0 {
		fmt.Println("目前没有联系人信息")
	} else {
		for _, value := range personList {
			fmt.Println("姓名：", value.userName)
			for k, v := range value.addressPhone {
				fmt.Printf("%s:%s\n", k, v)
			}
		}
	}
}

//删除联系人信息
func RemovePerson() {
	var name string
	var index int = -1

	//1.输入要删除联系人的姓名
	fmt.Println("Enter username you want delete:")
	fmt.Scan(&name)

	//2.遍历列表找到姓名对应下标
	for i := 0; i < len(personList); i++ {
		if personList[i].userName == name {
			index = i
			break
		}
	}

	//3.删除操作：切片的截取
	if index != -1 {
		personList = append(personList[:index], personList[index+1:]...)
	}

	ShowPersonList()
}

//获得联系人信息
func FindPerson() *Person {
	var name string
	var index int = -1

	//1.输入要查询的联系人名称
	fmt.Println("Enter username you want get:")
	fmt.Scan(&name)

	//2.循环遍历List获得联系人信息
	for key, value := range personList {
		if value.userName == name {
			fmt.Printf("联系人姓名：%s\n", value.userName)
			for k, v := range value.addressPhone {
				fmt.Printf("%s:%s\n", k, v)
			}
			index = key
			break
		}
	}

	//3.返回目标信息
	if index != -1 {
		return &personList[index]
	} else {
		fmt.Println("无此联系人信息")
		return nil
	}
}

//编辑联系人信息
func EditPerson() {
	var name string
	var EditEnter int
	var EditPhoneEnter int
	var NewPhone string
	menu := make([]string, 0)

	//1.输入要编辑的联系人姓名
	var person *Person
	person = FindPerson()

	//2.编辑操作
	if person != nil {

		for {
			fmt.Println("编辑用户5")
			fmt.Println("编辑电话6")
			fmt.Println("退出7")
			fmt.Scan(&EditEnter)
			switch EditEnter {
			case 5:
				fmt.Println("Enter new name:")
				fmt.Scan(&name)
				person.userName = name
			case 6:
				var j int
				for key, value := range person.addressPhone {
					fmt.Println("编辑(", key, ")", value, "请按", j)
					menu = append(menu, key)
					j++
				}
				fmt.Scan(&EditPhoneEnter)
				//编辑操作
				for k, v := range menu {
					if EditPhoneEnter == k {
						fmt.Println("Enter New Phone:")
						fmt.Scan(&NewPhone)
						person.addressPhone[v] = NewPhone
					}
				}
			}
			if EditEnter == 7 {
				break
			}
		}
	}

	ShowPersonList()
}
