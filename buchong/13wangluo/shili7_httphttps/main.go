// go语言服务端处理http，https请求
// www.kancloud.cn/imdszxs/golang/1509678
package main

import (
	"log"
	"net/http"
)

//服务端是如何处理HTTP和HTTPS请求的？
//1. 处理HTTP请求
//使用net/http包提供的http.ListenAndServe()方法
//可以对指定的地址进行监听，开启一个HTTP，服务端该方法的原型如下：
//func ListenAndServer(addr string, handler Handler)error
//该方法用于在指定的TCP网络地址addr进行监听，然后调用服务端处理程序来处理传入的连接请求
//ListenAndServe方法有两个参数，其中第一个参数addr即监听地址，
//第二个参数表示服务端处理程序，通常为空
//第二个参数为空时，意味着服务端调用http.DefaultServeMux进行处理
//而服务端编写的业务逻辑处理程序http.Handle()或http.HandleFunc()默认注入http.DefaultServeMux中，如下：
//http.Handle("/foo", fooHandler)
//http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request){
//	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
//} )
//
//log.Fatal(http.ListenAndServer(":8080", nil))

//如果想更多地控制服务端的行为，可以自定义http.Server,代码如下：
//s := &http.Server{
//	Addr : ":8080",
//	Handle : myHandler,
//	ReadTimeout : 10 * time.Second,
//	WriteTimeout : 10 * time.Second,
//	MaxHeaderBytes : 1 << 20，
//}
//log.Fatal(s.ListenAndServer())

//演示一个简单的服务端，如何处理HTTP请求，如下：
/*
func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "golang\n")
}

func main() {
	http.HandleFunc("/hello", HelloServer)
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe : ", err)
	}
}

*/

//运行后，我们可以使用浏览器访问
//http://localhost:12345/hello 查看运行结果

// 2. 处理HTTPS请求
// net/http包还提供http.ListenAndServerTLS()方法，用于处理HTTPS连接请求
//func ListenAndServerTLS(addr string, certFile string, keyFile string, handler Handler)error
//ListenAndServeTLS函数和ListenAndServe函数的行为基本一致，
//区别在于ListenAndServerTLS函数只处理HTTPS请求
//此外，服务器上必须提供证书文件和对应的私钥文件，比如certFile对应SSL证书文件存放路径，
//keyFile 对应证书私钥文件路径
//如果证书是由权威机构签发的，certFile参数指定的路径必须是存放在服务器上的经由CA认证果的SSL证书
//开启SSL监听服务也很简单，如下：
//http.Handle("/foo", fooHandler)
//http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request){
//fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
//})
//log.Fatal(http.ListenAndServeTLS(":10443","cert.pem","key.pem",nil))

//或者是：
//ss := &http.Server{
//	Addr : ":10443",
//	Handler : myHandler,
//	ReadTimeout : 10 * time.Second,
//	writeTimeout : 10 * time.Second,
//	MaxHeaderBytes : 1 << 20,
//}
//log.Fatal(ss.ListenAndServeTLS("cert.pem", "key.pem"))
//下面演示go如何处理https：

func handler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "text/plain")
	w.Write([]byte("golang"))
}

func main() {
	http.HandleFunc("/", handler)
	log.Printf("监听 1234 端口成功，可以通过 https://127.0.0.1:1234/ 访问")
	err := http.ListenAndServeTLS(":1234", "cert.pem", "key.pem", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// 运行上面的程序需要用到cert.pem和key.pem这两个文件
//可以使用crypto/tls包的generate_cert.go文件来生成cert.pem, key.pem这两个文件。
