// TCP server端
package main

//tcp服务端
//一个tcp服务端可以同时连接很多个客户端，例如世界各地的用户使用自己电脑上的浏览器访问淘宝
//因为go语言中创建多个goroutine实现并发非常方便和高效，所以我们可以每建立一次连接就创建一个goroutine去处理
//tcp服务端程序的处理流程：
//1.监听端口
//2.接收客户端请求建立链接
//3.创建goroutine处理链接

import (
	"bufio"
	"fmt"
	"net"
)

// 处理函数
func process(conn net.Conn) {
	defer conn.Close() // 关闭链接
	for {
		reader := bufio.NewReader(conn) //net.Conn接口实现了Reader接口
		var buf [128]byte
		n, err := reader.Read(buf[:]) //读取（接收）数据
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到client端发来的数据：", recvStr)
		conn.Write([]byte(recvStr + "服务端回信")) // 发送给客户端
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:20000") //tcp协议
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	for {
		conn, err := listen.Accept() //建立连接
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go process(conn) //启动一个goroutine处理连接
	}
}

// 将上面代码保存，编译成：server.exe可执行文件
