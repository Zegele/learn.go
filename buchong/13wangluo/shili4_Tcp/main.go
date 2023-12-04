// go语言tcp协议
// www.kancloud.cn/imdszxs/golang/1509675
package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
)

//详细看文档
//tcp数据报主要包括
//SYN包：请求建立连接的数据包
//ACK包：回应数据包，表示接收到了对方的某个数据包
//PSH包：正常数据包
//FIN包：通讯结束包
//RST包：重置连接，导致TCP协议发送RST包的原因
//SYN包：数据段指定的目的端口处没有接收进程在等待
//TCP协议想放弃一个已经存在的连接
//TCP接收到一个数据段，但是这个数据段所标识的连接不存在
//接收到RST数据段的TCP协议立即将这条连接非正常地断开，并向应用程序报告错误
//URG包：紧急指针

//TCP三次握手
//建立一个tcp连接时，需要客户端和服务端总共发送3个包以确认连接的建立
//第一次握手
//客户端向服务器发出连接请求报文，这时报文首部中的同部位SYN=1，
//同时随机生成初始序列号seq=x,此时tcp客户端进程进入了syn-sent (同步已发送状态)状态。
//TCP规定SYN报文段（SYN=1的报文段）不能携带数据，但需要消耗掉一个序号
//这是三次握手中的开始，表示客户端想要和服务端建立连接
//第二次握手
//TCP服务器收到请求报文后，如果同意连接，则发出确认报文。
//。。。
//tcp服务器进程进入SYN-RCVD（同步收到）状态
//第三次握手
//tcp客户端进程收到确认后，还要向服务器给出确认，此时tcp连接建立，
//客户端进入ESTABLISHED（已建立连接）状态
//完成三次握手后，客户端与服务器即开始传送数据

//TCP四次挥手（Four-Way-Wavehand）
//即终止TCP连接，就是指断开一个TCP连接时，需要客户端和服务端总共发送4个包以确认连接的断开
//由客户端或服务端任意一方执行close来触发
//第一次挥手：
//Tcp发送一个FIN（结束），用来关闭客户到服务端的连接
//第二次挥手：
//服务端收到这个FIN，它发回一个ACK（确认）
//此时服务端就进入了CLOSE-WAIT（关闭等待）状态
//TCP服务器通知高层的应用进程，客户端已经没有数据要发送过来了
//但是服务器若发送数据，客户端依然要接收，这个状态还要持续一段时间，也就是整个CLOSE-WAIT专题持续的时间
//客户端收到服务器的确认请求后，此时客户端就进入FIN-WAIT-2（终止等待）状态，
//等待服务器连接释放报文（在这之前还需要接受服务器发送的最后的数据）
//第三次挥手
//服务端发送一个FIN（结束）到客户端，服务端关闭客户端的连接
//服务器将最后的数据发送完毕后，就向客户端发送连接释放报文,等待客户端的确认
//第四次挥手
//客户端发送ACK（确认）报文确认，这样关闭完成
//注意此时TCP连接还没有释放，必须经过2**MSL（最长报文段寿命）的时间后，
//当客户端撤销相应的TCB后，才进入CLOSED状态
//服务器只要收到了客户端发出的确认，立即进入CLOSED状态
//同意撤销TCB后，就结束了这次的TCP连接
//可以看到服务器结束TCP连接的时间要比客户端早一些

//为什么建立连接是三次握手，关闭连接却是四次挥手
//详看：www.kancloud.cn/imdszxs/golang/1509675

//示例：建立TCP连接来实现初步的HTTP协议：

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]
	conn, err := net.Dial("tcp", service)
	checkError(err)
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)
	result, err := readFully(conn)
	checkError(err)
	fmt.Println(string(result)) // 把[]byte转为string类型
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func readFully(conn net.Conn) ([]byte, error) {
	defer conn.Close()
	result := bytes.NewBuffer(nil)
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return result.Bytes(), nil
}

// 执行：go run main.go baidu.com:80
