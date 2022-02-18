package main

import (
	"fmt"
	"testing"
)

type Change interface {
	ChangeName(newName string)
	ChangeAge(newAge int)
}

type Student struct {
	Name string
	Age  int
}

func (s *Student) ChangeName(newName string) { //这里是针对指针类型的结构体
	s.Name = newName
}

func (s Student) ChangeAge(newAge int) { //这里直接是结构体
	s.Age = newAge
}

func TestVal(t *testing.T) {
	var stdChg Change //定义stdChg是Change接口
	// stdChg = Student{} //当实现接口的方法有一个是在对象指针上时，只能用对象指针作为值赋值给接口
	stdChg = &Student{
		Name: "Tom",
		Age:  0,
	}

	fmt.Println(stdChg)
}

// func TestStd(teacher *testing.T){
// 	s := Student{Name:"Tom"}
// fmt.Println(s.Name)
// s.ChangeName("Jerry")
// fmt.Println(s.Name)
//}
