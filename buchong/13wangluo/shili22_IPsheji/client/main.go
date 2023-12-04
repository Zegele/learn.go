package main

import (
	"fmt"
	"net"
	"os"
)

//在ip网络编程中，客户机工作过程如下：
//1. IP客户机在获取了服务器的网络地址之后，
//可以调用DialIP()函数向服务器发出连接请求，
//如果请求成功会返回IPConn对象
//2. 如果连接成功，客户机可以直接调用IPConn对象的ReadFromIP()方法或WriteToIP()方法
//与服务器进行通信活动
//3. 通信完成后，客户机调用Close()方法关闭IPConn连接，断开通信链路

//函数DialIP()原型定义如下：
//func DialIP(netProto string, laddr, raddr *IPAddr)(*IPConn, error)
//在调用函数DialIP()时，参数netProto是“网络类型+协议名”或“网络类型+协议号”
//中间中“:”隔开，比如“IP4:IP”,"IP4:4"
//参数laddr是本地主机地址，可以设为nil
//参数raddr是对方主机地址，
//必须指定不能省略
//函数调用成功后，返回IPConn对象
//调用失败，返回一个错误类型

//方法Close()的原型定义如下：
//func (c *IPConn)Close()error
//该方法调用成功后，关闭IPConn连接，调用失败，返回一个错误类型
//示例2： IP Client端设计
//客户机通过内部测试地址“127.0.0.1”和服务器建立通信连接
//服务器主机地址可以使用Hostname()函数获取.

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host: port", os.Args[0])
	}
	service := os.Args[1]
	lAddr, err := net.ResolveIPAddr("ip4", service)
	checkError(err)
	name, err := os.Hostname()
	checkError(err)
	rAddr, err := net.ResolveIPAddr("ip4", name)
	checkError(err)
	conn, err := net.DialIP("ip4:ip", lAddr, rAddr)
	checkError(err)
	_, err = conn.WriteToIP([]byte("Hello Server!"), rAddr)
	checkError(err)
	var buf [512]byte
	n, addr, err := conn.ReadFromIP(buf[:])
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

// 编译并运行
//启动服务器： ./server.exe  // 启动报错
//客户机连接：./client.exe 127.0.0.1
//
//可以看出，TCP，UDP的服务器和客户机通信时必须使用端口号
//而IP服务器和客户机之间通信不需要端口号
//另外，如果在同一台计算机上，服务器，客户机要使用不同的地址进行测试
//比如本例服务器地址“172.18.128.1”，客户机使用内部测试地址“127.0.0.1”
//如果使用相同的地址，会发生自发自收的现象，原因是IP是底层通信，
//并没有像TCP，UDP那样使用端口号来区分不同的通信进程
//
