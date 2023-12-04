// go build命令（go语言编译命令）完全攻略
// c.biancheng.net/view/120.html
package main

import "fmt"

//go 编译速度非常快，1.9版本后，默认利用go语言的并发特性进行函数粒度的并发编译
//go语言的程序编写基本以源码方式，无论是自己的代码还是第三方代码
//并且以GOPATH作为工作目录和一套完整的工程目录规则
//因此go语言中日常编译时无须像C++一样配置各种包含路径、链接库地址等
//
//go语言中使用go build命令只要用于编译代码
//在包的编译过程中，若有必要，会同时编译与之相关联的包
//
//go build有很多种编译方法，如无参数编译，文件列表编译，指定包编译等
//使用这些方法都可以输出可执行文件
//

//1. go build 无参数编译
//。

func main() {
	// 同包的函数
	pkgFunc()
	fmt.Println("hello world")
}

//如果源码中没有依赖GOPATH的包引用，那么这些源码可以使用无参数go build
//go build
//在代码所在目录，下使用go build命令
//生成.exe执行文件（自动生成的文件名是用了文件夹的名字）。
//go build //在编译开始时，会搜索当前目录的go源码，
//这个例子中，go build会找到lib.go和main.go两个文件
//编译这两个文件后，生成当前目录名的可执行文件，并放置于当前目录下
//ls 列出当前目录的文件
//cd 进入某个目录
//./xxxx.exe 或 ./xxxx 运行当前目录的可执行文件
//
//2. go build + 文件列表
//编译同目录的多个源码文件时，可以在go build的后面提供多个文件名
//go build会编译这些源码，输出可执行文件
//go build+文件列表 的格式如下：
//go build file1.go file2.go
//在代码所在目录中使用go build，之后添加要编译的源码文件名：
//
//go build main.go lib.go
//ls
//lib.go main main.go
//./main 或 ./main.exe
//go build lib.go main.go
//ls
//lib lib.go main main.go
//注意：使用 go build+文件列表 方式编译时，
//可执行文件默认选择文件列表中第一个源码文件作为可执行文件名输出
//
//如果需要指定输出可执行文件名
//可以使用-o参数，
//go build -o myexec main.go lib.go
//ls
//lib.go main.go myexec
//./myexec
// 注意 我的windows怎么没有生成 myexec.exe ，只有myexec文件
//因为要这样写
//go build -o myexec.exe main.go lib.go
//才会生成可运行的.exe文件

//注意：
//使用 go build+文件列表 编译方式编译时，文件列表中的每个文件必须是同一个包的go源码
//也就是说，不能像c++一样将所有工程的go源码使用文件列表方式进行编译
//编译复杂工程时需要用“指定包编译”的方式
//"go build+文件列表"方式更适合使用go语言编写的只有少量文件的工具
//
// 下到 buildbao
