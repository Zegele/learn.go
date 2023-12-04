// go语言Dial()函数：建立网络连接
// www.kancloud.cn/imdszxs/golang/1509673
package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
)

//go语言中Dial()函数用于创建网络连接，函数原型如下：
//func Dial(nerwork, address string) (Conn, error) {
//	var d net.Dialer
//	return d.Dial(network, address)
//}

//network :参数表示传入的网络协议（如tcp，udp等 ）
//address：表示传入的IP地址或域名，而端口号是可选的，
//如果需要指定的话，以 : 的形式跟在地址或域名后面即可
//如果连接成功，该函数返回连接对象，否则返回error

//实际上，Dial()函数是对DialTCP(),DialUDP(),DialIP(),DialUnix()函数的封装：
//func DialTCP(net string, laddr, raddr *TCPAddr)(c *TCPConn, err error)、
//func DialUDP(net string, laddr, raddr *UDPAddr)(c *UDPConn, err error)、
//func DialIP(netProto string, laddr, raddr *IPAddr)(c *IPConn, err error)、
//func DialUnix(net string, laddr, raddr *UnixAddr)(c *UnixConn, err error)、

//常见的调用方式：
//1. TCP连接
//conn, err := net.Dial("tcp","192.168.10.10:80")
//2. UDP连接
//conn, err := net.Dial("udp", "192.168.10.10:8888")
//3. ICMP连接
//conn, err := net.Dial("ip4:icmp","c.biancheng.net")
//提示：ip4表示IPv4，相应的ip6表示IPv6
//4. ICMP连接（使用协议编号）
//conn, err := net.Dial("ip4:1","10.0.0.3")
//提示：我们可以通过以下连接查看协议编号的含义
//www.iana.org/assignments/protocol-numbers/protocol-numbers.xml

//目前 Dial()函数支持以下网络协议：tcp,tcp4,tcp6,udp,udp4,udp6,ip,ip4,ip6,unix,unixgram,unixpacket
//在成功建立连接后，我们就可以进行数据的发送和接收
//发送数据时使用连接对象conn的Write()方法，接收数据时使用Read()方法
//演示：tcp实现http协议，通过向网络主机发送http head请求，读取网络主机返回的信息

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}
	// 从参数中读取主机信息
	service := os.Args[1]
	//建立网络连接
	conn, err := net.Dial("tcp", service)
	//连接出错则打印错误消息并退出程序
	checkError(err)
	// 调用返回的连接对象提供的Write方法发送请求
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)
	// 通过连接对象提供的Read方法读取所有响应数据
	result, err := readFully(conn)
	checkError(err)
	// 打印响应数据
	fmt.Println(string(result))
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func readFully(conn net.Conn) ([]byte, error) {
	// 读取所有响应数据后主动关闭连接
	defer conn.Close()
	result := bytes.NewBuffer(nil)
	var buf [512]byte
	for {
		n, err := conn.Read(buf[:])
		result.Write(buf[:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return result.Bytes(), nil
}

//运行： go run main.go c.biancheng.net:80
//或运行： go run main.go c.biancheng.net:http
//HTTP/1.1 400 Bad Request
//Server: Tengine
//Date: Fri, 10 Mar 2023 08:53:11 GMT
//Content-Type: text/html
//Content-Length: 249
//Connection: close
//X-Tengine-Error: empty host
//Via: kunlun10.cn5037[,0]
//Timing-Allow-Origin: *
//EagleId: 7d40823a16784383918594662e
//可以看到，通过go语言编写的网络程序整体实现代码非常简单清晰，
//就是建立连接，发送数据、接收数据、不需要我们关注底层不同协议通信的细节
