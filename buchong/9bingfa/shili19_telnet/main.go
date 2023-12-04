package main

//www.likecs.com/show-306390433.html#sc=3773.3046875
// 这个逻辑要读懂，很重要
import (
	"learn.go/buchong/9bingfa/shili19_telnet/server"
	"os"
)

func main() {
	// 创建一个程序结束码的通道
	exitChan := make(chan int)

	// 将服务器并发运行
	go server.Server("127.0.0.1:7001", exitChan)
	//在操作系统中的命令行中输入： telnet 127.0.0.1 7001
	// 没搞懂这个是什么意思，解决：启用windows的telnet 客户端功能

	// 通道阻塞，等待接收返回值
	code := <-exitChan

	// 标记程序返回值，并退出
	os.Exit(code)
}
