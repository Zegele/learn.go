// http服务端
package main

import (
	"fmt"
	"net/http"
)

//使用go语言中net/http包来编写一个简单的接收http请求的server
// net/http包是对net包的进一步封装，专门用来处理HTTP协议的数据，如下：

func sayHello(w http.ResponseWriter, r *http.Request) { //请求处理函数
	fmt.Fprintln(w, "hello golang")
}

func main() {
	//http.HandleFunc("/", sayHello)
	http.HandleFunc("/a", sayHello) // 注册路由
	//err := http.ListenAndServe("localhost:9090", nil)//效果同下
	err := http.ListenAndServe(":9090", nil) //监听端口启动服务
	if err != nil {
		fmt.Printf("http server failed, err:%v\n", err)
		return
	}
}

//执行后，在浏览器输入： localhost:9090/a
//就会看到 hello golang/

//http.HandleFunc参考：blog.csdn.net/qq_34021712/article/details/109907490
//这个函数就是注册路由，
//第一个参数是路由表达式，也就是请求路径
//第二个参数是一个函数类型，也就是真正处理请求的函数，
//没有其他逻辑，直接调用DefaultServerMux.HandleFunc()处理

///
