package main

import (
	"fmt"
	"learn.go/buchong/12bianyi/shili1_build/buildbao/mypkg"
)

func main() {
	mypkg.CustomPkgFunc()

	fmt.Println("hello")
}

//3. go build+包
//go build+包 在设置GOPATH后，可以直接根据包名进行编译，即便包内文件被增加删除也不影响编译指令。
//。
//go build -o main.exe learn.go/buchong/12bianyi/shili1_build/buildbao
// 包名是相对于GOPATH下的src目录开始的
//./main.exe

//go build learn.go/buchong/12bianyi/shili1_build/buildbao
//./buildbao.exe
// 目录中不要包含中文

//4. go build编译时的附加参数
//go build还有一些附加参数，可以显示更多的编译信息和更多的操作
//-v 编译时显示包
//-p n 开启并发编译，默认情况下该值为cpu逻辑核数
//-a 强制重新构建
//-n 打印编译时会用到的所有命令，但不真正执行
//-x 打印编译时会用到的所有命令
//-race 开启竞态检测
//是按使用频率排列。
