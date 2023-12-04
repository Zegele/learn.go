// www.kancloud.cn/imdszxs/golang/1509646
// Go语言使用通道响应计时器的事件
package main

import (
	"fmt"
	"time"
)

// 定点计时
//Timer 计时器： 给定多少时间后触发。  使用time.Timer创建
//Ticker 打点器 ： 计时触发。 使用time.Ticker创建
// 里面通过一个C成员，类型是只能接收的时间通道（<-chan Time）（单向的），使用这个通道就可以获得时间触发的通知。

// 例子：创建一个打点器，每500毫秒触发一次，创建一个计时器，2秒后触发，只触发一次
func main() {
	// 创建一个打点器，每500毫秒触发一次
	ticker := time.NewTicker(time.Millisecond * 500)
	//创建一个计时器，2秒后触发
	stopper := time.NewTimer(time.Second * 2)
	// 声明计数变量
	var i int
	//不断地检查通道情况
	for {
		// 多路复用通道
		select {
		case <-stopper.C: //计时器到时了
			fmt.Println("stop")
			// 跳出循环
			goto StopHere
		case <-ticker.C: //打点器触发了
			// 记录触发了多少次
			i++
			fmt.Println("tick", i)
		}
	}
	// 退出的标签，使用goto跳转
StopHere:
	fmt.Println("done")

}

/*
// 一段时间之后（time.After）
func main() {
	// 声明一个退出用的通道， 想表示：往这个通道里写数据表示退出。
	exit := make(chan int)
	//打印开始
	fmt.Println("start")
	// 过1秒后，调用匿名函数
	time.AfterFunc(time.Second, func() {// 时间到达后，该匿名函数会在另外一个goroutine中被调用
		// 1秒后，打印结果
		fmt.Println("one second after")
		// 通知main（）的goroutine已经结束
		exit <- 0
	})
	//等待结束
	<-exit
}
// time.AfterFunc() 函数是在time.After 基础上增加了到时的回调，方便使用。



*/
