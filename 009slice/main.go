package main

import "fmt"

func main() {
	a := []int{} //slice 可添加元素
	fmt.Println(a)
	fmt.Println("追加元素到a中，a是切片")
	a = append(a, 333)
	a = append(a, 444)

	fmt.Println(a)

	b := [0]int{} //arr 没用的东西
	fmt.Println(b)
	//b = append(b, 333)// 编译错误，
}
