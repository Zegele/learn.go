package main

import "fmt"

func main() {
	const pi = 3.14159265359
	pi = 3.14 //这里会报错
	fmt.Println(pi)
}
