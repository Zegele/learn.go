package main

import (
	"fmt"
	"log"
	"net"
)

var qa = map[string]string{
	"你好":      "你好",
	"你是谁？":    "我是小小",
	"你是男是女？":  "你猜猜看",
	"今天天气怎样？": "今天天气不错！",
	"再见":      "再见",
}

func main() {
	ln, err := net.Listen("tcp", ":8080") //tcp: Transfer Control Protocol 传输控制协议;

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

		go talk(conn)
	}
}

func talk(conn net.Conn) {
	defer conn.Close()

	for {
		buf := make([]byte, 1024)
		valid, err := conn.Read(buf)
		if err != nil {
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
	}
}
