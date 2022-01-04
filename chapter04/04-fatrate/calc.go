package main

import (
	gobmi "github.com/armstrongli/go-bmi"
	"log"
)

type Calc struct { //用意是可能不同洲会有不同的算法。
	continental string
}

func (Calc) BMI(person *Person) error { //针对Calc的方法， BMI的参数也是个结构图体。可能只用到其中的参数，但写的时候写结构体即可。
	bmi, err := gobmi.BMI(person.weight, person.tall)
	if err != nil {
		log.Println("error when calculating bmi:", err)
		return err
	}
	person.bmi = bmi
	return nil
}

func (Calc) FatRate(person *Person) error {
	person.fatRate = gobmi.CalcFatRate(person.bmi, person.age, person.sex)
	return nil
}
