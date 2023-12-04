//部署Go语言程序到Linux服务器
//www.kancloud.cn/imdszxs/golang/1509751

package main

import (
	"fmt"
	"log"
	"net/http"
)

//平时都是在本地开发调试访问的
//那怎么打包到服务器上呢？ 如下：

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "golang")
}

func main() {
	http.HandleFunc("/", sayHello)
	log.Println("启动成功，可以通过localhost:9000访问")
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("List 9000")
	}
}

//上面代码，现在编译，由于是window环境编译到linux下运行
//所有涉及到跨平台编译，如下
//set GOARCH=amd64 //设置目标可执行程序操作系统构架，包括：386，amd64, arm
//set GOOS=linux //设置可执行程序运行操作系统，支持darwin，freebsd，linux，windows
//go build ./main.go
//注意：使用Windows 10系统必须用cmd工具执行上述命令，不能使用powershell
//
//编译完成后会生成一个main可执行文件，没有后缀，
//这时只需要把这个文件上传到虚拟机，直接运行就好了。
