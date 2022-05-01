package main

import (
	"fmt"
	"net/http"
)

func main() {

	m := http.NewServeMux() // 使用serveMux实现路由分发  分发不同的请求，实现不同的返回， 没有注册路由，默认是404
	m.Handle("/hello", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		//是m.Handle 不是m.Handler
		writer.Write([]byte(`hello`))
	}))
	m.Handle("/ranks", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte(`ranks`))
	}))
	m.Handle("/history/xiao", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte(`xiao的历史`))
	}))

	//....
	m.Handle("/history/jesse", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte(`jesse的健身馆`))
	}))
	// http://localhost:8088/history/jesse
	// jesse的健身馆

	m.Handle("/history", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		qp := request.URL.Query()
		name := qp.Get("name")
		//switch request.Method {
		//case "GET": ///...
		//case "P0ST": ///...
		//}
		writer.WriteHeader(http.StatusOK) //点进去可以看到很多状态码 这里本来默认会加个ok
		writer.Write([]byte(fmt.Sprintf(`%s: %s的健身馆`, request.Method, name)))
	}))
	//浏览器输入：
	// http://localhost:8088/history?name=nine
	// nine的健身馆

	http.ListenAndServe(":8088", m) // 不是8088， 是 :8088 有冒号
}

// 面向接口，但实现接口的是函数。不是通常的结构体。
// 我现在只需要用个func（函数），但是函数不能传，只能传接口怎么办。就可以定义一个function的类型，去实现一个接口，这样就可以相当于传递了函数。
//type Handler interface {
//	ServeHTTP(ResponseWriter, *Request)
//}
//
//type HandlerFunc func(ResponseWriter, *Request) //定义了一个函数类型
//
//说明// ServeHTTP calls f(w, r).
//func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
//	f(w, r)
//}
