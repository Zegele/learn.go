package main

// 通过stringer工具来演示go generate命令的使用
// stringer并不是Go语言自带的工具，需要手动安装
// 我们可以通过下面的命令来安装stringer工具
// go get golang.org/x/tools/cmd/stringer
// 以上用fq， 然后git没成功，下次用goland试试
//
// 上面的命令需要fq。条件不允许的化也可以通过github上的镜像来安装：
// git clone https://github.com/golang/tools/ $GOPATH/src/golang.org/x/tools
// 把clone的tools包，安装在$GOPATH/src/golang.org/x/tools目录下
// go install golang.org/x/tools/cmd/stringer
// 在bin目录下生成string.exe文件
// install用git没有成功，使用goland成功
// 安装好的stringer工具位于GOPATH/bin 目录下，
// 想要正常使用它，需要先将GOPATH/bin目录添加到系统的环境变量PATH中
//
// 示例strings工具实现String()方法：
// 1. 在项目目录下新近啊一个painkiller文件夹，并在该文件中创建painkiller.go，如下：
// 下接 painkiller/painkiller.go 中、
func main() {

}
