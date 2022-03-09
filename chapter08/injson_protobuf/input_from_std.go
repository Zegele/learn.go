package main

import (
	"fmt"
	"learn.go/pkg/apis"
)

type inputFromStd struct {
}

func (inputFromStd) GetInput() *apis.PersonalInfomation {
	//录入各项
	var name string

	fmt.Print("姓名：")
	fmt.Scanln(&name)

	var weight float64
	fmt.Print("体重(千克)：")
	fmt.Scanln(&weight)

	var tall float64
	fmt.Print("身高（米）：")
	fmt.Scanln(&tall)

	var age int
	fmt.Print("年龄：")
	fmt.Scanln(&age)

	sex := "男" //初始值设置成“男”
	fmt.Print("性别（男/女）:")
	fmt.Scanln(&sex)

	return &apis.PersonalInfomation{
		Name:   name,
		Sex:    sex,
		Tall:   float32(tall),
		Weight: float32(weight),
		Age:    int64(age),
	}
}
