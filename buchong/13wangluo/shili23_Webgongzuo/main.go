// Go语言是如何使得web工作的？
// www.kancloud.cn/imdszxs/golang/1509694
package main

//前面已经介绍了如何通过Go语言搭建一个Web服务
//我们可以看到简单应用一个net/http包就方便的搭建起来了
//那么GO语言在底层到底是怎么做的呢？

//Web工作方式的几个概念
//以下均是服务器端的几个概念：
//1. Request： 用户请求的信息，用来解析用户的请求信息
//包括post，get，cookie，url等信息
//2. Reponse: 服务器需要反馈给用户的信息
//3. Conn：用户的每次请求连接
//4. Handler: 处理请求和生成返回信息的处理逻辑

//分析http包运行机制
//下图是Go实现Web服务的工作模式流程图
//图：http包执行流程
//1. 创建Listen Socket， 监听指定的端口，等待客户端请求到来
//2. Listen Socket接受客户端的请求，得到Client Socket
//接下来通过Client Socket与客户端通信
//3. 处理客户端的请求，首先从Client Socket读取HTTP请求的协议头
//如果是POST方法，还可能要读取客户端提交的数据
//然后交给响应的handler处理请求
//handler处理完毕准备好客户端需要的数据
//通过Client Socket写给客户端

//这整个的过程里面我们只要了解清楚下面3个问题，也就知道Go是如何让Web运行起来了
//1. 如何监听端口
//2. 如何接收客户端请求
//3. 如何分配handler

//前面小节的代码里面我们可以看到，go是通过一个函数ListenAndServe来处理这些事情的
//这个底层其实这样处理的：
//1.初始化一个server对象
//2. 然后调用了net.Listen("tcp", addr),也就是底层用tcp协议搭建了一个服务
//3. 然后监控我们设置的端口
//下面代码来自go的http包的源码
//通过下面的源码我们可以看到整个的http处理过程：
//func (srv *Server) Serve(l net.Listener)error{
//	defer l.Close()
//	var tempDelay time.Duration // how long to sleep on accept failure
//	for{
//		rw, e := l.Accept()
//		if e != nil{
//			if ne, ok := e.(net.Error); ok && ne.Temporary(){
//				if tempDelay == 0{
//					tempDalay = 5 * time.Millisecond
//				}else{
//						tempDelay *=2
//				}
//				if max := 1 * time.Second; tempDelay > max {
//					tempDelay = max
//				}
//				log.Printf("http: Accept error: %v; retrying in %v", e, tempDelay)
//				time.Sleep(tempDelay)
//				continue
//			}
//			return e
//		}
//		tempDelay = 0
//		c, err := srv.newConn(rw)
//		if err != nil{
//			continue
//		}
//		go c.serve()
//	}
//}

//监控之后如何接收客户端的请求呢？
//上面代码执行监控端口之后，调用了 srv.Serve(net.Listener)函数
//这个函数就是处理接收客户端的请求信息
//这个函数里面起了一个for循环
//首先通过Listener接收请求，
//其次创建一个Conn
//最后单独开了一个goroutine
//把这个请求的数据当做参数扔给这个conn去服务 go c.serve()
//这个就是高并发体现了
//用户的每一个请求都是在一个新的goroutine去服务，相互不影响

//那么如何具体分配到相应的函数来处理请求呢？
//conn首先会解析request:c.readRequest(),
//然后获取相应的handler:handler := c.server.Handler
//也就是我们刚才调用函数ListenAndServe时候的第二个参数
//我们前面例子传递的是nil，也就是为空
//那么默认获取handler= DefaultServeMux

//那么这个变量用来做什么？
//这个变量就是一个路由器，
//它用来匹配url跳转到其相应的handle函数
//那么这个我们有设置过吗？
//有，我们调用的代码里面第一句不是调用了http.HandleFunc("/",sayhelloName)嘛
//这个作用就是注册了请求 '/' 的路由规则
//当请求url为“/”,
//路由就会转到函数sayhelloName,DefaultServeMux会调用ServeHTTP方法
//这个方法内部其实就是调用sayhelloName本身
//最后通过写入response的信息反馈到客户端
//详细流程如下图所示：
//图：一个http连接处理流程
//至此我们的三个问题已经全部得到了解答
//现在对于Go如何让Web跑起来的已经基本了解了