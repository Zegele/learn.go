package main

import (
	"fmt"
	"time"
)

var tall float64
var weight float64

func main() {
	//panicAndRecover()

	//deferGuess()
	//fmt.Println("sleep somewhile")
	//time.Sleep(10 * time.Second)

	//openFile()
	//fmt.Println("sleep somewhile")
	//time.Sleep(10 * time.Sleep())

	closureMain()
	fmt.Println("sleep somewhile")
	time.Sleep(10 * time.Second)

	//close call
	fmt.Println("1, 2, 3, 4, 5, 6, 7, 8, 9, 0")
	fmt.Println("1, 2, 3, 4, 5, 6, 7, 8, 9, 0")
	fmt.Println("1, 2, 3, 4, 5, 6, 7, 8, 9, 0")
	fmt.Println("1, 2, 3, 4, 5, 6, 7, 8, 9, 0")
	fmt.Println("1, 2, 3, 4, 5, 6, 7, 8, 9, 0")
	fmt.Println("1, 2, 3, 4, 5, 6, 7, 8, 9, 0")

	showUsedTimes()
	calcSum()
	//guess(1, 100)
	//fmt.Println("done guess, sleep somewhile")
	//time.Sleep(10 * time.Second)

	fmt.Println(fib(10)) // 斐波那契数列 递归练习

	//fmt.Println("done 04-fatrate, sleep somewhile")
	//time.Sleep(10 * time.Second)

	sampleSubdomain2()
	fmt.Print("全局变量赋值前：")
	calcAdd() //0
	tall, weight = 1.80, 66.0
	fmt.Print("全局变量赋值后：")
	calcAdd() //67.8

	//重新定义重名的局部变量
	tall, weight := 100.00, 70.00
	calcAdd() // ? 67.8  因为calcAdd中的tall和weight是函数内部的，引用的变量需要是全局变量。
	tall, weight = 200.0, 70.0
	calcAdd() //? 67.8

	calculatorAdd := func(a, b float64) float64 {
		return a + b
	}

	result := calculatorAdd(1, 3)
	fmt.Println(result)

	{
		//fmt.Scanln...
		personTall := 1.81
		personWeight := 90.0
		calculatorAdd(personTall, personWeight)
		//suggestions
	}

	{
		//fmt.Scanln...
		personTall := 1.81
		personWeight := 90.0
		calculatorAdd(personTall, personWeight)
		//suggestions
	}
	// fmt.Println(personTall) // personTall 的有效范围 { } 内部。外部无效。
	fmt.Println(tall, weight)
}

func funcDef(nums ...int) (addResult int) {
	for _, item := range nums {
		addResult += item
	}
	return
}

func funcDef1(nums ...int) int {
	sum := 0
	for _, item := range nums {
		sum += item
	}
	return sum
}

func calcAdd() float64 {
	fmt.Println(tall + weight) //这种函数要注意
	return tall + weight
}

func sampleSubdomain() {
	name := "小强"              // 声明变量 name, 值是“小强”
	fmt.Println("名字是：", name) //小强
	{
		fmt.Println("名字是：", name) //小强
		name = "Kr"               //重新赋值，给了 name。 name的值是kr
		fmt.Println("名字是：", name)
	}
	fmt.Println(">>>名字是：", name) //Kr?? 小强？？ ==》 Kr name的地址没变
}

func sampleSubdomain2() {
	name := "小强"              // 声明变量 name, 值是“小强”
	fmt.Println("名字是：", name) //小强
	{
		name = "小强-Update"
		fmt.Println("名字是：", name) //小强
		name := "Kr"              //重新赋值，给了 name。 name的值是kr
		fmt.Println("名字是：", name) //Kr
	}
	fmt.Println(">>>名字是：", name) //Kr?? 小强？？ ==》 小强-Update

	if name == "小强" {
		a := 3
		fmt.Println(a)
	} else {
		a := 4
		fmt.Println(a)
	}
}

func sampleSubdomainIf() {
	if v := calcAdd(); v == 0 {
		fmt.Println(v)
	} else {
		fmt.Println(v + 1)
	}
	// fmt.Println(v) // 无效。v的有效范围为if block
}

func sampleSubdomainFor() {
	for i := 0; i < 10; i++ {
		fmt.Println("hello golang, ", i)
	}
	// fmt.Println(i) // i 的有效范围为 for block
}

func fib(n uint) uint {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}
