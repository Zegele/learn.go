package main

import (
	"fmt"
	"sync"
)

// 在读多写少的环境中，可以优先使用读写互斥锁（sync.RWMutex），它比互斥锁更加高效。
// sync包中的RWMutex提供了读写互斥锁的封装。

var (
	// 逻辑中使用的某个变量
	count int

	// 与变量对应的使用读写互斥锁
	countGuard sync.RWMutex
)

func GetCount() int {
	// 锁定
	countGuard.RLock()
	// 获取count的过程是一个读取count数据的过程，适用于读写互斥锁。
	//RLock将互斥锁标记为读状态。如果此时另外一个goroutine并发访问了countGuard,
	//同时也调用了countGuard.RLock()时，并不会发送阻塞。

	// 在函数退出时解除锁定
	defer countGuard.RUnlock() //读模式解锁
	return count
}

func SetCount(i int) {
	countGuard.Lock()
	count = i
	countGuard.Unlock()

}
func main() {
	SetCount(100)
	fmt.Println(GetCount())
}
