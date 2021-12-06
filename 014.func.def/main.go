package main

import "fmt"

func main() {
	hello()
	helloToSomeone("aa")
	constructHelloMessage("bb")
	v := constructHelloMessage("bb")
	fmt.Println(v)
}

func hello() {
	fmt.Println("你好，golang！")
}

func helloToSomeone(name string) {
	fmt.Println("你好，", name)
}

func constructHelloMessage(name string) string {
	fmt.Println("你好，", name)
	return "你好，" + name + name
}
