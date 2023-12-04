// go语言UDP网络程序设计
// www.kancloud.cn/imdszxs/golang/1509692
package main

import (
	"fmt"
	"net"
	"os"
)

//UDP和上一节《TCP网络程序设计》中的TCP一样
//也工作在网络传输层
//但和TCP不同的是，它提供不可靠的通信服务
//UDP网络编程也为C-S模式
//要设计一个服务器，一个或多个客户机
//另外，UDP是不保证可靠性的通信协议
//所以客户机和服务器之间只要建立连接，
//就可以直接通信
//而不用调用Accept()进行连接确认
//本节将详细讲解UDP网络编程服务器，客户机的设计原理和设计过程

//UDPAddr地址结构体
//在进行UDP网络编程时，服务器或客户机的地址使用UDPAddr地址结构体表示，
//UDPAddr包括两个字段：IP和Port，形式如下：
//type UDPAddr struct{
//	IP IP
//	Port int
//}
//函数ResolveUDPAddr()可以把网络地址转换为UDPAddr地址结构
//该函数原型定义如下：
//func ResolveUDPAddr(net, addr string)(*UDPAddr, error)
//在调用函数ResolveUDPAddr()时，
//参数net是网络协议名
//可以是“udp”，“udp4” 或 “udp6”
//参数addr是IP地址或域名
//如果是IPv6地址必须“[]”括起来
//另外，端口号以“:”的形式跟随在IP地址或域名的后面，端口是可选的
//函数ResolveUDPAddr()调用成功后返回一个指向UDPAddr结构体的指针，否则返回一个错误类型
//另外，UDPAddr地址对象还有两个方法：Newwork() String()
//Network() 方法用于返回UDPAddr地址对象的网络协议名， 比如：“udp”
//String() 方法可以将UDPAddr地址转换成字符串形式，
//这两个方法原型定义如下：
//func(a *UDPAddr) Network()string
//func(a *UDPAddr) String()string

//UDPConn对象
//在进行UDP网络编程时，
//客户机和服务器之间是通过UDPConn对象实现连接的
//UDPConn是Conn接口的实现
//UDPConn对象绑定了服务器的网络协议和地址信息
//UDPConn对象定义如下：
//type UDPConn struct{
//	//空结构
//}
//通过UDPConn来连接对象在客户机和服务器之间进行通信
//UDP并不能保证通信的可靠性和有序性
//这些都要由程序员来处理
//为此，TCPConn对象提供了ReadFromUDP()方法和WriteToUDP()方法
//这两个方法直接使用远端主机地址进行数据发送和接收，
//即便在链路失效的情况下
//通信操作都能正常进行
//ReadFromUDP()方法和WriteToUDP()方法的原型定义如下：
//func (c *UDPConn)ReadFromUDP(b []byte)(n int, addr *UDPAddr, err error)
//func (c *UDPConn)WriteToUDP(b []byte, addr *UDPAddr)(int, error)
//ReadFromUDP()方法调用成功后，返回接收字节数和发送方地址
//否则返回一个错误类型
//WriteToUDP()方法调用成功后，返回发送字节数，否则返回一个错误类型
//
//UDP服务器设计
//在UDP网络编程中，服务器工作过程如下：
//1. UDP服务器首先注册一个公知端口，然后调用ListenUDP()函数在这个端口上创建一个UDPConn连接对象
//并在该对象上和客户机建立不可靠连接
//2. 如果服务器和某个客户机建立了UDPConn连接
//就可以使该对象的ReadFromUDP()方法和WriteToUDP()方法相互通信了
//3. 不管上一次通信是否完成或正常，UDP服务依然会接受下一次连接请求

//函数ListenUDP()原型定义如下：
//func ListenUDP(net string, laddr *UDPAddr)(*UDPConn, error)
//在调用函数ListenUDP()时
//参数net是网络协议名，可以是“udp”，“udp4”，“udp6”
//参数laddr是服务器本地地址，可以是任意活动的主机地址，
//或者是内部测试地址“127.0.0.1”
//该函数调用成功，返回一个UDPConn对象，
//调用失败，返回一个错误类型

//示例1： UDP Server端设计
//服务器使用本地地址，服务端口号为5001
//服务器设计工作模式采用循环服务器
//对每一个连接请求调用线程handleClient来处理/

func main() {
	serveice := ":5001"
	udpAddr, err := net.ResolveUDPAddr("udp", serveice)
	checkError(err)
	conn, err := net.ListenUDP("udp", udpAddr)
	checkError(err)
	for {
		handleClient(conn)
	}
}

func handleClient(conn *net.UDPConn) {
	var buf [512]byte
	n, addr, err := conn.ReadFromUDP(buf[:]) //读到buf里
	if err != nil {
		return
	}
	fmt.Println("Receive from client: ", addr.String(), string(buf[:n]))
	conn.WriteToUDP([]byte("Welcome Client"), addr)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

// UDP 客户机设计 //转client-main.go
