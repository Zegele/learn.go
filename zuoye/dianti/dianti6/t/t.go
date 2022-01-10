package main

import "fmt"

type Person struct {
	name string
	cl   *Class
}

type Class struct {
	nageClsee string
}

func t() *Person {
	person := Person{name: "a", cl: &Class{nageClsee: "math"}}
	fmt.Println(person)
	fmt.Println(person.cl)
	return &Person{cl: &Class{nageClsee: "china"}} //要盯住，需要实例化！
}

func main() {
	p := t()
	fmt.Println("-->", p)
	fmt.Println("-->", *p.cl)
}
