package main

import (
	"fmt"
	"sync"
)

var (
	// 逻辑中使用的某个变量
	count int // 无论是包级的变量还是结构体成员字段，都可以

	// 与变量对应的使用互斥锁
	countGuard sync.Mutex
	//一般，将互斥锁的颗粒度设置得越小越好，降低因为共享访问时等待的时间。
)

func GetCount() int { // 是一个获取count值的函数封装，通过这个函数可以并发安全的访问变量count
	// 锁定
	countGuard.Lock()
	// 尝试对countGuard互斥量进行加锁。一旦countGuard发生加锁，如果另外一个goroutine尝试继续加锁时，
	//将会发送阻塞，直到这个countGuard被解锁

	// 在函数退出时解除锁定
	defer countGuard.Unlock()

	return count
}

func SetCount(c int) { // 在设置count值时，同样使用countGuard进行加锁，解锁操作，保证修改count值的过程时一个原子过程，不会发送访问冲突。
	countGuard.Lock()
	count = c
	countGuard.Unlock()
}

func main() {
	// 可以进行并发安全的设置
	SetCount(1)

	// 可以进行并发安全的获取
	fmt.Println(GetCount())
}
