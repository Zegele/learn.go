package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

//c.biancheng.net/view/5408.html
/*
服务端程序
服务端程序包含4个goroutine，分别是一个主goroutine，和广播（broadcaster）goroutine
每一个连接里面又包含一个连接处理（handleConn）goroutine和一个客户写入（clientwrite）goroutine。

广播器（broadcaster）是用于如何使用select的一个规范说明，因为它需要对三种不同的消息进行响应。

主goroutine的工作是监听端口，接受连接客户端的网络连接，每一个连接，它将创建一个新的handleConn goroutine

如下：
*/

func main() {
	// main 函数，服务器要做的事情就是1. 获得listener对象，然后2. 不停的获取链接上来的conn对象，
	//最后3.把conn对象丢给处理链接的函数。
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue // 跳下次循环
		}
		go handleConn(conn)
		// 在使用handleConn方法处理conn对象的时候，对不同的链接都启一个goroutine去并发处理每个conn这样则无需等待。
		// 由于要给所有在线的用户发送消息，而不同的用户conn对象都在不同的goroutine里，
		// 所以要使用go语言中的channel来处理不同goroutine之间的消息传递。
		// 所以使用了不同的channel（全局，以及handleConn内），来传递消息。
	}

}

type client chan<- string // 对外发送消息的通道 这个不需要往外传数据

var (
	entering = make(chan client) // 通道类型的通道（client也是通道）
	leaving  = make(chan client)
	messages = make(chan string) // 所有连接的客户端
)

func broadcaster() {
	clients := make(map[client]bool)
	// 记录当前链接的客户集合
	//每个客户唯一被记录的信息是其对外发送消息通道的ID。
	for {
		select {
		case msg := <-messages: // messages是哪里给的？
			// 把所有接收到的消息广播给所有客户端
			// 发送消息通道
			for cli := range clients {
				cli <- msg
			}

			//用户来了，true
		case cli := <-entering: // cli 本身就是chan类型的（且是只接收），还能这样用，学习了。
			clients[cli] = true

			//用户走了，clents的map中删掉对应client的key
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string) // 对外发送客户消息的通道
	// 创建一个对外发送消息的新通道。
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String() // 这是干嘛？ id地址？
	ch <- "欢迎" + who
	messages <- who + "上线" // 全局接收消息通道，每条信息前有id
	entering <- ch         //channel 对另一个channel 传递消息。
	// 问题这两个channel 接收到信息后，在广播那有先后顺序么？

	input := bufio.NewScanner(conn) //输入的消息？
	// 这里参数是io.Reader类型，为啥conn类型能放这里？
	// 因为io.Reader类型是接口类型，里面有个read方法，conn接口里也有read方法，方法一致所以可以用在这里。
	// 对接口的理解和使用还要再熟悉熟悉！！！
	for input.Scan() { // 返回值是个布尔值
		messages <- who + ": " + input.Text()
	}
	// 注意：忽略input.Err() 中可能的错误

	leaving <- ch //如果关闭了客户端，那么把队列离开写入leaving交给广播函数去删除这个客户端并关闭这个客户端。
	// 这个是怎么触发的？

	messages <- who + " 下线"
	conn.Close() // 关闭这个客户端的连接。conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // 注意：忽略网络层面的错误
	}
}
