// go语言DialTCP() ： 网络通信
// www.kancloud.cn/imdszxs/golang/1509676
package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

//shili4_main.go实现的基于TCP发送的HTTP请求，读取服务器信息并返回HTTP Head 的示例程序也可以使用下面的方式实现
//。

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	fmt.Println("------", tcpAddr) //------ 39.156.66.10:80

	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	fmt.Println("------conn:", conn) //------conn: &{{0xc00014ea00}}

	checkError(err)
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)
	result, err := ioutil.ReadAll(conn)
	checkError(err)
	fmt.Println(string(result))
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

// 与之前使用Dial()的例子相比，有两个不同点：
//1. net.ResolveTCPAddr()用于解析地址和端口号
//2. net.DialTCP()用于建立连接
//提示：这两个函数在Dial()函数中都得到了封装

//net包中还包含了一系列的工具函数
//合理地使用这些函数可以更好地保障程序的质量
//验证IP地址有效性的函数如下：
//func net.ParseIP
//创建子网掩码的函数如下：
//func IPv4Mask(a, b, c, d byte)IPMask
//获取默认子网掩码的函数如下：
//func (ip IP) DefaultMask()IPMask
//根据域名查找IP的函数如下：
//func ResolveIPAddr(net, addr string)(IPAddr, error)
//func LookupHost(name string)(cname string, addrs []string, err error)
//。
