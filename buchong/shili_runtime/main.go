package main

import (
	"fmt"
	"runtime"
)

func say(s string) {
	for i := 0; i < 2; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

func main() {
	for i := 0; i < 10; i++ {

		go say("world")
		say("hello")
		fmt.Println("----------")
	}
}

// 假设先执行go（world） - rt - 执行 hello  - rt - 打印 world  - 执行 go world - rt - 打印 hello - 执行 hello - rt -- 打印world 完成go world  打印hello
// 假设先执行hello -rt - 执行 world - rt 打印hello - 执行 hello循环 - rt - 打印world - 执行world循环，- rt 打印hello ,主程序结束，所有goroutine结束
