// go语言获取IP地址和域名解析
// www.kancloud.cn/imdszxs/golang/1509690
package main

import (
	"fmt"
	"net"
	"os"
)

//主机地址是网络通信最重要的数据之一，
//net包中定义了三种类型的主机地址数据类型：IP，IPMask，IPAddr
//他们分别用来存储协议相关的网络地址

//IP地址类型
//在net包中，IP地址类型被定义成一个byte型切片
//即若干个8位组，格式如下：
//type IP []byte
//在net包中，有几个函数可以将IP地址类型作为函数的返回类型
//比如ParseIP()函数
//该函数原型定义如下：
//func ParseIP(s string)IP
//ParseIP()函数的主要作用是分析IP地址的合法性，
//如果是一个合法的IP地址
//ParseIP()函数将返回一个IP地址对象
//如果是一个非法IP地址，ParseIP()函数将返回nil

//还可以使用IP对象的String()方法将IP地址转换成字符串格式，
//String()方法的原型定义如下：
//func (ip IP)String()string
//如果是IPv4地址，String()方法将返回一个点分十进制格式的IP地址
//如“192.168.0.1”
//如果是IPv6地址，String()方法将返回使用“:”分割的地址形式：如：2000:0:0:0:0:0:0:1
//另外注意一个特例，对于地址“0:0:0:0:0:0:0:1”的返回结果是省略格式“::1”
//示例1：IP地址类型：

/*
func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s ip.addr\n", os.Args[0])
		os.Exit(1)
	}
	addr := os.Args[1]
	ip := net.ParseIP(addr)
	if ip == nil {
		fmt.Println("Invalid address")
	} else {
		fmt.Println("The address is", ip.String())
		os.Exit(0)
	}
}

*/

//从键盘输入：./main.exe 192.168.0.1
//输出结果为：The address is 192.168.0.1
//从键盘输入：192.168.0.256  //最多255
//输出结果为：Inval id address
//从键盘输入：0:0:0:0:0:0:0:1
//输出结果为：::1 //简写这样

//IPMack地址类型
//在Go语言中，为了方便子网掩码操作与计算
//net包中还提供了IPMask地址类型
//在前面讲过，子网掩码地址其实就是一个特殊的IP地址
//所以IPMask类型也是一个byte型数组
//格式如下
//type IPMask []byte
//函数IPv4Mask()可以通过一个32位IPv4地址生成子网掩码地址
//调用成功后返回一个4字节16进制子网掩码地址
//IPv4Mask()函数原型定义如下：
//func IPv4Mask(a, b, c, d byte)IPMask
//另外，还可以使用主机地址对象的DefaultMask()方法获取主机默认子网掩码地址
//DefaultMask()方法原型定义如下：
//func(ip IP)DefaultMask()IPMask
//要注意的是，只有IPv4地址才有默认子网掩码
//如果不是IPv4地址 ,DefaultMask()方法将返回nil
//不管是通过调用IPv4Mask()函数
//还是执行DefaultMask()方法
//获取的子网掩码地址都是十六进制格式的
//例如：子网掩码地址“255.255.255.0”的十六进制格式是“ffffffOO”
//主机地址对象还有一个Mask()方法
//执行Mask()后，会返回IP地址与子网掩码地址相“与”的结果
//这个结果即是主机所处的网络的“网络地址”
//Mask()方法原型定义如下：
//func(ip IP)Mask(mask IPMask)IP
//还可以通过子网掩码对象的Size()方法获取掩码位数（ones）
//和掩码总长度（bits）
//如果是一个非标准的子网掩码地址，
//则Size()方法将返回“0,0”
//Size()方法的原型定义如下：
//func (m IPMask)Size()(ones, bits int)

//示例2:子网掩码地址
/*
//子网掩码地址
func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s ip.addr\n", os.Args[0])
		os.Exit(1)
	}
	dotaddr := os.Args[1]
	addr := net.ParseIP(dotaddr)
	if addr == nil {
		fmt.Println("Invalid address")
	}
	mask := addr.DefaultMask()//获取主机默认子网掩码地址
	fmt.Println("Subnet mask is: ", mask, mask.String())
	network := addr.Mask(mask)//IP与子网掩码的‘与’运算，获得主机地址
	fmt.Println("Network address is :", network.String())
	ones, bits := mask.Size()//获得掩码位数，和掩码总长度
	fmt.Println("Mask bits: ", ones, "Total bits: ", bits)
	os.Exit(0)
}

*/

//./main 192.168.0.1
//Subnet mask is:  ffffff00
//Network address is : 192.168.0.0
//Mask bits:  24 Total bits:  32

// 域名解析
// 在net包中，许多函数或方法调用后返回的是一个指向IPAddr结构体的指针
// 结构体IPAddr内只定义了一个IP类型的字段，格式如下：
//
//	type IPAddr struct{
//		IP IP
//	}
//
// IPAddr结构体的主要作用是用于域名解析服务（DNS）
// 例如：函数ResoloveIPAddr()可以通过主机名解析主机网络地址
// ResoloveIPAddr()函数原型定义如下：
// func ResolveIPAddr(net, addr string)(*IPAddr, error)
// 在调用ResoloveIPAddr()函数时，参数net表示网络类型
// 可以是“ip”，“ip4”或“ip6”
// 参数addr可以是IP地址或域名
// 如果是IPv6地址必须使用“[]”括起来
// ResolveIPAddr()函数调用成功后返回指向IPAddr结构体的指针
// 调用失败返回错误类型error
// 示例3：DNS域名解析
// DNS域名解析
func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s hostname\n", os.Args[0])
		fmt.Println("Usage: ", os.Args[0], "hostname")
		os.Exit(1)
	}
	name := os.Args[1]
	addr, err := net.ResolveIPAddr("ip", name)
	if err != nil {
		fmt.Println("Resolvtion error", err.Error())
		os.Exit(1)
	}
	fmt.Println("Resolved address is: ", addr.String())
	os.Exit(0)
}

//./main c.biancheng.net
//Resolved address is:  118.180.56.187
