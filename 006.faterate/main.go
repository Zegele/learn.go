package main

import "fmt"

func main() {
	var bmi float64
	var weight float64
	fmt.Print("体重（kg）：")
	fmt.Scanner(&weight)

	var height float64 = 1.80
	fmt.Print("身高（米）：")

	bmi = weight / (height * height)

	var age int

}
