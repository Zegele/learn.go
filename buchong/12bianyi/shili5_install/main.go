// go install命令————编译并安装
// c.biancheng.net/view/122.html
package main

//go install 与 go build 命令类似，附加参数绝大多数可以与go build通用
//go install 只是将编译的中间文件放在GOPATH的pkg目录下，以及固定地将编译结果放在GOPATH的bin目录下
//以及固定地将编译结果放在GOPATH的bin目录下
//
//这个命令在内部实际上分成了两步操作，第一步时生成结果文件（可执行文件或者.a包）
//第二部会把编译好的结果移到pkg或bin目录下
//
//go install的编译过程有如下规律：
//1. go install 是建立在GOPATH上的，无法在独立的目录里使用go install
//2. GOPATH下的bin目录放置的是使用 go install 生成的可执行文件，
//可执行文件的名称来自于编译时的包名
//3. go install 输出目录始终为GOPATH下的bin目录，无法使用-o附加参数进行自定义
//4. GOPATH下的pkg目录放置的是编译期间的中间文件。

func main() {

}
