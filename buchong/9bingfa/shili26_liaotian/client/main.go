// 是一个简单的TCP服务器读/写客户端
package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() { // 得看看小强课程
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() { // 这个goroutine是干什么的？？？
		io.Copy(os.Stdout, conn) //注意：忽略错误
		log.Println("done")
		done <- struct{}{} // 向主Goroutine发出信号
	}()
	mustCopy(conn, os.Stdin)
	conn.Close() // 什么情况下，conn就close了？？？
	<-done       // 等待后台goroutine完成
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

// 服务端中，clients map 限制在广播器这一个goroutine中被访问，所以不会并发访问它。
// 唯一被多个goroutine共享的变量是通道以及net.Conn的实例，他们又都是并发安全的。
