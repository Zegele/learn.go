package main

import "fmt"

func main() {
	var a string = "您好"
	fmt.Println(a)

	fmt.Println("byte a:", []byte(a), len([]byte(a))) //byte a: [230 130 168 229 165 189] 6
	fmt.Println("rune a:", []rune(a), len([]rune(a))) //rune a: [24744 22909] 2

	abytes := []byte(a)
	arunes := []rune(a) // []int 不能转，[]byte可以和string互通

	fmt.Println(abytes)
	fmt.Println(arunes)
	fmt.Println("修改切片内的内容")

	abytes[0] = 'H'
	arunes[0] = 'H'
	a = string(arunes)
	fmt.Println(a) //H好
	a = string(abytes)
	fmt.Println(a) //H��好
}
