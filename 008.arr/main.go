package main

import "fmt"

func main() {
	//xqnames := [3]string{"小强", "男", "在职"}
	//xlnames := [3]string{"小李", "男", "在职"}
	//xhnames := [3]string{"小红", "男", "在职"}
	all := [3][3]string{
		{"小强", "男", "在职"},
		{"小李", "男", "在职"},
		{"小红", "男", "在职"},
	}
	for d1, d1val := range all {
		fmt.Println(d1, d1val)
		for d2, d2val := range d1val {
			fmt.Println(d2, d2val)
		}
	}
}
