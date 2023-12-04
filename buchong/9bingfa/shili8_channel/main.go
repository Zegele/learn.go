package main

import (
	"fmt"
	"time"
)

// 循环接收
func main() {
	// 构建一个通道
	ch := make(chan int)

	// 开启一个并发匿名函数
	go func() {

		// 从3循环到0
		for i := 3; i >= 0; i-- {
			// 发送3到0之间的数值
			ch <- i

			// 每次发送完时等待
			time.Sleep(time.Second)
		}
	}()

	// 遍历接收通道数据
	// 循环遍历要记着退出。
	for data := range ch {
		// 打印通道数据
		fmt.Println(data)

		// 当遇到数据0时，退出循环
		// 如果没有这个break退出，会死锁
		if data == 0 {
			break
		}
	}
}

/*
func main() {
	ch := make(chan int) // channel 必须在不同的goroutine间传递值

	go func() {
		fmt.Println("start gorountine")
		// 通过通道通知main的goroutine
		ch <- 1
		fmt.Println("exit goroutine")
	}()

	fmt.Println("wait goroutine")

	// 在main的goroutine中，等待接收 匿名goroutine 中的值
	fmt.Println(<-ch)
	fmt.Println("all down")
}


*/
