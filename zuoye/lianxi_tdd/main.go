package main

import (
	"fmt"
	"learn.go/zuoye/lianxi_tdd/calcbft"
	"learn.go/zuoye/lianxi_tdd/calcbmi"
	"learn.go/zuoye/lianxi_tdd/healthsug"
	"learn.go/zuoye/lianxi_tdd/putin"
)

func main() {
	//输入
	name, age, sex, tall, weight, err := putin.PutinAll()
	if err != nil {
		fmt.Println(err)
	}
	//计算bmi
	bmi, err := calcbmi.CalcBMI(tall, weight)
	if err != nil {
		fmt.Println(err)
	}
	//计算bft
	bft, err := calcbft.CalcBFT(bmi, age, sex)
	if err != nil {
		fmt.Println(err)
		//给出建议
	}
	sug := healthsug.HeslthSug(sex, age, bft)
	fmt.Printf("%s的体脂率是%f,健康建议是：%s。", name, bft, sug)
}
