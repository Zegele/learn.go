package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

/*
// 死锁 进程彼此等待

type value struct {
	memAccess sync.Mutex
	value     int
}

func main() {
	runtime.GOMAXPROCS(4)
	var wg sync.WaitGroup

	sum := func(v1, v2 *value) {
		defer wg.Done()
		v1.memAccess.Lock()
		time.Sleep(2 * time.Second)
		v2.memAccess.Lock()
		fmt.Printf("sum = %d\n", v1.value+v2.value)
		v1.memAccess.Unlock()
		v2.memAccess.Unlock()
	}

	product := func(v1, v2 *value) {
		defer wg.Done()
		v2.memAccess.Lock() //这里的v2已经上锁了，上面的sum函数，v2一直在等上锁
		time.Sleep(2 * time.Second)
		v1.memAccess.Lock() // 同理上面sum函数的v1已经上锁了，这里的v1还在等上锁。所以死锁了。
		fmt.Printf("product = %d\n", v1.value*v2.value)
		v1.memAccess.Unlock()
		v2.memAccess.Unlock()
	}

	var v1, v2 value
	v1.value = 1
	v2.value = 2
	wg.Add(2)
	go sum(&v1, &v2)
	go product(&v1, &v2)
	wg.Wait()
}

*/

/*
// 活锁
// 活跃性问题，尽管不会阻塞线程，但也不能继续执行，因为线程不断重复同样的操作，而且总会失败。
// 例如：线程1可以使用资源，但它礼貌让其他线程先使用资源，线程2也可以使用资源，但它同样很绅士，让其他线程先使用资源，这样你让我，我让你，最后没有使用资源。
// 活锁通常发生在处理事务消息中，如果不能成功处理某一个消息，那么消息处理机制将回滚事务，并将它重新放在队列的开头。
// 这样错误的事务被一直回滚重复执行，这种形式的活锁是由于过度的错误恢复代码造成的，因为它的错误将不可修复的错误认为可以修复的错误。

// 解决这种活锁的问题，需要在重试机制中引入随机性，例如在网络上发送数据包，如果检测冲突，都要停止并在一段时间后重发，如果都在1秒后重发，还是会冲突。
// 所以引入随机性可以解决该类问题
func main() {
	runtime.GOMAXPROCS(4)
	cv := sync.NewCond(&sync.Mutex{})
	go func() {
		for range time.Tick(1 * time.Second) { // for range 还能这样用？？？
			// 通过tick控制两个人的步调
			cv.Broadcast()
		}
	}()

	takeStep := func() {
		cv.L.Lock()
		cv.Wait() // 这个wait是等待broadcast唤醒？
		cv.L.Unlock()
	}

	tryDir := func(dirName string, dir *int32, out *bytes.Buffer) bool {
		fmt.Fprintf(out, "%+v", dirName)
		atomic.AddInt32(dir, 1)
		takeStep()
		fmt.Println("asdgasdf", *dir)   // 关键在这，假设男人的dir（方向）从0加到1，但是女人的dir（方向）和男人的一样，同时又加了1，变成了2。所有走不成功。
		if atomic.LoadInt32(dir) == 1 { // 走成功就返回
			fmt.Println(out, ".Success!")
			return true
		}
		takeStep() // 没走成功， 再走回来
		atomic.AddInt32(dir, -1)
		fmt.Println("dgadg", *dir)
		return false
	}

	var left, right int32
	tryLeft := func(out *bytes.Buffer) bool {
		return tryDir("向左走 ", &left, out)
	}
	tryRight := func(out *bytes.Buffer) bool {
		return tryDir("向右走 ", &right, out)
	}

	walk := func(walking *sync.WaitGroup, name string) {
		var out bytes.Buffer
		defer walking.Done()
		defer func() {
			fmt.Println(out.String())
		}()
		fmt.Fprintf(&out, "%v is tring to scoot:", name)
		for i := 0; i < 5; i++ {
			if tryRight(&out) || tryLeft(&out) {
				return
			}
		}
		fmt.Fprintf(&out, "\n%v is tried !", name)
	}

	var trail sync.WaitGroup
	trail.Add(2)
	go walk(&trail, "男人")
	go walk(&trail, "女人")
	trail.Wait()
}

*/

// 饥饿 ：
//饥饿指一个可运行的进程尽管能继续执行，但被调度器无限期地忽视，而不能被调度执行的情况

// 与死锁不同，饥饿是锁在一段时间内，优先级低的线程最终还是会执行的，比如高优先级的线程执行完之后释放了资源。
// 活锁与饥饿是无关的，因为在活锁中，所有并发进程都是相同的，并且没有完成工作。
// 更广泛地说，饥饿通常意味着有一个活多个贪婪的并发进程，他们不公平地阻止一个或多个并发进程，以尽可能有效地完成工作，或者阻止全部并发进程。
func main() {
	runtime.GOMAXPROCS(4)
	var wg sync.WaitGroup
	const runtime = 1 * time.Second
	var shareLock sync.Mutex
	greedyWork := func() {
		defer wg.Done()
		var count int
		for begin := time.Now(); time.Since(begin) <= runtime; {
			shareLock.Lock()
			time.Sleep(3 * time.Nanosecond)
			shareLock.Unlock()
			count++
		}
		fmt.Printf("Greedy worker was able to execute %v work loops\n", count)
	}

	politeWork := func() {
		defer wg.Done()
		var count int
		for begin := time.Now(); time.Since(begin) <= runtime; {
			shareLock.Lock()
			time.Sleep(1 * time.Nanosecond)
			shareLock.Unlock()

			shareLock.Lock()
			time.Sleep(1 * time.Nanosecond)
			shareLock.Unlock()

			shareLock.Lock()
			time.Sleep(1 * time.Nanosecond)
			shareLock.Unlock()
			count++
		}
		fmt.Printf("Polite work was able to execute %v work loops\n", count)
	}
	wg.Add(2)
	go greedyWork()
	go politeWork()
	wg.Wait()
	//贪婪几乎是平和的3倍工作量，实际是没有平衡好
}

// 总结 死锁， 活锁， 饥饿
// 不使用锁肯定会出问题。如果用了，虽然解决了前面的问题，但是有出现了更多的新问题。
// 死锁：是以内错误的使用了锁，导致异常；
// 活锁：一直在僵尸跑。活死人（程序）
// 饥饿：与锁使用的粒度有关，通过计数取样，可以判断进程的工作效率
// 只要有共享资源的访问，必定要使其逻辑上进行顺序化和原子化，确保访问一致，这绕不开锁这个概念。
