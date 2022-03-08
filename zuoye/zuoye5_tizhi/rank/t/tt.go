package main

//关于指针

import "fmt"

func main() {
	//关于指针
	var p *int //不能直接使用
	var q *int //
	p = new(int)
	*p = 1
	q = p
	fmt.Println(p, &p, *p) //p内储存的值：0xc000018088 p自身所在的值：0xc000006028 p内储存的值对应的值1
	fmt.Println(q, &q, *q) //q内储存的值：0xc000018088 q自身所在的值：0xc000006030 q内储存的值对应的值1
}
