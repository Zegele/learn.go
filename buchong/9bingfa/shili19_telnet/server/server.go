package server

import (
	"fmt"
	"learn.go/buchong/9bingfa/shili19_telnet/client"
	"net"
)

//www.kancloud.cn/imdszxs/golang/1509649

// 服务逻辑，传入地址核退出的通道
func Server(address string, exitChan chan int) { // address为传入的地址， 退出服务器使用exitChan通道控制
	// 根据给定地址进行侦听
	l, err := net.Listen("tcp", address) //net.Listen()函数进行侦听
	// 第一个参数为协议类型，本例需要做的是TCP连接 因此填入tcp；
	// 第二个参数为地址，格式为：主机 + 端口号

	// 如果侦听发生错误，打印错误并退出
	if err != nil {
		fmt.Println(err.Error())
		exitChan <- 1
	}
	// 打印侦听地址，表示侦听成功
	fmt.Println("listen: " + address)
	// 延迟关闭侦听
	defer l.Close()
	// 侦听循环
	for {
		// 新连接没有到来时，Accept是阻塞的
		conn, err := l.Accept()
		//在没有连接时，Accept()函数调用后会一直阻塞。
		// 连接到来时，返回conn和错误变量，conn的类型时*tcp.Conn

		// 发生任何的侦听错误，打印错误并退出服务器
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		// 根据连接开启会话，这个过程需要并行执行
		go session.HandleSession(conn, exitChan)
		// 每个连接会生成一个会话，这个会话的处理与接受逻辑需要并行执行，彼此不干扰
	}
}
