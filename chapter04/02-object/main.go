package main

import (
	"fmt"
	gobmi "github.com/armstrongli/go-bmi"
)

type Person struct {
	name string
	// Alias []string
	sex    string
	tall   float64
	weight float64
	age    int
}

func main() {
	persons := []Person{ //结构体的切片
		{
			name:   "小强",
			sex:    "男",
			tall:   1.7,
			weight: 70,
			age:    35,
		},
	}
	xq := Person{
		name:   "xiaoqiang",
		sex:    "sex",
		tall:   1.8,
		weight: 65,
		age:    30,
	}
	fmt.Println(xq)
	for _, item := range persons {
		bmi, err := gobmi.BMI(item.weight, item.tall)
		fmt.Println(bmi, err)
	}

	a := new(Person) //new创建的是指针类型的
	a.name = "AAA"
	fmt.Printf("a的类型是：%T,值是：%v\n", a, a)
	b := &Person{} //b是Person结构体的指针
	b.name = "BBB"
	fmt.Printf("b的类型是：%T,值是：%v\n", b, b)
}
