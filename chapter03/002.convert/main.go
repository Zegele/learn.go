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
			debug.PrintStack() //这什么意思？ debug.PrintStack 表示输出调用栈
		}
	}()
	var a interface{} //any

	a = "string abc"

	b := a.(int) //这什么意思？ 将a这个黑盒子（空interface）中的内容，转成int类型的。 //这种方式能用？
	fmt.Println(b)
}
