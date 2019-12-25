package main

import "fmt"

type Person struct {
	name   string
	age    int
	gender string
}

func (p *Person) SetValue(name string, age int, gender string) {
	p.name = name
	p.age = age
	p.gender = gender
}

type Reporter struct {
	Person
	hobby string
}

func (r *Reporter) RepSayHello(hobby string) {
	fmt.Printf("我叫%s,爱好%s,年龄%d，是个%s狗仔\n", r.name, r.hobby, r.age, r.gender)
}

type Programmer struct {
	Person
	workDay int
}

func (p *Programmer) ProSayHello(workDay int) {
	p.workDay = workDay
	fmt.Printf("我叫%s,年龄%d，性别%s，工作年限%d年\n", p.name, p.age, p.gender, p.workDay)
}

func main() {
	var reporter Reporter
	var programmer Programmer
	reporter.SetValue("张三", 34, "男")
	programmer.SetValue("孙权", 23, "男")
	reporter.RepSayHello("偷拍") //结果：我叫张三,爱好,年龄34，是个男狗仔
	programmer.ProSayHello(3)  //结果：我叫孙权,年龄23，性别男，工作年限3年
}
