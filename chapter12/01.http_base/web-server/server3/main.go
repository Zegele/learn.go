package main

import (
	"fmt"
	"net/http"
)

func main() {

	m := http.NewServeMux()
	m.Handle("/hello", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		//是m.Handle 不是m.Handler
		writer.Write([]byte(`hello`))
	}))
	m.Handle("/rank", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte(`rank`))
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
	// http://localhost:8088/history?name=nine
	// nine的健身馆

	http.ListenAndServe(":8088", m) // 不是8088， 是 :8088 有冒号
}

/*
Handle that calls f.
type HandlerFunc func(ResponseWriter, *Request) // 定义了一个func的类型

// ServeHTTP calls f(w, r)
func(f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request){// 让func的类型去实现这个接口 （某个接口有ServeHTTP方法）
	f(w, r)
}
*/
