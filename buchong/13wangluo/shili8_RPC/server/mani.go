// go语言RPC协议：远程过程调用
// www.kancloud.cn/imdszxs/golang/1509679
package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
)

//RPC协议构建于TCP,UDP或者是HTTP之上，允许开发人员直接调用另一条计算机上的程序，
//而开发人员无需额外地为这个调用过程编写网络通信相关代码
//使得开发网络分布式类型的应用程序更加容易
//go语言的标准库提供了RPC框架和不同的RPC实现。

//1. 什么RPC
//远程过程调用（Remote Procedure Call，简称RPC）是一个计算机通信协议
//该协议允许运行一台计算机的程序调用另一台计算机的子程序
//而开发人员无需额外地为这个交互作用编程
//如果设计的软件采用面向对象编程
//那么远程过程调用亦可称作远程调用或远程方法调用
//通俗来讲就是，RPC允许跨机器，跨语言调用计算机程序
//例如：用go语言写了一个获取用户信息的方法getUserInfo，并把go语言程序部署在阿里云服务器上，
//另外我们还有一个部署在腾讯云上的php项目，需要调用go语言的getUserInfo方法获取用户信息
//php跨机器调用go方法的过程就是RPC调用
//
//调用流程：
//1. 调用客户端句柄，执行传送参数
//2. 调用本地系统内核发送网络消息
//3. 消息传送到远程主机
//4. 服务器句柄得到消息并取得参数
//5. 执行远程过程
//6. 执行的过程将结果返回服务器句柄
//7. 服务器句柄返回结果，调用远程系统内核
//8. 消息传回本地主机
//9. 客户句柄由内核接收消息
//10. 客户接收句柄返回的数据

//2. go语言如何实现RPC
//go实现RPCf非常简单，有封装好的官方包，和一些第三方包提供支持
//go语言中RPC可以利用tcp或http来传递数据
//可以对要传递的数据使用多种类型的编码方式
//go语言的net/rpc包使用encoding/gob进行编解码，
//支持tcp或http数据传输方式
//由于其他语言不支持gob编码方式，所以使用net/rpc包实现的RPC方法没办法进行跨语言调用
//此外，go语言还提供了net/rpc/jsonrpc包实现RPC方法
//JSON RPC采用JSON进行数据编解码，因而支持跨语言调用，
//但目前的jsonrpc包是基于tcp协议实现的，暂时不支持使用http进行数据传递
//除了go语言官方提供的rpc包，还有许多第三方包为在go语言中实现RPC提供支持
//大部分第三方rpc包的实现都是使用protobuf进行数据编解码
//根据protobuf声明文档自动生成rpc方法定义与服务注册代码
//所以在go语言中可以很方便的进行rpc服务调用

//3. net/rpc包
//rpc 包提供了通过网络或其他i/o连接对一个对象的导出方法的访问
//服务端注册一个对象，使他作为一个服务被暴露
//服务的名字是该对象的类型名
//注册之后，对象的导出方法就可以被远程访问
//服务端可以注册多个不同类型的对象（服务）
//但注册具有相同类型的多个对象是错误的。
//只有满足如下标准的方法才能用于远程访问，其余方法会被忽略
//1. 方法是可导出的
//2. 方法有两个参数，都是导出类型或内建类型
//3. 方法的第二个参数是指针类型
//4. 方法只有一个error接口类型的返回值
//下面演示go语言net/rpc包实现RPC方法，使用http作为RPC的载体，
//通过net/http包监听客户端连接请求

// 算数运算结构体
type Arith struct {
}

// 算数运算请求结构体
type ArithRequest struct {
	A int
	B int
}

// 算数运算响应结构体
type ArithResponse struct {
	Pro int // 乘积
	Quo int // 商
	Rem int // 余数
}

// 乘法运算方式
func (this *Arith) Multiply(req ArithRequest, res *ArithResponse) error {
	res.Pro = req.A * req.B
	return nil
}

// 除法运算方法
func (this *Arith) Divide(req ArithRequest, res *ArithResponse) error {
	if req.B == 0 {
		return errors.New("除以零")
	}
	res.Quo = req.A / req.B
	res.Rem = req.A % req.B
	return nil
}

func main() {
	rpc.Register(new(Arith)) // 注册rpc服务
	rpc.HandleHTTP()

	// 采用http协议作为rpc载体
	lis, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatalln("致命错误：", err)
	}

	fmt.Fprintf(os.Stdout, "%s", "开始连接")
	http.Serve(lis, nil)
}

//服务端程序运行后，将会监听本地8080端口
//客户端转看client-main.go，

// 4. net/rpc/jsonrpc库 (详见server2-main.go)
//上面的示例只试用go语言远程调用，因为使用的是gob
