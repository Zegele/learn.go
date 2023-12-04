// go语言如何搭建网站程序
// www.kancloud.cn/imdszxs/golang/1509682
package main

import (
	"io"
	"log"
	"net/http"
)

// 如何搭建一个简单的网站程序
func helloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, world!")
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}

// go run main.go
// localhost:8080/hello

//我们引入go语言标准库中的net/http包
//主要用于提供web服务，响应并处理客户端（浏览器）的HTTP请求
//同时，使用io包而不是fmt包来输出字符串
//这样源文件编译成可执行文件后，体积要小很多
//运行起来也更省资源

//简单了解go语言http包再上述示例中所作的工作。
//net/http包简介
//在main()中调用了http.HandleFunc()，该方法用于分发请求
//即针对某一路径请求将其映射到指定的业务逻辑处理方法中
//可以将其形象地理解为提供类似URL路由或者URL映射之类的功能
//在hello.go中，http.HandleFunc()方法接收两个参数，
//第一个参数是HTTP请求的目标路径"/hello"
//该参数值可以是字符串，也可以是字符串形式的正则表达式，
//第二个参数指定具体的回调方法，比如helloHandler
//当我们的程序运行起来后，访问http://localhost:8080/hello
//就会调用helloHandler()方法中的业务逻辑程序
//在上述例子中，helloHandler() 方法是http.HandlerFunc类型的实例，
//并传入http.ResponseWriter和http.Request作为其必要的两个参数
//http.ResponseWriter 类型的对象用于包装处理HTTP服务端的响应信息
//我们将字符串”Hello, world!“写入类型为http.ResponseWriter的w实例中，
//即可将该字符串数据发送到HTTP客户端
//第二个参数r *http.Request表示的是此次HTTP请求的一个数据结构体
//即代表一个客户端，不过该实例中我们尚未用到它

//还看到在main()方法中调用了http.ListenAndServer()，该方法用于在示例中监听8080端口
//接收并调用内部程序来处理连接到此端口的请求
//如果端口监听失败，会调用log.Fatal()方法输出异常出错信息
//如你所见，main()方法中的短短两行即开启了一个http服务，
//使用go语言的net/http包
//搭建一个Web 是如此简单
//当然，net/http包的作用远不止这些。
