package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	var p string
	flag.StringVar(&p, "path", ".", "the path to expose as http") //路径
	// "path"是命令参数的名字，"."是默认值 ，
	var port int
	flag.IntVar(&port, "port", 8089, "the port to expose") // 命令行的使用方法
	//./00_cool.exe --help  查看有哪些参数可用。
	//./00_cool.exe --port 8080 把端口设置成8080 运行。
	flag.Parse()
	http.Handle("/", http.FileServer(http.Dir(p)))                // p是文件路径
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil)) //port是端口

	//http.Handle("/", http.FileServer(http.Dir("E:/Geek/src/learn.go")))
	//log.Fatal(http.ListenAndServe(":8089", nil))
}
