// go语言TCP网络程序设计
// www.kancloud.cn/imdszxs/golang/1509691
package main

import (
	"fmt"
	"net"
	"os"
)

//TCP工作在网络的传输层
//它属于一种面向连接的可靠的通信协议
//TCP网络程序设计属于C-S模式，一般要设计一个服务器程序
//一个或多个客户机程序
//另外，TCP是面向连接的通信协议
//所以客户机要和服务器进行通信
//首先要在通信双方之间建立通信连接。
//本节将详细讲解TCP网络编程服务器，客户机的设计原理和设计过程

//TCPAddr地址结构体
//在进行TCP网络编程时，服务器或客户机的地址使用TCPAddr地质结构体表示
//TCPAddr包含两个字段：IP和Port，形式如下：
//type TCPAddr struct{
//	IP IP
//	Port int
//}

//函数ResolveTCPAddr()
//可以把网络地址转换为TCPAddr地址结构
//该函数原型定义如下：
//func ResolveTCPAddr(net, addr string)(*TCPAddr, error)
//在调用函数ResolveTCPAddr()时，参数net时网络协议名，
//可以是“tcp”，"tcp4"或“tcp6”
//参数addr是IP地址或域名
//如果是IPv6地址则必须使用“[]”括起来
//另外，端口号以“:”的形式跟随在IP地址或域名的后面
//端口是可选的
//例如：www.google.com:80  或  127.0.0.1:21
//还有一种特例，
//就是对于HTTP服务器
//当主机地址为本地测试地址时（127.0.0.1）
//可以直接使用端口号作为TCP连接地址，形如":80"(其实就是127.0.0.1：80的缩写)
//函数ResolveTCPAddr()调用成功后返回一个指向TCPAddr结构体的指针
//否则返回一个错误类型
//另外，TCPAddr地址对象还有两个方法：Network()和String()
//Netwrok()方法用于返回TCPAddr地址对象的网络协议名 ，比如“tcp”;
//String()方法可以将TCPAddr地址转换层字符串形式
//这两个方法原型定义如下：
//func(a *TCPAddr)Network()string
//func(a *TCPAddr)String()string
//示例1：TCP连接地址：

/*
func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s networkType addr\n", os.Args[0])
		os.Exit(1)
	}
	networkType := os.Args[1]
	addr := os.Args[2]
	tcpAddr, err := net.ResolveTCPAddr(networkType, addr)
	if err != nil {
		fmt.Println("ResolveTCPAddr error: ", err.Error())
		os.Exit(1)
	}
	fmt.Println("The network type is: ", tcpAddr.Network())
	fmt.Println("The IP address is: ", tcpAddr.String())
	os.Exit(0)
}

*/

//./main tcp c.biancheng.net:80
//The network type is:  tcp
//The IP address is:  222.75.48.146:80

//TCPConn 对象
//在进行TCP网络编程时
//客户机和服务器之间是通过TCPConn对象实现连接的
//TCPConn是Conn接口的实现
//TCPConn对象绑定了服务器的网络协议和地址信息
//TCPConn对象定义如下：
//type TCPConn struct{
//	//空结构
//}
//通过TCPConn连接对象，
//可以实现客户机和服务器的全双工通信
//可以通过TCPConn对象的Read()方法和Write()方法
//在服务器和客户机之间发送和接收数据
//Read()方法和Write()方法的原型定义如下：
//func (c *TCPConn)Read(b []byte)(n int, err error)
//func (c *TCPConn)Write(b []byte)(n int, err error)
//Read()方法调用成功后会返回接收到的字节数
//调用失败返回一个错误类型
//Write()方法调用成功后会返回正确发送的字节数
//调用失败返回一个错误类型
//另外，这两个方法的执行都会引起阻塞
//
//TCP服务器设计
//前面讲了GO语言网络编程和传统Socket网络程序有所不同
//TCP服务器的工作过程如下：
//1. TCP服务器首先注册一个公知端口，
//然后调用ListenTCP()函数在这个端口上创建一个TCPListener监听对象
//并在该对象上监听客户机的连接请求
//2. 启用TCPListener对象的Accept()方法接收客户机的连接请求
//并返回一个协议相关的Conn对象
//这里就是TCPConn对象
//3. 如果返回了一个新的TCPConn对象，
//服务器就可以调用该对象的Read()方法接收客户机发来的数据
//或者调用Write()方法向客户机发送数据了
//
//TCPListener对象，ListenTCP()函数的原型定义如下：
//type TCPListener struct{
// 	//contains filtered or unexported fields
//}
//func ListenTCP(net string, laddr *TCPAddr)(*TCPListener, error)
//在调用函数ListenTCP()时，参数net是网络协议名，可以是“tcp”，“tcp4”，或“tcp6”
//参数laddr是服务器本地地址，可以是任意活动的主机地址
//或者是内部测试地址“127.0.0.1”
//该函数调用成功，返回一个TCPListener对象
//调用失败，返回一个错误类型

//TCPListener对象的Accept()方法原型定义如下：
//func(I *TCPListener)Accept()(c Coon, err error)
//Accep()方法调用成功后，返回TCPConn对象，否则返回一个错误类型
//服务器和客户机的通信连接建立成功后
//就可以使用Read(),Write()方法收发数据
//在通信过程中，如果还想获取通信双方的地址信息
//可以使用LocalAddr()方法和RemoteAddr()方法来完成
//这两个方法原型定义如下：
//func(c *TCPConn)LocalAddr()Addr
//func(c *TCPConn)RemoteAddr()Addr
//LocalAddr()方法会返回本地主机地址，
//RemoteAddr()方法返回远端主机地址
//示例2：TCP Server端设计
//服务器使用本地地址，服务端口号：5000
//服务器设计工作模式采用循环服务器，
//对每一个连接请求调用线程handleClient来处理

/*
func main() {
	service := ":5000"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)
	fmt.Println("tcpAddr:", tcpAddr.String()) //tcpAddr: :5000

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		handleClient(conn)
		conn.Close()
	}
}
func handleClient(conn net.Conn) {
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		if err != nil {
			return
		}
		rAddr := conn.RemoteAddr()
		fmt.Println("Receive from client", rAddr.String(), string(buf[0:n]))
		_, err2 := conn.Write([]byte("Welcome client"))
		if err2 != nil {
			return
		}

	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error %s", err.Error())
		os.Exit(1)
	}
}

*/

// TCP 客户机设计 （见client-main.go）

//使用Goroutine实现并发服务器
//前面的讲解中，服务器设计采用循环服务器设计模式
//这种服务器设计简单但缺陷明显
//因为这种服务器一旦启动
//就一直阻塞监听客户机的连接请求
//直至服务器关闭
//所以，循环服务器很耗费系统资源
//解决问题的方法是采用并发服务器模式，
//在这种模式中，对每一个客户端的连接请求
//服务器都会创建一个新的进程，线程或协程 进行响应
//而服务器还可以去处理其他任务
//Goroutine即协程是一种比线程更轻量级的任务单位
//所以这里就使用Goroutine来实现并发服务器的设计

//示例4：并发服务器设计
//服务器使用本地地址，服务端口号为5000
//服务器设计工作模式采用并发服务器模式，
//对每一个连接请求创建一个能调用handleClient()函数的Goroutine来处理

func main() {
	service := ":5000"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn) // 创建goroutine
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close() // 逆序调用 Close()保证连接能正常关闭
	var buf [512]byte
	for {
		n, err := conn.Read(buf[:])
		if err != nil {
			return
		}
		rAddr := conn.RemoteAddr()
		fmt.Println("Receive from client: ", rAddr.String(), string(buf[:n]))
		_, err2 := conn.Write([]byte("Welcome client"))
		if err2 != nil {
			return
		}
	}
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error %s", err.Error())
		os.Exit(1)
	}
}

//编译并运行
//通过测试可以发现，并发服务器可以同时响应多个客户机的连接请求
//并能和多个客户机并发通信，
//尤其在多核心系统平台上
//这种通信模式效率更高。
//而循环服务器只能按客户机的请求队列次序
//一个一个地为客户机提供通信服务，通信效率低下
///
