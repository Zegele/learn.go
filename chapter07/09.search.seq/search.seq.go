package main

import "fmt"

var A = []int64{
	1,
}

func main() {
	arr := A
	fmt.Print(arr) //undefined: A 看来main包，最好只有一个文件。
}
