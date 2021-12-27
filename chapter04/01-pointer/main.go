package main

import "fmt"

//指针 &a, *a *int 区别这3个东西
func main() {
	a, b := 1, 2
	add(&a, &b) //&a 是a盒子， 不是a内部的值
	fmt.Println(a)
}

/*
	func add(a, b int){
		a = a+b
	}
*/

func add(a, b *int) { //指针类型
	*a = *a + *b //*a a盒子内的东西
}
