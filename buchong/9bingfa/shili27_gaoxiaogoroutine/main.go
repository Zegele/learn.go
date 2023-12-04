// goroutine（Go 语言开发）如何使用才更加高效
// 如何高效使用goroutine
// 源码参考：github.com/TianhaoLi/GoLearn/blob/master/src/go_dev/part12/2_exitnotify.go
// 文档参考：www.h5w3.com/linux/536960.html
package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

// 优化后
// 套接字接收过程
func socketRecv(conn net.Conn, wg *sync.WaitGroup) {
	// 创建一个接收的缓冲
	buff := make([]byte, 1024)
	// 不停地接收数据
	for {
		// 从套接字中读取数据
		_, err := conn.Read(buff)
		// 需要结束接收，退出循环
		if err != nil {
			break
		}
		// 函数已经结束，发送通知
		wg.Done()
	}
}

// 优化 通道内部也是锁实现的。用等待锁来优化锁操作的消耗。
func main() {
	// 连接一个地址
	conn, err := net.Dial("tcp", "www.163.com:80")
	// 发生错误时打印错误退出
	if err != nil {
		fmt.Println(err)
		return
	}
	// 退出通道
	var wg sync.WaitGroup
	// 添加一个任务
	wg.Add(1)
	// 并发执行接收套接字
	go socketRecv(conn, &wg)
	// 在接收时，等待1秒
	time.Sleep(time.Second)
	//主动关闭套接字
	conn.Close()
	// 等待goroutine退出完毕
	wg.Wait()
	fmt.Println("recv done")
}

// 优化前
//避免在不必要的地方使用通道

// 套接字接收过程
//func socketRecv(conn net.Conn, exitChan chan string) {
//	创建一个接收缓冲
//	buff := make([]byte, 1024)
// 不停地接收数据
//	for {
//		//从套接字中读取数据
//		_, err := conn.Read(buff)
//  	// 需要结束接收，退出循环
//		if err != nil {
//			break
//		}
//	}
// 函数已经结束，发送通知。
//	exitChan <- "recv exit"
//}

//func main() {
//	// 连接一个地址
//	conn, err := net.Dial("tcp", "www.163.com:80")
// 发生错误时打印错误退出
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
// 创建退出通道
//	exit := make(chan string)
// 并发执行套接字接收
//	go socketRecv(conn, exit)
//在接收时，等待1秒
//	time.Sleep(time.Second)
//主动关闭套接字
//	conn.Close()
//等待goroutine退出完毕
//	fmt.Println(<-exit)
//}
// 例子中，goroutine退出使用通道来通知，这种做法可以解决问题，但是实际上通道中的数据并没有完全使用。
// 意思是通道的数据只是用来发信号了，并没有使用数据？

/*
// 改良
// 一段耗时的计算函数
func consumer(ch chan int) {
	// 无限获取数据的循环
	for {
		// 从通道获取数据
		data := <-ch
		fmt.Println("----------?????", data)
		if data == 0 {
			break
		}
		// 打印数据
		fmt.Println("----------", data)
	}
	fmt.Println("goroutine exit")
}

func main() {
	// 传递数据用的通道
	ch := make(chan int)
	for {
		// 空变量， 什么也不做
		var dummy string
		// 获取输入，模拟进程持续运行
		fmt.Scan(&dummy)
		if dummy == "quit" {
			for i := 0; i < runtime.NumGoroutine()-1; i++ {
				ch <- 0 //让运行的goroutine都去结束
			}
			continue
		}
		// 启动并发执行consumer()函数
		go consumer(ch)
		// 输出现在的goroutine数量
		fmt.Println("goroutines:", runtime.NumGoroutine())
	}
}

*/

/*
//一段耗时的计算函数
func consumer(ch chan int) {
	// 无限获取数据的循环
	for {
		// 从通道获取数据
		data := <-ch
		// 打印数据
		fmt.Println(data)
	}
}
func main() {
	// 创建一个传递数据用的通道
	ch := make(chan int)
	for { // goroutine越来越多，因为consumer没有退出。
		// 空变量，什么也不做
		var dummy string
		// 获取输入，模拟进程持续运行
		fmt.Scan(&dummy)
		// 启动并发执行consumer()函数
		go consumer(ch)
		// 输出现在的goroutine数量
		fmt.Println("goroutine:", runtime.NumGoroutine())
	}
}


*/
