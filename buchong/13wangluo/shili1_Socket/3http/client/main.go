// http用户端
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
func main() {
	conn, err := net.Dial("tcp", "www.baidu.com:80")
	if err != nil {
		fmt.Println("dial faild, err:", err)
		return
	}
	defer conn.Close()

	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n") //发送请求

	var buf [8192]byte
	// 接收响应
	for {
		n, err := conn.Read(buf[:])
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("get response failed, err:", err)
			break
		}
		fmt.Print(string(buf[:n]))
	}
}
//编译成可执行文件后，执行就能在终端打印xxx.com网站首页内容
//我们的浏览器其实就是一个发送和接收HTTP协议数据的客户端
//我们平时通过浏览器访问网页其实就是从网站的服务器接收http数据
//然后浏览器会按照html，css等规则将网页渲染展示出来
//

*/

//我们还可以直接使用go语言封装好的net/http包
//它提供了http客户端和服务端的实现
//有了net/http包，我们请求xxx.com网站的页面就会比较简单，如下：

func main() {
	resp, err := http.Get("https://www.baidu.com/")
	if err != nil {
		fmt.Println("get failed, err:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("%q\n", body)
}

// 编译可执行文件后，执行就能在终端输出xxx.com网站首页的内容
