package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

var qa = map[string]string{
	"你好":      "你好",
	"你是谁？":    "我是小小",
	"你是男是女？":  "你猜猜看",
	"今天天气怎样？": "今天天气不错！",
	"再见":      "再见",
}

func main() {
	var port string                                 // port（端口）
	flag.StringVar(&port, "port", "8080", "配置启动端口") // 什么意思？？？ 给参数port赋值 默认8080
	// ./server --port=8081 设置端口的命令参数
	flag.Parse() //把端口绑进去 // 就是按上面的样式，可以更换端口，如果不更换，就默认是8080端口

	//	ln, err := net.Listen("tcp", ":8080") //tcp: Transfer Control Protocol 传输控制协议; 只有8080端口
	ln, err := net.Listen("tcp", ":"+port) //tcp: Transfer Control Protocol 传输控制协议; tcp+端口号。

	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("warning:建立连接失败：", err)
			continue
		}
		fmt.Println(conn)

		//talk(conn)//如果只是调用这个函数，就会随主函数只执行一次。
		go talk(conn) //加深理解goroutine 每调用一次，都会有个goroutine
	}
}

func talk(conn net.Conn) {
	defer fmt.Println("结束链接：", conn)
	defer conn.Close()

	for {
		buf := make([]byte, 1024)
		valid, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF { // 如果不加这个，会一直报错EOF，如果加上会睡一秒，然后continue
				time.Sleep(1 * time.Second)
				continue
			}
			log.Println("WARNING：读取数据失败：", err)
			continue
		}
		content := buf[:valid]
		resp, ok := qa[string(content)] //string(content) 把 content 转化成 字符串类型
		if !ok {
			log.Println("没有找到回答，问他说了什么")
			conn.Write([]byte(`我听不懂你在讲什么。`)) //[]byte(`我听不懂你在讲什么。`) 把这句话，转成 byte切片 类型
			continue
		}
		conn.Write([]byte(resp))
		if string(content) == "再见" {
			break
		}
	}
}
