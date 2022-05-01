package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.ListenAndServe(":8088", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		//Handler是什么东西
		qp := request.URL.Query() // URL请求参数
		// 服务器从request中读取数据
		data, _ := json.Marshal(qp)
		writer.Write([]byte("hello,您好：" + string(data)))
		// 服务器从writer中写出数据

		// 浏览器内输入 http://localhost:8088/?name=xiao,sex=男
		//hello,您好：{"name":["xiao,sex=男"]} // 当作一个整体
		//http://localhost:8088/?name=xiao&sex=男
		//hello,您好：{"name":["xiao"],"sex":["男"]}// 输入了两个参数
		// 如果输入了xiao的名字中就含有个&符号，结果输出的还是原来的。也就是说会有信息丢失。所以这种方式使用有限。 所以要用post 会把所有数据传过去。
	}))
}

//浏览器中：
//localhost：8088
//
//返回了：
//hello,您好：{}
//
//localhost：8088/?name=xiaoqiang
//返回了：
//hello,您好：{"name":["xiaoqiang"]}
//
//localhost：8088/?name=xiaoqiang,sex=男 //均无空格
//返回了：
//hello,您好：{"name":["xiaoqiang,sex=男"]} // 都当成名字了
//
//localhost:8088/?name=xiaoqiang&sex=男 //均无空格
//返回了：
//hello,您好：{"name":["xiaoqiang"],"sex":["男"]} 这样才对了
//
//但是，name是xiaoqiang&呢？：
//localhost:8088/?name=xiaoqiang&&sex=男 //均无空格
//返回了：
//hello,您好：{"name":["xiaoqiang"],"sex":["男"]} // 名字中的&符号被忽略了。所以存在泄露风险不安全。
