package main

import (
	"encoding/base64"
	"io/ioutil"
	"net/http"
)

func main() {
	http.ListenAndServe(":8088", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		//http.HandlerFunc 注册一些对http请求的方法
		if request.Body == nil { // request.Body相当于payload
			writer.Write([]byte("no body"))
			return
		}
		data, _ := ioutil.ReadAll(request.Body)
		defer request.Body.Close()
		encoded := base64.StdEncoding.EncodeToString(data) // 把data转成base64的编码

		writer.Write(append(data, []byte(encoded)...))
	}))
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
