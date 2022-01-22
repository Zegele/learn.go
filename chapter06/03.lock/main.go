package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	for i := 0; i < 1; i++ {
		// countDict()
		// countDictGoroutineSafe()
		// countDictForgetUnlock()
		// countDictLockPrice()
		countDictGoroutineSafeRW()
	}
}

func countDict() {
	fmt.Println("开始数")
	var totalCount int64 = 0
	wg := sync.WaitGroup{} // wg := sync.WaitGroup{}/ wg.Add/ wg.Done
	wg.Add(5000)           //因为想让5000个人数，所以这里是5000，每个人完成后，这个Add中的数会减少1
	for p := 0; p < 5000; p++ {
		go func() {
			defer wg.Done()
			totalCount += 100
		}()
	}
	wg.Wait()
	fmt.Println("预计有", 100*5000, "字")
	fmt.Println("总共有", totalCount, "字")
}

func countDictGoroutineSafe() { //同步锁
	fmt.Println("开始数")
	var totalCount int64 = 0
	totalCountLock := sync.Mutex{} //同步锁，让goroutine排队拿数据，不会出现插队的错误。

	wg := sync.WaitGroup{}
	wg.Add(5000)
	for p := 0; p < 5000; p++ {
		go func() {
			defer wg.Done()
			totalCountLock.Lock()
			defer totalCountLock.Unlock()
			totalCount += 100
		}()
	}
	wg.Wait()
	fmt.Println("预计有：", 100*5000, "字")
	fmt.Println("总共有：", totalCount, "字")
}

func countDictGoroutineSafeRW() { //读写锁
	fmt.Println("开始数")
	var totalCount int64 = 0
	totalCountLock := sync.RWMutex{}

	wg := sync.WaitGroup{}
	workerCount := 5
	wg.Add(workerCount)

	doneCh := make(chan struct{})
	for p := 0; p < workerCount; p++ {
		go func(p int) { // 读锁可以多个go routine同时拿到。读可以同时读
			fmt.Println(p, "读锁开始时间：", time.Now())
			totalCountLock.RLock()
			fmt.Println(p, "读锁拿到锁时间：", time.Now())
			totalCountLock.RUnlock()
			time.Sleep(1 * time.Second)
		}(p)
	}

	for p := 0; p < workerCount; p++ {
		go func(p int) {
			defer wg.Done()

			fmt.Println(p, "写锁开始时间：", time.Now())
			totalCountLock.Lock() //写还是得排队
			fmt.Println(p, "写锁拿到锁时间：", time.Now())
			defer totalCountLock.Unlock()
			totalCount += 100
		}(p)
	}
	wg.Wait()
	close(doneCh)
	time.Sleep(1 * time.Second)
	fmt.Println("预计有：", 5000*100, "字")
	fmt.Println("总共有：", totalCount, "字")
}

func countDictLockPrice() {
	fmt.Println("开始数")
	var totalCount int64 = 0
	totalCountLock := sync.Mutex{} //注意：有大括号

	wg := sync.WaitGroup{}
	wg.Add(5)
	for p := 0; p < 5; p++ {
		go func(pageNum int) {
			defer wg.Done()
			totalCountLock.Lock()
			totalCount += 100
			//if pageNum == 3{
			//	time.Sleep(3*time.Second)
			//}
			totalCountLock.Unlock()
		}(p)
	}
	wg.Wait()
	fmt.Println("预计有：", 100*5000, "字")
	fmt.Println("总共有：", totalCount, "字")
}

func countDictForgetUnlock() {
	fmt.Println("开始数")
	var totalCount int64 = 0
	totalCountLock := sync.Mutex{}

	wg := sync.WaitGroup{}
	wg.Add(5)
	for p := 0; p < 5; p++ {
		go func() {
			defer wg.Done()
			totalCountLock.Lock()
			totalCount += 100
			//没有解锁
		}()
	}
	wg.Wait()
	fmt.Println("预计有：", 5000*100, "字")
	fmt.Println("总共有：", totalCount, "字")
}
