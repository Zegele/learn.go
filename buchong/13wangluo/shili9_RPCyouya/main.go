// 如何设计优雅的RPC接口
// www.kancloud.cn/imdszxs/golang/1509680
package main

// RPC是一种方便的网络通信编程模型，由于和编程语言的高度结合
//大大减少了处理网络数据的复杂度，让代码可读性也有可观的提高
//但是RPC本身的构成却比较复杂
//由于受到编程语言，网络模型，使用习惯的约束，有大量的妥协和取舍之处

//认识RPC（远程调用）
//指的是用一行简单的代码，通过网络调用另外一个计算机商的某段程序。 如
//RMI（Remote Method Invoke）：调用远程的方法，
//“方法”一般是附属于某个对象上的
//所以通常RMI指对在远程的计算机上的某个对象，进行其方法函数的调用
//RPC（Remote Procedure Call） : 远程过程调用，指对网络上另外一个计算机上的，某段特定的函数代码的调用

//远程调用本身是网络通信的一种概念，它的特点是把网络通信封装成一个类似函数的调用
//网络通信在远程调用外，一般还有其他的几种概念：数据包处理，消息队列，流过滤，资源拉取等，他们的差异如下：
//方案		编程方式											信息封装						传输模型							典型应用
//远程调用	调用函数，输入参数，获得返回值。						使用编程语言的变量,类型，函数	发出请求，获得响应					JavaRMI
//数据包处理	调用Send()/Recv(),使用字节码数据，编解码，处理内容	把通信内容构造成二进制的协议包	发送，接收						UDP编程
//消息队列	调用Put()/Get()，使用“包”对象，处理其包含的内容		消息被封装成语言可用的对象或结构	对某队列存入一个消息或取出一个消息	ActiveMQ
//流过滤		读取一个流或写出一个流，对流中的单元包即刻处理			单元长度很小的统一数据结构		连接，发送/接收，处理				网络视频
//资源拉取	输入一个资源ID，获得资源内容						请求或响应都包含：头部和正文	请求后等待响应					WWW

//远程调用的优势
//1. 屏蔽了网络层
//因此在传输协议和编码协议上，我们可以选择不同的方案
//比如WebService方案就是用的HTTP传输协议+SOAP编码协议
//而REST的方案往往使用HTTP+JSON协议
//Facebook的Thrift可以定制任何不同的传输协议和编码协议
//可以用TCP+Google Protocol Buffer 也可以用UDP+JSON等
//由于屏蔽了网络层，可以根据实际需要来独立的优化网络部分
//而无需涉及业务逻辑的处理代码，这对于需要在各种网络环境下运行的程序来说，非常有价值
//
//2. 函数映射协议
//可以直接用编程语言来书写数据结构和函数定义，取代编写大量的编码协议格式和分包处理逻辑
//对于那些业务逻辑非常复杂的系统
//比如网络游戏，可以节省大量定义消息格式的时间
//函数调用模型非常容易学习，不需要歇息通信协议和流程，让经验较浅的程序员也能很容易的开始使用网络编程

//远程调用的缺点
//1. 增加了性能消耗
//由于把网络通信包装成“函数”，需要大量额外的处理，比如需要预生产代码，
//或者使用反射机制，这些都是额外消耗cpu和内存的操作
//而且为了表达复杂的数据类型
//比如变长的类型string/map/list,这些都要数据包中增加更多的描述性信息，则会占用更多的网络包长度
//2. 不必要的复杂化
//如果是为了某些特定的业务需求，比如传送一个固定的文件，那么应用HTTP/FTP协议模型
//如果为了做监控或者IM软件，用简单的消息编码手法会更快速高效
//如果是为了做代理服务器，用流式的处理会很简单
//另外，如果要做数据广播，那么消息队列会很容易做到
//而远程调用这几乎无法完成
//远程调用最适合是业务需求变多或网络环境多变的场景

//RPC结构拆解
//RPC服务端通过 RpcServer去导出（export）远程接口方法，
//而客户端通过RpcClient去引入（import）远程接口方法
//客户端像调用本地方法一样去调用远程接口方法
//RPC框架提供接口的代理实现，实际的调用将委托给代理RpcProxy
//代理封装调用信息并将调用转交给RpcInvoker去实际执行
//在客户端的RpcInvoker通过连接器RpcConnector去维持与服务端的通道RpcChannel,
//并使用RpcProtocol执行协议编码（encode）并将编码后的请求消息通过通道发送给服务端
//RPC服务端接收器RpcAcceptor接收客户端的调用请求，同样使用RpcProtocol执行协议解码（decode）
//解码后的调用信息传递给RpcProcessor去控制处理调用过程
//最后再委托调用给RpcInvoker去实际执行并返回调用结果
//RPC各个组件的职责如下所示：
//1. RpcServer: 负责导出（export）远程接口
//2. RpcClient: 负责导入（import）远程接口的代理实现
//3. RpcProxy: 远程接口的代理实现
//4. RpcInvoker：
//5. 客户方实现：负责编码调用信息和发送调用请求到服务方并等待调用结果返回
//6. 服务方实现：负责调用服务端接口的具体实现并返回调用结果
//7. RpcProtocol: 负责协议编解码
//8. RpcConnector: 负责维持客户方和服务方的连接通道和发送数据到服务方
//9. RpcAccptor: 负责接收客户方请求并返回请求结果
//10. RpcProcessor：负责在服务方控制调用过程，包括管理调用线程池，超时时间等
//11. RpcChannel： 数据传输通道

//RPC接口设计
//go语言的net/rpc很灵活，他在数据传输前后实现了编码解码的接口定义
//这意味着，开发者可以自定义数据的传输方式以及RPC服务端和客户端之间的交互行为
//RPC提供的编码解码器接口如下：
//type ClientCodec interface{
//WriteRequest(*Request, interface{})error
//ReadResponseHeader(*Response)error
//ReadResponseBody(interface{})error
//Close()error
//}

//type ServerCodec interface{
//ReadRequestHeader(*Request)error
//ReadRequestBody(interface{})error
//WriteResponse(*Response, interface{})error
//Close()error
//}

//接口ClientCodec定义了RPC客户端如何在一个RPC会话中发送请求和读取响应
//客户端程序通过WriteRequest()方法将一个请求写入到RPC连接中
//并通过ReadResponseHeader()和ReadResponseBody()读取服务端的响应信息
//当整个过程执行完毕后，再通过Close()方法关闭该连接
//接口ServerCodec定义了RPC服务端如何在一个RPC会话中接收请求并发送响应
//服务端程序通过ReadRequestHeader()和ReadRequestBody()方法从一个RPC连接中读取请求信息
//然后再通过WriteResponse()方法向该连接中的RPC客户端发送响应
//当完成该过程后，通过Close()方法关闭连接
//通过实现上述接口，我们可以自定义数据传输前后的编码解码方式
//而不仅仅局限于Gob，
//同样，可以自定义RPC服务端和客户端的交互行为
//实际上，Go标准库提供的net/rpc/json包就是一套实现了rpc.ClientCodec和rpc.ServerCodec接口的JSON-RPC模块/

func main() {

}
