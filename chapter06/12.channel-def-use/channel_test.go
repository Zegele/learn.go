package main

import (
	"fmt"
	"testing"
	"time"
)

func TestDefChannel(t *testing.T) {
	var sampleMap map[string]int = map[string]int{} //map可以直接这样实例化
	fmt.Println("sampleMap", sampleMap)
	var intCh chan int           //= make(chan int) // = chan int {} channel不能这样实例化 只能用make关键字实例化
	fmt.Println("intCh:", intCh) //intCh: <nil> 未初始化是个nil
	intCh = make(chan int, 1)    //intCh = make(chan int)       //默认长度是0的channel
	fmt.Println("intCh:", intCh) //intCh: 0xc0000461e0 已初始化，且无法通过*intCh获取值
	fmt.Println(len(intCh))

	fmt.Println("装入数据")
	intCh <- 3
	fmt.Println("取出数据")
	out := <-intCh
	fmt.Println("取出的数据是：", out)
}

func TestChanPutGet(t *testing.T) {
	intCh := make(chan int) // 创建一个不带size的channel（不带buffer）
	workerCount := 10
	for i := 0; i < workerCount; i++ {
		go func(i int) {
			intCh <- i
			fmt.Printf("注入%d\n", i)
		}(i)
	}

	for o := 0; o < workerCount; o++ {
		go func(o int) {
			out := <-intCh
			fmt.Printf("出%d拿到%d\n", o, out)
		}(o)
	}
	time.Sleep(1 * time.Second)
}

// 这是一个让out部分等待一段时间再取数据，来观察i的行为
// 结果：如果没有out，那么in部分会等待，直到有out开始。这样效率不高。
func TestChanPutGet2_Owait(t *testing.T) { //让Out等待
	intCh := make(chan int)
	workerCount := 10
	for i := 0; i < workerCount; i++ {
		go func(i int) {
			fmt.Println(i, "开始工作", time.Now())
			intCh <- i //数据装进去后，如果没有被取出，就会一直等待。等待有人拿取数据。
		}(i)
	}
	fmt.Println(time.Now())
	time.Sleep(2 * time.Second)
	fmt.Println(time.Now())

	for o := 0; o < workerCount; o++ {
		go func(o int) {
			out := <-intCh
			fmt.Printf("%s, 出%d拿到%d\n", time.Now(), o, out)
		}(o)
	}
	time.Sleep(1 * time.Second)
}

// 这是一个让out部分等待一段时间再取数据，来观察i的行为
// 结果：带有buffer的会直接允许in，不影响out。
func TestChanPutGet2_Owait_withBuffer(t *testing.T) {
	intCh := make(chan int, 10) //创建一个带size的channel（带buffer）
	workerCount := 10
	for i := 0; i < workerCount; i++ {
		go func(i int) {
			fmt.Println(i, "开始工作", time.Now())
			intCh <- i //数据都装在buffer（缓冲）中。最多装10个数
			fmt.Println(i, "结束工作", time.Now())
		}(i)
	}
	fmt.Println(time.Now())
	time.Sleep(2 * time.Second)
	fmt.Println(time.Now())

	for o := 0; o < workerCount; o++ {
		go func(o int) {
			out := <-intCh
			fmt.Printf("%s, 出%d拿到%d\n", time.Now(), o, out)
		}(o)
	}
	time.Sleep(1 * time.Second)
}

// 这是一个让out部分等待一段时间再取数据，来观察i的行为
// 结果：带有buffer的会直接允许in，不影响out。
func TestChanPutGet2_Owait_withSmallBuffer(t *testing.T) {
	intCh := make(chan int, 2) //创建一个带size的channel（带buffer）,最多只能装2个数
	workerCount := 10
	for i := 0; i < workerCount; i++ {
		go func(i int) {
			fmt.Println(i, "开始工作", time.Now())
			intCh <- i //数据都装在buffer（缓冲）中。最多装2个数
			fmt.Println(i, "结束工作", time.Now())
		}(i)
	}
	fmt.Println(time.Now())
	time.Sleep(2 * time.Second)
	fmt.Println(time.Now())

	for o := 0; o < workerCount; o++ {
		go func(o int) {
			out := <-intCh
			fmt.Printf("%s, 出%d拿到%d\n", time.Now(), o, out)
		}(o)
	}
	time.Sleep(1 * time.Second)
}

// out 先启动， in 后续启动
// 结果：如果先out， 因为没有数据则一直在等待传入数据。
func TestChanPutGet2_OFirst_withBuffer(t *testing.T) {
	intCh := make(chan int, 10) //创建一个带size的channel（带buffer）,最多能装10个数
	workerCount := 10
	for o := 0; o < 10; o++ {
		go func(o int) {
			out := <-intCh
			fmt.Printf("%s, 出%d拿到%d\n", time.Now(), o, out)
		}(o)
	}

	fmt.Println(time.Now())
	time.Sleep(2 * time.Second)
	fmt.Println(time.Now())

	for i := 0; i < workerCount; i++ {
		go func(i int) {
			fmt.Println(i, "开始工作", time.Now())
			intCh <- i //数据都装在buffer（缓冲）中。最多装2个数
			fmt.Println(i, "结束工作", time.Now())
		}(i)
	}

	time.Sleep(1 * time.Second)
}

//遍历channel
// range 没有 close 的话，在去除所有数据后，panic deadlock
func TestRangeChannel(t *testing.T) { //遍历channel
	intCh := make(chan int, 10) //创建一个带size的channel（带buffer）
	intCh <- 1
	intCh <- 2
	intCh <- 3
	intCh <- 4
	intCh <- 5

	for o := range intCh {
		fmt.Println(o)
	}
}

// range 没有 close 的话，在去除所有数据后，panic deadlock
func TestRangeClosedChannel(t *testing.T) { //遍历channel
	intCh := make(chan int, 10) //创建一个带size的channel（带buffer）
	intCh <- 1
	intCh <- 2
	intCh <- 3
	intCh <- 4
	intCh <- 5

	close(intCh)

	{
		o1, ok := <-intCh
		fmt.Println("直接取数：", o1, ok)
	}

	for o := range intCh {
		fmt.Println("range 取数：", o)
	}

	o1 := <-intCh   //o1是假值
	fmt.Println(o1) //如果没有ok，输出的0值不知道是真0值，还是假0值。
	{
		o1, ok := <-intCh
		fmt.Println("遍历后取数：", o1, ok)
	}
}

// select 没有default
func TestSelectChannel(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- 1
	}()

	go func() {
		//time.Sleep(1 * time.Second)
		ch2 <- "hh"
	}()

	select {
	case <-ch1: //case o := <-ch1: //也可以定义变量o接住取出的数据。
		fmt.Println("ch1 ready, go")
	case <-ch2:
		fmt.Println("ch2 ready, go")
	}
}

// select 有default 并且channel没有准备好
func TestSelectDefaultChannel(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan string)

	go func() {
		//time.Sleep(1 * time.Second)
		ch1 <- 1
	}()

	go func() {
		//time.Sleep(1 * time.Second)
		ch2 <- "hh"
	}()

	select {
	case o := <-ch1: //case o := <-ch1: //也可以定义变量o接住取出的数据。
		fmt.Println(time.Now(), "ch1 ready, go", o)
	case o := <-ch2: //o的定义域是这个case
		fmt.Println(time.Now(), "ch2 ready, go", o)
	default: //顺序执行的时候，channel还没有赋值，程序已经跑到default了，所以直接打印了default
		fmt.Println(time.Now(), "所有的channel都不ready，直接走default")
	}
}

// select 有default，并且channel准备好了
// case 的优先级高于default， 只要有case中的channel ready， default不会运行。
func TestSelectDefaultChannelAndReady(t *testing.T) {
	ch1 := make(chan int, 1)
	ch2 := make(chan string)
	fmt.Println("start:", time.Now())
	ch1 <- 1
	/*
		go func() {
			//time.Sleep(1 * time.Second)
			ch1 <- 1
		}()

	*/

	go func() {
		//time.Sleep(1 * time.Second)
		ch2 <- "hh"
	}()

	select {
	case o := <-ch1: //case o := <-ch1: //也可以定义变量o接住取出的数据。
		fmt.Println(time.Now(), "ch1 ready, go", o)
	case o := <-ch2: //o的定义域是这个case
		fmt.Println(time.Now(), "ch2 ready, go", o)
	default:
		fmt.Println(time.Now(), "所有的channel都不ready，直接走default")
	}
	fmt.Println("DONE")
}

// 只要channel是关闭的，select时，总是ready好的。总是可用的。
func TestSelectChannelWithDefaultAndClosedChannel(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan string)
	fmt.Println("start:", time.Now())

	//go func() {
	//	//time.Sleep(1 * time.Second)
	//	ch1 <- 1
	//}()
	fmt.Println("关闭 ch1")
	close(ch1)

	go func() {
		//time.Sleep(1 * time.Second)
		ch2 <- "hh"
	}()

	fmt.Println("select:", time.Now())
	select {
	case o := <-ch1: //case o := <-ch1: //也可以定义变量o接住取出的数据。
		fmt.Println(time.Now(), "ch1 ready, go", o)
	case o := <-ch2: //o的定义域是这个case
		fmt.Println(time.Now(), "ch2 ready, go", o)
	default: //顺序执行的时候，channel还没有赋值，程序已经跑到default了，所以直接打印了default
		fmt.Println(time.Now(), "所有的channel都不ready，直接走default")
	}
}

func TestMultipleSelect(t *testing.T) {
	ch1 := make(chan int)
	for i := 0; i < 10; i++ {
		go func(i int) {
			select {
			case <-ch1: //ch1中没有值，一直在等待。但是当ch1 close后，则总是被执行。就可以直接运行了。
				fmt.Println(time.Now(), i)
			}
		}(i)
	}

	fmt.Println("关闭channel", time.Now())
	close(ch1)

	time.Sleep(1 * time.Second)
}

func TestDualCloseChannel(t *testing.T) {
	c := make(chan struct{})
	close(c)
	//close(c) // panic: close of closed channel
}

func TestPut2ClosedChannel(t *testing.T) { //遍历channel
	intCh := make(chan int, 10) //创建一个带size的channel（带buffer）
	intCh <- 1
	intCh <- 2
	intCh <- 3
	intCh <- 4
	intCh <- 5

	close(intCh)
	intCh <- 5 //panic: send on closed channel [recovered]
	//channel 关闭了，可取数，但不能存入数据。
}

func TestOutOnly(t *testing.T) {
	intCh := make(chan int, 10)
	<-intCh //fatal error: all goroutines are asleep - deadlock!
}

func TestSingelGoroutineApp(t *testing.T) {
	intCh := make(chan int) //单线程谨慎使用channel
	intCh <- 1
	<-intCh //fatal error: all goroutines are asleep - deadlock!
}

func TestMultipleChannelReadySelect(t *testing.T) {
	ch1, ch2 := make(chan int), make(chan int)
	close(ch1)
	close(ch2)

	ch1Counter, ch2Counter := 0, 0
	for i := 0; i < 10000; i++ {
		select {
		case <-ch1:
			ch1Counter++
		case <-ch2:
			ch2Counter++
		}
	}
	fmt.Println("ch1Counter:", ch1Counter, "ch2Counter:", ch2Counter) //随机的
}
