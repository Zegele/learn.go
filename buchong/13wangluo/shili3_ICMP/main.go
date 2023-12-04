// go语言ICMP协议：向主机发送消息
// www.kancloud.cn/imdszxs/golang/1509674
package main

import (
	"fmt"
	"net"
	"os"
)

//ICMP是用来对网络状态进行反馈的协议，可以用来侦测网络状态检测网络错误
//ICMP协议（Internet Control Message Protocol 因特网控制报文协议）简介
//是IPv4协议族中的一个子协议，用于IP主机，路由器之间传递控制消息
//控制消息是网络是否通畅、主机是否可达，路由是否可用等网络本身的消息
//这些控制消息虽然不传输用户数据，但是对于用户数据的传递起着重要的作用
//ICMP协议是一种面向无连接的协议，用户传输出错误报告控制信息，
//是一个非常重要的协议，对于网络安全具有机器重要的意义
//ICMP数据网络层协议
//主要用于在主机与路由器之间传递控制信息，包括报告错误，交换受限控制和状态信息等
//当遇到IP数据无法访问目标、IP路由器无法按当前的传输速率转发数据包等情况时，会自动发送ICMP消息
//ICMP是TCP/IP模型中网络层的重要成员。
//与IP协议，ARP协议，RARP协议，IGMP协议共同构成TCP/IP模型中的网络层
//ping和tracert是两个常用网络管理命令
//ping用来测试网络可达性
//tracert用来显示达目的主机的路径
//ping和tracert都利用ICMP协议来实现网络功能
//他们是把网络协议应用到日常网络管理的典型示例
//从计数角度，ICMP就是一个”错误侦测与回报机制“
//目的就是让我们能够检测网络的连线状况 ，也能确保连线的准确性
//当路由器在处理一个数据包的过程中发送了意外
//可以通过ICMP向数据包的源报告有关事件
//其功能主要有：侦测远端主机是否存在，建立及维护路由资料，
//重导资料传送路径（ICMP重定向），资料流量控制
//icmp在沟通之中，主要是透过不同的类别（Type）与代码(Code)让机器来识别不同的连线状况
//ICMP协议大致可以分为两类，一种是查询报文，一种是差错报文。
//其中查询报文有以下几种用途：
//1. ping 查询
//2. 子网掩码查询（用于无盘工作站在初始化自身的时候初始化子网掩码）
//时间戳查询（可以用来同步时间）
//而差错报文则产生在数据传送发生错误的时候

//ICMP消息类型
//icmp报告无法传送数据报的错误，且无法帮助对这些错误进行疑难解答
//例如：IPv4不能将数据报传送到目标主机，路由器或目标主机上的ICMP会向主机发送一条”无法到达目标“消息
//ICMP消息类型   用途
//回显请求：		ping工具通过发送icmp回显消息检查特定节点的ipv4连接以排查网络问题，类型值为0
//回显应答：		节点发送回显答复消息响应icmp回显信息，类型值为8
//重定向：		路由器发送”重定向“消息，告诉发送主机到目标IPv4地址更好的路由，类型值为5
//源抑制：		路由器发送”源结束“消息，告诉发送主机他们的IPv4数据报被丢弃，因为路由器上发生了拥塞，于是发送主机将以较低的频度发送数据报，类型值为4
//超时：			这个消息有两种用途。当超过IP生存期时向发送系统发出错误信息；如果分段的IP数据报没有在某种期限内重新组合，这个消息将通知发送系统，类型值为11
//无法到达目标	路由器和目标主机发送”无法到达目标“消息，通知发送主机他们的数据无法传送，类型值为3
//
//其中无法到达目标消息中可以细分以下几项 ：
//无法到达目标消息	说明
//不能访问主机		路由器找不到目标的IPv4地址的路由时发送”不能访问主机“消息
//无法访问协议		目标ipv4节点无法将ipv4报头中的”协议“字段与当前使用的ipv4客户端协议相匹配时会发送”无法访问协议“消息
//无法访问端口		IPv4节点在UDP报头中的”目标端口“字段与使用该UDP端口的应用程序相匹配时发送”无法访问端口“消息
//需要分段设置了DF	当必须分段但发送节点在ipv4报头中设置了”不分段（DF）“标志时，ipv4路由器会发送”需要分段但设置了DF“消息
//ICMP协议只是试图报告错误，并对特定的情况提供反馈，但最终并没有使IPv4称为一个可靠的协议
//icmp消息是以未确认的Ipv4数据报传送的，他们自己也不可靠

//ICMP的报文格式
//icmp报文包含在IP数据报中，IP报头在ICMP报文的最前面。
//一个ICMP报文包括IP报头（至少20字节），icmp报头（至少8字节）和ICMP报文（属于icmp报文的数据部分）
//当ip报头中的协议字段值为1时，就说明这是一个icmp报文

//常见的icmp报文
//1.响应请求
//日常进行的Ping操作中就包括了响应请求（类型字段值为8）和应答（类型字段值为0）icmp报文
//一台主机向一个节点发送一个类型字段值为8的icmp报文，如果途中没有异常（如果没有被路由丢弃，目标不回应icmp或者传输失败）
//则目标返回类型字段值为0的icmp报文，说明这台主机存在
//2. 目标不可达，源抑制和超时报文
//目标不可达报文（类型值为3）在路由器或者主机不能传递数据时使用
//例如我们要连接对方一个不存在的系统端口（端口号小于1024）时，将返回类型字段值3，代码字段值为3的icmp报文
//常见的不可达类型还有网络不可达（代码字段值为0），主机不可达（代码字段值为1），协议不可达（代码字段值为2）等等
//源抑制报文（类型字段值为4，代码字段值为0），则充当一个控制流量的角色，
//通知主机减少数据报流量。
//由于icmp没有回复传输的报文，所以只要停止该报文，主机就会恢复传输速率
//最后，无连接方式网络的问题就是数据报会丢失
//或者长时间在网络游荡而找不到目标
//或者拥塞导致主机在规定的时间内无法重组数据报分段，
//这时就要触发icmp超时报文的产生
//超时报文（类型字段值为11）的代码域有两种取值，
//代码字段值为0表示传输超时；代码字段值为1表示分段重组超时。
//3.时间戳请求
//时间戳请求报文（类型值字段13）和时间戳应答报文（类型值字段14）用于测试两台主机之间数据报来回一次的传输时间
//传输时，主机填充原始时间戳，接受方收到请求后填充接受时间戳后以类型字段14的报文格式返回，
//发送方计算这个时间差。有些系统不响应这种报文。

//icmp的应用
//Ping
//ping可以说是icmp的最著名的应用，当我们某一个网站上不去的时候，通常会ping一下这个网站
//ping会回显出一些有用的信息，一般的信息如下：
//Reply from 10.4.24.1: bytes=32 time<1ms TTL=255
//Reply from 10.4.24.1: bytes=32 time<1ms TTL=255
//Reply from 10.4.24.1: bytes=32 time<1ms TTL=255
//Reply from 10.4.24.1: bytes=32 time<1ms TTL=255
//
//Ping statistics for 10.4.24.1:
//    Packets: Sent = 4, Received = 4, Lost = 0 (0% loss),
//Approximate round trip times in milli-seconds:
//    Minimum = 0ms, Maximum = 0ms, Average = 0ms
//ping利用icmp协议包来侦测另一个主机是否可达，
//原理是用类型码为0的icmp发请求，收到请求的主机则用类型码为8的icmp回应
//ping程序来计算间隔时间，并计算有多少个包被送达
//用户就可以判断网络大致的情况，
//我们可以看到，ping给出来了传送的时间和TTL的数据
//ping还给我们一个看主机到目的主机的路由的机会
//这是因为icmp的ping请求数据报在每经过一个路由器的时候，路由器都会把自己的ip放到该数据报中
//而目的主机则会把这个ip列表复制到回应icmp数据包中发回给主机/

func checkSum(msg []byte) uint16 {
	sum := 0
	len := len(msg)
	for i := 0; i < len-1; i += 2 {
		sum += int(msg[i])*256 + int(msg[i+1])
	}
	if len%2 == 1 {
		sum += int(msg[len-1]) * 256 // notice here, why *256
	}
	sum = (sum >> 16) + (sum & 0xffff)
	sum += (sum >> 16)
	var answer uint16 = uint16(^sum)
	return answer
}

func checkError(err error) {
	if err != nil {
		fmt.Fprint(os.Stderr, "Fatal error:", err.Error())
		os.Exit(1)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:", os.Args[0], "host")
		os.Exit(1)

	}
	service := os.Args[1]
	conn, err := net.Dial("ip4:icmp", service)
	checkError(err)
	var msg [512]byte
	msg[0] = 8
	msg[1] = 0
	msg[2] = 0
	msg[3] = 0
	msg[4] = 0
	msg[5] = 13
	msg[6] = 0
	msg[7] = 37
	msg[8] = 99
	len := 9
	check := checkSum(msg[0:len])
	msg[2] = byte(check >> 8)
	msg[3] = byte(check & 0xff)
	fmt.Println(msg[0:len])
	_, err = conn.Write(msg[0:len])
	checkError(err)
	_, err = conn.Read(msg[0:])
	checkError(err)
	fmt.Println(msg[0 : 20+len])
	fmt.Println("Got response")
	if msg[20+5] == 13 {
		fmt.Println("Identifier matches")
	}
	if msg[20+7] == 37 {
		fmt.Println("Sequence matches")
	}
	if msg[20+8] == 99 {
		fmt.Println("Custom data matches")
	}
	os.Exit(0)
}

// 运行
//go run main.go c.biancheng.net
//[8 0 148 205 0 13 0 37 99]
//[69 0 0 29 6 249 0 0 56 1 171 94 222 75 48 146 192 168 1 3 0 0
// 156 205 0 13 0 37 99]
//Got response
//Identifier matches
//Sequence matches
//Custom data matches
//但是无论如何，ip头能记录的路由列表是非常有限的
//如果要观察路由，我们还是需要使用更好的工具
//就是要讲到的Traceroute（windows下的名字叫tracert）
//
//Traceroute
//记录下所有经过的路由器
//：
