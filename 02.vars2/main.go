package main

import (
	"fmt"
	"strconv"
)

func main() {
	name := "小强"
	{
		val, err1 := strconv.Atoi(name) //这样不能把string转变成int
		fmt.Println(val, err1)

		var age1 string = "30"
		age2, err := strconv.Atoi(age1) //字符串转数字(int类型)
		fmt.Println(age2, err)
		fmt.Printf("%T\n", age2)
		age3 := strconv.Itoa(age2)
		fmt.Printf("%T,%v\n", age3, age3)

	}
	age := 30
	fmt.Printf("%p\n", &age)
	age, time := 32, "时间"    //有新变量（time），则可以用 := 符号
	fmt.Printf("%p\n", &age) //%p打印指针地址
	fmt.Println(name, age, time)
	{
		age := 3
		fmt.Printf("%p\n", &age) //不同的作用域，即使相同的变量名，实质是不一样的东西。
	}

}
