package main

import "fmt"

func main() {
	var fruit string = "6 apples"
	var watermelon bool = true // true or false
	if watermelon {
		fruit = "1 apple"
	} else {
		fruit = "7 apples"
	}
	fmt.Println("buy: ", fruit)
}
