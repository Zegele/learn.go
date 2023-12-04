// 并发始终服务器
// www.kancloud.cn/imdszxs/golang/1509685
package main

import (
	"io"
	"log"
	"net"
	"time"
)

//网络是一个自然使用并发的领域，因为服务器通常一次处理很多来自客户端的连接
//每一个客户端通常和其他客户端保持独立
//本节介绍net包，它提供构建客户端和服务器程序的组件
//这些程序通过TCP，UDP或者UNIX套接字进行通信
//net/http包就是在net包基础上构建的
//示例：顺序始终服务器，它以每秒一次的频率向客户端发送当前时间，如下；

// clock1 是一个定期报告时间的TCP服务器，
func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // 例如，连接中止
			continue
		}
		//handleConn(conn) // 一次处理一个连接
		go handleConn(conn) // 并发处理连接
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	for {
		_, err := io.WriteString(conn, time.Now().Format("15:04:05\n"))
		if err != nil {
			return //例如，连接断开
		}
		time.Sleep(1 * time.Second)
	}
}

//Listen函数创建一个net.Listener对象，它在一个网络端口上监听进来的连接，
//这里是TCP端口localhost:8000
//监听器的Accept方法被阻塞，直到有连接请求进来，然后返回net.Conn对象来代表一个连接
//handleConn函数处理一个完整的客户连接，
//在循环里，它将time.Now()获取的当前时间发送给客户端
//因为net.Conn满足io.Writer接口，
//所以可以直接向它进行写入
//当写入失败时循环结束，很多时候是客户端断开连接，这时handleconn函数使用延迟的Close调用关闭自己这边的连接
//然后继续等待下一个连接请求
//time.Time.Format方法提供了格式化日期和时间信息的方式
//它的参数是一个模板，只是如何格式化一个参考时间，具体如 Mon Jan 2 03：04：05PM 2006 UTC-0700这样的形式
//参考时间有8个部分（本周第几天，月，本月第几天，等）
//他们可以以任意的组合和对应数目的格式化字符出现在格式化模板中，所选择的日期和时间将通过所选择的格式进行显示
//这里只使用时间的小时，分钟和秒部分
//time包定义了许多标准时间格式的模板
//如time.RFC1123
//相反，当解析一个代表时间的字符串的时候使用相同的机制
//为了连接到服务器，需要一个像nc("netcat")这样的程序
//以及一个用来操作网络连接的标准工具：
//go build main.go
//./main &
//nc localhost 8000
//
//
//客户端先似乎每秒从服务器发送的时间
//直到使用control+c 快捷键中断它
//Unix系统shell上面回显为^C
//如果系统上没有安装nc或netcat
//可以使用telnet或者一个使用net.Dial实现的Go版的netcat来连接tcp服务器
//
//
//go handleConn() 多个客户端/
