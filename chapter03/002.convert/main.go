package main

//没看懂

import (
	"fmt"
	"runtime/debug"
)

func main() {
	convertType()
	fmt.Println("finish")
}
func convertType() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic啦：", r)
			debug.PrintStack() //这什么意思？
		}
	}()
	var a interface{} //any
	a = "string aaa"

	b := a.(int) //这什么意思？
	fmt.Println(b)
}
