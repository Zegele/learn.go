package main

import "fmt"

func (c Calculator) Add() int {
	fmt.Printf("c的地址是=%p\n", &c) //在调用这个方法的时候，其实是复制，所以这个c的地址是变化的。
	tempResult := c.left + c.right
	c.result = tempResult //由于c的地址变化了，同时又没有使用指针，这样return的结果不会影响的原本的c。
	fmt.Println("调用Add函数，c的result=", c.result)
	return tempResult
}

func (c Calculator) Sub() int {

	return 0
}

func (c Calculator) Multiple() int {
	return 0
}

func (c Calculator) Divide() int {
	return 0
}

func (c Calculator) Reminder() int {
	return c.left % c.right
}

func (c *Calculator) SetResult(result int) {
	c.result = result
}
