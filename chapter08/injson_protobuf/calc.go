package main

import (
	gobmi "github.com/armstrongli/go-bmi"
	"learn.go/pkg/apis"
	"log"
)

type Calc struct {
	continental string // 这个有啥用？
}

func (c *Calc) BMI(person *apis.PersonalInfomation) (bmi float64, err error) {
	//bmi, err = gobmi.BMI(client.Weight, client.Tall)// 演示直接使用JSON格式
	bmi, err = gobmi.BMI(float64(person.Weight), float64(person.Tall)) // 演示使用protobuf格式化后，JSON格式
	// protobuf中是 float tall = 3; float weight = 4; 所以需要把float类型转换为float64类型
	if err != nil {
		log.Println("error when calculating bmi:", err)
		return -1, err
	}
	return bmi, nil
}

func (c *Calc) FatRate(person *apis.PersonalInfomation) (fatRate float64, err error) {
	bmi, err := c.BMI(person)
	if err != nil {
		return -1, err
	}
	return gobmi.CalcFatRate(bmi, int(person.Age), person.Sex), nil

}
