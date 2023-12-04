// socket
// www.cnblogs.com/aaronthon/p/10951929.html
package main

//socket是BSD Unix的进程通信机制，通常也称作“套接字”
//用于描述IP地址和端口，是一个通信链的句柄
//可以理解为TCP/IP网络的API，它定义了许多函数或例程
//可以用它开发tcp/ip网络上的应用程序
//电脑上运行的应用程序通常通过“套接字”向网络发出请求或应答网络请求
//socket把复杂的tcp/ip协议族隐藏再socket后面，对用户来说只需要调用socket规定的相关函数，
//让socket去组织符合指定的协议数据然后进行通信
//
//
//go 语言实现
//一、TCP通信
//tcp服务端（转看1tcp-server-main.go）
//一个tcp服务端可以同时连接很多个客户端，例如世界各地的用户使用自己电脑上的浏览器访问淘宝
//因为go语言中创建多个goroutine实现并发非常方便和高效，所以我们可以每建立一次连接就创建一个goroutine去处理
//tcp服务端程序的处理流程：
//1.监听端口
//2.接收客户端请求建立链接
//3.创建goroutine处理链接

// 客户端（转看1tcp-client-main.go）
//一个tcp客户端进行tcp通信的流程如下：
//1. 建立与服务端的连接
//2. 进行数据收发
//3. 关闭连接

//二、UDP协议（User Datagram Protocol）用户数据报协议
//不需要建立间接就能直接进行数据发送和接收，属于不可靠的，没有时序的通信，但是UDP协议的实时性比较好，通常用于视频直播相关领域
//UDP服务端（转看2pdp-server-main.go）
//UDP客户端（转看2pdp-client-main.go）

//三、HTTP协议（HyperText Transfer Protocol）
//是互联网上应用最为广泛的一种网络传输协议，所有的www文件都必须遵守这个标准
//设计http最初的目的是为了提供一种发布和接收html页面的方法
///
