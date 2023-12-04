package main

import (
	"fmt"
	"net"
	"os"
)

//UDP客户机设计
//在UDP网络编程中，客户机工作过程如下：
//1. UDP客户机在获取了服务器的服务端口号和服务地址之后
//可以调用DialUDP()函数向服务器发出连接请求
//如果请求成功会返回UDPConn对象
//2. 客户机可以直接调用UDPConn对象的ReadFromUDP()方法或WriteToUDP()方法
//与服务器进行通信活动
//3. 通信完成后，客户机调用Close()方法关闭UDPConn连接
//断开通信链路

//函数DialUDP()原型定义如下：
//func DialUDP(net string, laddr, raddr *UDPAddr)(*UDPConn, error)
//在调用函数DialUDP()时，
//参数net是网络协议名 可以是“udp”，“udp4”，“udp6”
//参数laddr是本地主机地址，可以设为nil
//参数raddr是对方主机地址，必须指定不能省略
//函数调用成功后，返回UDPConn对象，调用失败，返回一个错误类型

//方法Close()的原型定义如下：
//func (c *UDPConn)Close()error
//该方法调用成功后，关闭UDPConn连接
//调用失败，返回一个错误类型

//示例2：UDP Client端设计
//客户机通过内部测试地址“127.0.0.1:5001”和服务器建立通信连接

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
	}
	service := os.Args[1]
	udpAddr, err := net.ResolveUDPAddr("udp", service)
	checkError(err)
	conn, err := net.DialUDP("udp", nil, udpAddr)
	checkError(err)
	_, err = conn.Write([]byte("Hello server!"))
	checkError(err)
	var buf [512]byte
	n, addr, err := conn.ReadFromUDP(buf[:])
	checkError(err)
	fmt.Println("Reply from server: ", addr.String(), string(buf[:n]))
	conn.Close()
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

// 服务端和客户端编译运行后
// ./client.exe 127.0.0.1:5001
//服务端响应：Receive from client:  127.0.0.1:57996 Hello server!
//客户端接收：Reply from server:  127.0.0.1:5001 Welcome Client

//通过测试会发现，采用TCP时必须先启动服务器
//然后才能正常启动客户机
//如果服务器中断，则客户机也会异常退出
//而采用UDP时，客户机和服务器启动没有先后次序
//而且即便是服务器异常退出，客户机也能正常工作
//总之，TCP可以保证客户机，服务器双方按照可靠有序的方式进行通信
//但通信效率低
//而UDP虽然不能保证通信的可靠性，但通信效率要高得多
//在有些场合是非常有用的。
