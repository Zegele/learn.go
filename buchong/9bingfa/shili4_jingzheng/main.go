package main

import (
	"fmt"
	"runtime"
	"sync"
)

// c.biancheng.net/view/4358.html

// 互斥锁
var (
	counter int64
	wg      sync.WaitGroup
	mutex   sync.Mutex
)

func main() {
	for i := 0; i < 10; i++ {

		wg.Add(2)

		go incCounter(1)
		go incCounter(2)

		wg.Wait()
		fmt.Println(counter)

		fmt.Println("-------------")
	}
}

func incCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		// 同一时刻只允许一个goroutine进入这个临界区
		mutex.Lock()
		{
			value := counter
			runtime.Gosched() // 在这种情况下（互斥锁），当调用runtime.Gosched函数强制将当前goroutine退出当前线程后，
			// 调度器会再次分配这个goroutine继续运行。不会执行另一个goroutine。
			value++
			counter = value
		}
		mutex.Unlock() // 释放锁，允许其他正在等待的goroutine进入临界区
	}
	fmt.Println(id)
}

// 同一时刻只有一个goroutine可以进入临界区。之后直到调用Unlock函数之后，其他goroutine才能进去临界区。

/*
// 原子函数 atomic.StoreInt64(), atomic.StoreInt64()

var (
	shutdown int64
	wg       sync.WaitGroup
)

func main() {
	wg.Add(2)
	go doWork("A")
	go doWork("B")

	time.Sleep(1 * time.Second)
	fmt.Println("Shutdown Now")
	atomic.StoreInt64(&shutdown, 1)
	// main函数使用StoreInt64安全地修改shutdown变量的值。
	// 如果哪个dowork goroutine 试图在main函数调用StoreInt64的同时调用LoadInt64函数，那么原子函数会将这些调用互相同步。
	// 保证这些操作都是安全的，不会进入竞争状态。
	wg.Wait()
}

func doWork(name string) {
	defer wg.Done()

	for {
		fmt.Printf("Doing %s Work\n", name)
		time.Sleep(250 * time.Millisecond)

		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("Shutting %s Down\n", name)
			break
		}
	}

}
*/

/*
// 原子函数 atomic.AddInt64(&counter, 1)
var (
	counter int64
	wg      sync.WaitGroup
)

func main() {
	wg.Add(2)
	go incCounter(1)
	go incCounter(2)

	wg.Wait() // 等待goroutine结束
	fmt.Println(counter)
}

func incCounter(id int) { //传入这个参数没有使用
	defer wg.Done()
	for count := 0; count < 2; count++ {
		atomic.AddInt64(&counter, 1) //安全的对counter加1
		// atomic包的AddInt64函数，会同步整型值的加法，方法是强制同一时刻只能有一个goroutine运行并完成这个加法操作。
		// 当goroutine试图去调用任何原子函数时，这些goroutine都会自动根据所引用的变量做同步处理。
		runtime.Gosched()
	}
}
*/

/* 竞争资源导致错误
import (
	"fmt"
	"runtime"
	"sync"
)

var (
	count int32
	wg    sync.WaitGroup
)

func main() {

	wg.Add(2)
	go incCount()
	go incCount()
	wg.Wait()
	fmt.Println(count) //没有保护同时读写，导致结果众多。
	count = 0

}

func incCount() {
	defer wg.Done()
	for i := 0; i < 2; i++ {
		value := count
		runtime.Gosched() // 是让当前goroutine暂停，退回执行队列，让其他等待的goroutine运行。
		// 这里的目的是让资源竞争的结果更明显
		value++
		count = value
	}
}
*/
