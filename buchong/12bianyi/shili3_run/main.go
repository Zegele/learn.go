// go run 命令 ————编译并运行
// c.biancheng.net/view/121.html
package main

import (
	"fmt"
	"os"
)

//go run 命令会编译源码，并且直接执行源码的main()函数，不会在当前目录留下可执行文件
//。

func main() {
	fmt.Println("args:", os.Args)
}

// 这段代码的功能是将输入的参数打印出来
//使用go run 运行这个源码文件：
//go run main.go --filename xxx.go
//go run 不会在运行目录下生成任何文件，
//可执行文件被放在临时文件中被执行
//工作目录被设置为当前目录
//在 go run 的后部可以添加参数，
//这部分参数会作为代码可以接收的命令行输入提供给程序
//
//go run 不能使用 go run +包 的方式进行编译，
//如需快速编译运行包，需要使用如下步骤：
//1. 使用go build生成可执行文件
//2. 运行可执行文件 =。
