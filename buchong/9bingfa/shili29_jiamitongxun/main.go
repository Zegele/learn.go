package main

import (
	"fmt"
	"net/http"
)

// c.biancheng.net/view/vip_7357.html 看不了
// www.zixuephp.com/manual/body/id/221/pid/1%3E
// 为什么运行不出想要的结果？
const SERVER_PORT = 8080
const SERVER_DOMAIN = "localhost"
const RESPONSE_TEMPLATE = "c语言中文网"

func rootHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Content-Length", fmt.Sprint(len(RESPONSE_TEMPLATE)))
	w.Write([]byte(RESPONSE_TEMPLATE))
}

//func rootHandler(w http.ResponseWriter, req *http.Request) {
//	w.Header().Set("Content-Type", "text/plain")
//	w.Header().Set("Content-Length", fmt.Sprint(len(RESPONSE_TEMPLATE)))
//	w.Write([]byte(RESPONSE_TEMPLATE))
//}

func main() {
	http.HandleFunc(fmt.Sprintf("%s:%d/", SERVER_DOMAIN, SERVER_PORT), rootHandler)
	http.ListenAndServeTLS(fmt.Sprintf("%d", SERVER_PORT), "cert.pem", "key.pem", nil)
	// 注意： 上面的冒号   上面"rui.crt", "rui.key"对么？ 不应是cert.pem, key.pem 么？存放私钥、公钥的文件？
}

/*
//http://code.js-code.com/android/222221.html
func main() {
http.HandleFunc("/", rootHandler)
http.ListenAndServeTLS(":8080", "cert.pem", "key.pem", nil)
//使用http.ListenAndServeTLS实现一个支持https的Web服务器。
// 运行上面的程序需要用到cert.pem和 key.pem 两个文件，
//可以使用crypto/tls包的generate_cert.go文件来生成。  怎么生成这些文件，参考该目录下其他文件。
//运行成功后，我们可以在浏览器中通过localhost：8080 查看访问效果。
}

*/
