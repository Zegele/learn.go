package main

import (
	"encoding/base64"
	"io/ioutil"
	"net/http"
)

func main() {
	http.ListenAndServe(":8088", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		//http.HandlerFunc 注册一些对http请求的方法
		if request.Body == nil {
			writer.Write([]byte("no body"))
			return
		}
		data, _ := ioutil.ReadAll(request.Body)
		defer request.Body.Close()
		encoded := base64.StdEncoding.EncodeToString(data) // 把data转成base64的编码

		writer.Write(append(data, []byte(encoded)...))
	}))
}
