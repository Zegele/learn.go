// go语言IP网络程序设计
// www.kancloud.cn/imdszxs/golang/1509693
package main

import (
	"bytes"
	"fmt"
	"net"
	"os"
	"time"
)

//IP 是Internet网络层的核心协议
//是一种不可靠的，无连接的通信协议
//TCP，UDP都是在IP的基础上实现的通信协议
//所以IP属于一种底层协议，它可以直接对网络数据包（package） 进行处理
//另外，通过IP用户还可以实现自己的网络服务协议
//本节将详细讲解IP网络编程服务器，客户机的设计原理和设科过程

//IPAddr地址结构体
//在进行IP网络编程时，服务器或客户机的地址使用IPAddr地址结构表示
//IPAddr结构体只有一个字段IP，形式如下：
//type IPAddr struct{
//	IP IP
//}
//通过了解IPAddr地址结构可以发现，IP网络编程属于一种底层网络程序设计
//它可以直接对IP包进行处理
//所以IPAddr地址中没有端口地址，这个和TCPAddr地址结构，UDPAddr地址结构都不同，
//在应用时要特别注意

//函数ResolveIPAddr()可以把网络地址转换为IPAddr地址结构，该函数原型定义如下：
//func ResolveIPAddr()函数时，参数net表示网络类型，可以是"ip","ip4","ip6"
//参数addr是IP地址或域名，如果是IPv6地址必须使用“[]”括起来
//函数ResolveIPAddr()调用成功后返回一个指向IPAddr结构体的指针
//否则返回一个错误类型
//另外，IPAddr地址对象还有两个方法：Network()和String()
//NetWork()方法用于返回IPAddr地址对象的网络协议名，比如“ip”
//String()方法可以将IPAddr地址转换成字符串形式
//这两个方法原型定义服下：
//func (a *IPAddr)Network()string
//func (a *IPAddr)String()string

//IPConn对象
//在进行IP网络编程时，客户机和服务器之间是通过IPConn对象实现连接的
//IPConn是Conn接口的实现
//IPConn对象绑定了服务器的网络协议和地址信息，
//IPConn对象定义如下：
//type IPConn struct{
//	//空结构体
//}
//由于IPConn是一个无连接的通信对象，
//所以IPConn对象提供了ReadFromIP()方法和WriteToIP()方法用于在客户机和服务器之间进行数据收发操作
//原型定义如下：
//func(c *IPConn)ReadFromIP(b []byte)(int, *IPAddr,error)
//func(c *IPConn)ReadFromIP(b []byte, addr *IPAddr)(int, error)
//ReadFromIP()方法调用成功后返回接收字节数和发送方地址，否则返回一个错误类型
//WriteToIP()方法调用成功后返回发送字节数，否则返回一个错误类型

//IP服务器设计
//由于工作在网络层，ip服务器并不需要在一个指定的端口上和客户机进行通信连接
//IP服务器的工作过程如下：
//1. IP服务器使用指定的协议簇和协议，调用ListenIP()函数创建一个IPConn连接对象
//并在该对象和客户机建立不可靠连接
//2. 如果服务器和某个客户机建立了IPConn连接，就可以使用该对象的ReadFromIP()方法
//和WriteToIP()方法相互通信了
//3. 如果通信结束，服务器还可以调用Close()方法关闭IPConn连接

//函数ListenIP()原型定义如下：
//func ListenIP(netProto string, laddr *IPAddr)(*IPConn, error)
//在调用函数ListenIP()时，参数netProto是“网络类型+协议名”或“网络类型+协议号”
//中间用“:”隔开 比如“IP4:IP”或“IP4:4”
//参数laddr是服务器本地地址
//可以是任意活动的主机地址，或者是内部测试地址“127.0.0.1”
//该函数调用成功，返回一个IPConn对象
//调用失败，返回一个错误类型

//示例1：
//IP Server端设计
//服务器使用本地主机地址，调用Hostname()函数获取
//服务器设计工作模式采用循环服务器
//对每一个连接请求调用线程handleClient来处理。

/*
func main() {
	name, err := os.Hostname()
	checkError(err)

	ipAddr, err := net.ResolveIPAddr("ip4", name)
	checkError(err)
	fmt.Println(ipAddr)
	conn, err := net.ListenIP("ip4:ip", ipAddr)
	// 这步报错为啥？？
	//报错 socket: An attempt was made to access a socket in a way forbidden by its access perm
	//端口被用了？？
	checkError(err)
	for {
		handleClient(conn)
	}
}

func handleClient(conn *net.IPConn) {
	var buf [512]byte
	n, addr, err := conn.ReadFromIP(buf[:])
	if err != nil {
		return
	}
	fmt.Println("Receive from Client: ", addr.String(), string(buf[:n]))
	conn.WriteToIP([]byte("Welcome Client."), addr)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
//编译运行不了
//	conn, err := net.ListenIP("ip4:ip", ipAddr)
// 这步报错为啥？？
//报错 socket: An attempt was made to access a socket in a way forbidden by its access perm


*/

// IP 客户机设计（见client-main.go）

//Ping程序设计
//不管是UNIX还是Windows系统中都有一个Ping命令
//利用它可以检查网络是否连通
//分析判断网络故障
//Ping会向目标主机发送测试数据包
//看对方时候有响应并统计响应时间，以此测试网络
//Ping命令的这些功能是使用IP层的ICMP实现的
//在测试过程中，源主机向目标主机发送回显请求报文 （ICMP_ECHO_REQUEST,type=8,code=0）
//目的主机返回回显响应报文（ICMP_ECHO_REPLY,type=0,code=0）
//相关的数据报格式如下图所示：
//图：ICMP回显请求和响应数据报格式
//其中，标识符是源主机的进程号，
//序列码用来标识发出回显请求的次序
//时间戳表示数据报发出的时刻
//通过比较回显响应时刻和源主机当前时刻的差值
//可以测出ICMP数据包的往返时间

//示例3：使用原始套接字和ICMP设计Ping程序
//函数makePingRequest() 的功能是生成ICMP请求包
//函数parsePingReply()用于解析目标主机发回的相应包
//函数elapsedTime()的功能是计算ICMP数据报往返时间

const (
	ICMP_ECHO_REQUEST = 8
	ICMP_ECHO_REPLY   = 0
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host", os.Args[0])
		os.Exit(1)
	}
	dst := os.Args[1]
	raddr, err := net.ResolveIPAddr("ip4", dst)
	checkError(err)
	ipconn, err := net.DialIP("ip4:icmp", nil, raddr)
	fmt.Println(111111)
	checkError(err)
	fmt.Println(222222)
	sendid := os.Getpid() & 0xfff
	sendseq := 1
	pingpktlen := 64
	for {
		sendpkt := makePingRequest(sendid, sendseq, pingpktlen, []byte(" ")) //这个[]byte("")有问题
		start := int64(time.Now().Nanosecond())
		_, err := ipconn.WriteToIP(sendpkt, raddr)
		fmt.Println(3333333)
		checkError(err)
		fmt.Println(4444444)
		resp := make([]byte, 1024)
		for {
			n, from, err := ipconn.ReadFrom(resp)
			fmt.Println(55555555)
			checkError(err)
			fmt.Println(66666666)
			fmt.Printf("%d bytes from %s: icmp_req = %d time = %.2f ms\n", n, from, sendseq, elapsedTime(start))
			if resp[0] != ICMP_ECHO_REPLY {
				continue
			}
			rcvid, rcvseq := parsePingReply(resp)
			if rcvid != sendid || rcvseq != sendseq {
				fmt.Printf("Ping reply saw id", rcvid, rcvseq, sendid, sendseq)
			}
			break
		}
		if sendseq == 4 {
			break
		} else {
			sendseq++
		}
		time.Sleep(1e9)
	}
}

func makePingRequest(id, seq, pktlen int, filler []byte) []byte {
	p := make([]byte, pktlen)
	copy(p[8:], bytes.Repeat(filler, (pktlen-8)/(len(filler)+1))) // IP运算出了问题
	p[0] = ICMP_ECHO_REQUEST
	//type
	p[1] = 0
	//cksum
	p[2] = 0
	//cksum
	p[3] = 0
	//id
	p[4] = uint8(id >> 8)
	//id
	p[5] = uint8(id & 0xff)
	// id
	p[6] = uint8(seq >> 8)
	//sequence
	p[7] = uint8(seq & 0xff)
	//sequence
	cklen := len(p)
	s := uint32(0)
	for i := 0; i < (cklen - 1); i += 2 {
		s += uint32(p[i+1])<<8 | uint32(p[i]) // 这个 “|” 是什么意思？？？
	}
	if cklen&1 == 1 {
		s += uint32(p[cklen-1])
	}
	s = (s >> 16) + (s & 0xffff)
	s = s + (s >> 16)
	p[2] ^= uint8(^s & 0xff)
	p[3] ^= uint8(^s >> 8)
	return p
}

func parsePingReply(p []byte) (id, seq int) {
	id = int(p[4])<<8 | int(p[5])
	seq = int(p[6])<<8 | int(p[7])
	return
}

func elapsedTime(start int64) float32 {
	t := float32((int64(time.Now().Nanosecond()) - start) / 1000000.0)
	return t
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

// 不知道哪里错了
