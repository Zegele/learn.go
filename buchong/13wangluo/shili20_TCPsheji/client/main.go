package main

import (
	"fmt"
	"net"
	"os"
)

//在TCP网络编程中，客户机的工作过程如下：
//1. TCP客户机在获取了服务器的服务端口号和服务地址之后
//可以调用DialTCP()函数向服务器发出连接请求
//如果请求成功会返回TCPConn对象
//2. 客户机调用TCPConn对象的Read()或Write()方法
//与服务器进行通信活动
//3. 通信完成后，客户机调用Close()方法关闭连接，断开通信链路

//DialTCP()函数原型定义如下：
//Func DialTCP(net string, laddr, raddr *TCPAddr)(*TCPConn, error)
//在调用函数时DialTCP()时，参数net时网络协议名，可以是“tcp”，“tcp4”，“tcp6”
//参数laddr是本地主机地址
//可以设为nil
//参数raddr是对方主机地址
//必须指定不能省略
//函数调用成功后，返回TCPConn对象，
//调用失败，返回一个错误类型

//方法Close()的原型定义如下：
//func (c *TCPConn) Close() error
//该方法调用成功后，关闭TCPConn连接，调用失败，返回一个错误类型：
//示例3： TCP Client端设计
//客户机通过内部测试地址“127.0.0.1”和端口5000和服务器建立通信连接

func main() {
	var buf [512]byte
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
	}

	service := os.Args[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp", service) //把网络地址转换为TCPAddr地址结构
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr) //拨号
	checkError(err)
	rAddr := conn.RemoteAddr() //获取服务端ip地址
	n, err := conn.Write([]byte("hello server"))
	checkError(err)
	n, err = conn.Read(buf[0:])
	checkError(err)
	fmt.Println("Reply from server: ", rAddr.String(), string(buf[0:n]))
	conn.Close()
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

//编译并运行：
//启动服务器 go run main.go
//启动客户机连接：./client 127.0.0.1:5000
//服务器响应：Receive from client: 127.0.0.1:50999 hello server
//客户机接收：Reply from server : 127.0.0.1:5000 Welcome client
//服务器注册了一个公知端口5000，而当客户机与服务器建立连接后
//客户机会生成一个临时端口“50999”与服务器进行通信
//服务器不管启动多少次端口号都是5000
//而客户端每次启动端口号都不一样
//
