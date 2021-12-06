package main

import "fmt"

func main() {
	v := []byte("hello")
	fmt.Println(v)
	v2 := string(v)
	fmt.Println(v2)
}
