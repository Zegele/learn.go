package main

import (
	"fmt"
	"sync"
)

/*
func main() {
	var mu sync.Mutex

	go func() {
		fmt.Println("Golang!")
		mu.Lock()
	}()
	mu.Unlock()
}



// 由于mu.Lunk() 和 mu.Unlock()并不在同一个goroutine中
// 所以也就不满足顺序一致性内存模型
//同时他们也没有其他的同步事件可以参考，也就是说这两件事是可以并发的。
// 因为可能是并发的事件，所以main()函数中的muUnlock()很有可能先发生，而这个时刻mu互斥对象还处于未加锁的状态，因而会导致运行时异常。

*/

/*
// 修复后，如下：
func main() {
	var mu sync.Mutex
	mu.Lock()
	go func() {
		fmt.Println("Golang！")
		mu.Unlock()
	}()
	mu.Lock()
}

// 修复的地方实在main() 函数所在线程中执行两次mu.Lock() 当第二次加锁时会因为锁已经被占用（不是递归锁）而阻塞，
//main()函数的阻塞状态驱动后台线程继续向前执行。
// 当后台线程执行到mu.Unlock()时解锁，此时打印工作已经完成，解锁会导致main函数中的第二个mu.Lock()阻塞状态取消
// 此时后台线程和主线程在没有其他的同步事件参考，他们退出的事件将是并发的
// 在main函数退出导致程序退出时，后台线程可能已经退出了，也可能没有退出。
// 虽然无法确定两个线程退出的时间，但是打印工作是可以正确完成的。


*/

/*
// 使用sync.Mutex 互斥锁同步是比较低级的做法，我们现在改用无缓存通道来实现同步：
func main() {
	done := make(chan int)// 无缓冲，先准备好接收，然后传入数据

	go func() {
		fmt.Println("Golang")
		<-done
	}()
	done <- 1
}

// 根据go语言内存模型规范，对于无缓存通道进行的接收，发生在对该通道进行的发送完成之前。
// 因此，后台线程 <-done 接收操作完成之后，main线程的 done<-1 发送操作才可能完成（从而退出main，退出程序）
// 而此时打印工作已经完成
*/

/*
//上面的代码虽然可以正确同步，但是对通道的缓存大小太敏感，如果通道有缓存，就无法保证main函数退出之前后台线程能正常打印了。
//更好的做法是将通道的发送和接收方向调换一下，这样可以避免同步事件受通道缓存大小的影响。

func main() {
	done := make(chan int, 1) // 带缓存通道, 先传入，后接收 //要理解这两个示例!!!

	go func() {
		fmt.Println("golang!")
		done <- 1 //和上面的代码调换了位置
	}()
	<-done
}
// 通道是带缓存的，main线程接收完成，是在后台线程发送开始但还未完成的时刻，此时打印工作也是已经完成的。
// 对于带缓存的通道，对通道的第K个接收完成操作发生在第K+C个发送操作完成之前，其中C是通道的缓存大小。
*/

/*
// 基于带缓存通道，我们可以很容易将打印线程扩展到N个，下面的示例是开启10个后台线程分别打印：

func main() {
	done := make(chan int, 10) // 带10个缓存

	// 开n个后台打印线程
	for i := 0; i < cap(done); i++ {
		go func() {
			//fmt.Println("Golang", i)//不要让for循环的参数和你的goroutine有影响
			// 因为在你打印的时候，for循环可能都已经循环完很久了。
			fmt.Println("Golang")
			done <- 1
		}()
	}
	// 等待n个后台线程完成
	for i := 0; i < cap(done); i++ {
		<-done
	}
}

*/

// 对于这种要等待N个线程完成后再进行下一步的同步操作有一个简单的做法，就是使用sync.WaitGroup来等待一组事件。
func main() {
	var wg sync.WaitGroup
	//wg.Add(10) 为什么不放在for循环外面？
	// 因为实际使用中你可能不知道有多少个goroutine，所以放到循环里，有一个搞一个，减一个
	// 而且我分析，下面那种来一个搞一个，减一个，wait完一个，速度更快，因为不用等着wait10个都完成。
	// 开n个后台打印线程
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			fmt.Println("Golang")
			wg.Done()
		}()
	}
	// 等待N个后台线程完成
	wg.Wait()
}
