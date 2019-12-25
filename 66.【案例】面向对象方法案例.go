package main

import "fmt"

type Student struct {
	name    string
	gender  string
	age     int
	chinese float64
	math    float64
	english float64
}

//打招呼
func (s *Student) Hello(name string, age int, gender string) {
	s.name = name
	s.age = age
	s.gender = gender
	fmt.Printf("我叫%s，今年%d岁，%s\n", s.name, s.age, s.gender)
}

//计算总分与平均分
func (s *Student) Calculate(chinese float64, math float64, english float64) {
	s.chinese = chinese
	s.math = math
	s.english = english

	total := s.chinese + s.math + s.english
	average := float64(total) / 3
	fmt.Println("总分", total)
	fmt.Printf("平均分%.1f", average)
}

func main() {
	var stu Student
	stu.Hello("Hunter", 18, "男")
	stu.Calculate(22, 77, 1000)
}
