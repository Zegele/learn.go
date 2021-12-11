package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b int8 = 30, 11
	a = math.MaxInt8 // 2^(8-1)-1=128-1=127
	//1+2+4+8+16=31

	fmt.Println(a)
	fmt.Printf("%T,%T\n", a, b)
	fmt.Println("a + b = ", a+b)
	fmt.Println("a - b = ", a-b)
	fmt.Println("a * b = ", a*b) //溢出错误
	fmt.Println("a / b = ", a/b)
	fmt.Println("a % b = ", a%b) //取余两个数只能是整型

	var c, d int = 30, 11
	fmt.Printf("%T,%T\n", c, d)
	fmt.Println("c + d = ", c+d)
	fmt.Println("c - d = ", c-d)
	fmt.Println("c * d = ", c*d)
	fmt.Println("c / d = ", c/d)
	fmt.Println("c % d = ", c%d)

	fmt.Println(true && false == false)
	fmt.Println("a>b=", a > b)
	fmt.Println("a<b=", a < b)

}
