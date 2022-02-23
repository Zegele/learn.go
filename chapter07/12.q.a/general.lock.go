package main

import "sync"

var counter *safeCount = &safeCount{}

type safeCount struct {
	totalNumber      int
	totalLetterCount int
	totalWordCount   int

	// ...
	sync.Mutex //锁是结构体的一部分 一把锁管该结构体的3个变量
}

//线程安全的加数据
func (c *safeCount) AddNumber(totalNumber int, totallLetterCount int, totalWordCount int) { //参数是3个变量
	c.Lock()
	defer c.Unlock()
	c.totalNumber += totalNumber
	//...
}

func countNumber() {
	counter.AddNumber(100, 5, 500)
}
