// 一个只读的TCP客户端程序
package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

//这个程序从网络连接中读取，然后写到标准输出
//直到到达EOF或者出错
//mustCopy函数是这一节的多个例子中使用的一个使用程序
//在不同的终端上同时运行两个客户端，一个显示在作弊那一个在右边
//go build client
// ./client
//killall server
//killall命令是UNIX的一个实用程序 (所以windows用不了？？？)
//用来终止所有指定名字的进程
//第二个客户端必须等到第一个结束才能正常工作
//因为服务器是顺序的
//一次只能处理一个客户请求
//让服务器支持并发只需要一个很小的改变：
//在调用handleconn的地方添加一个go关键字
//使它在自己的goroutine内执行
//，
