package main

import "fmt"

type Student struct {
	id   int
	name string
	age  int
	addr string
}

func main() {
	Stu := make([]Student, 3)
	InitData(Stu)
	GetMax(Stu)
}

//初始化信息
func InitData(Stu []Student) {
	for i := 0; i < len(Stu); i++ {
		fmt.Printf("请输入第%d名学生的信息:\n", i+1)
		fmt.Scan(&Stu[i].id, &Stu[i].name, &Stu[i].age, &Stu[i].addr)
	}
}

//获取最大值
func GetMax(Stu []Student) {
	MaxAge := Stu[0].age
	MaxIndex := 0
	for i := 0; i < len(Stu); i++ {
		if Stu[i].age > MaxAge {
			MaxAge = Stu[i].age
			MaxIndex = i
		}
	}
	fmt.Println("年龄最大：", Stu[MaxIndex])
}
