package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.ListenAndServe(":8088", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		qp := request.URL.Query()
		data, _ := json.Marshal(qp)
		writer.Write([]byte("hello,您好：" + string(data)))
		// 浏览器内输入 http://localhost:8088/?name=xiao,sex=男
		//hello,您好：{"name":["xiao,sex=男"]} // 当作一个整体
		//http://localhost:8088/?name=xiao&sex=男
		//hello,您好：{"name":["xiao"],"sex":["男"]}// 输入了两个参数
		// 如果输入了xiao的名字中就含有个&符号，结果输出的还是原来的。也就是说会有信息丢失。所以这种方式使用有限。 所以要用post 会把所有数据传过去。
	}))
}
