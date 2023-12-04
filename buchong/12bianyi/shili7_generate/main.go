// go generate 命令 ————在编译前自动化生成某类代码
// c.biancheng.net/view/4442.html
package main

import "fmt"

// go generate命令时go1.4新添加的命令，
//当运行该命令时，它将扫描与当前包相关的源码文件，找出所有包含//go:generate的特殊注释
//提取并执行该特殊注释后面的命令
//
//使用go generate命令时有以下几点需要注意：
//1. 该特殊注释必须在.go源码文件中
//2. 每个源码问价你可以包含多个generate特殊注释
//3. 运行 go generate 命令时，才会执行特殊注释后面的命令
//4. 当go generate 命令执行出错时，将终止程序的运行
//5. 特殊注释必须以//ge:generate开头，双斜线后面没有空格

//在下面这些场景下，我们会使用 go generate 命令：
//yacc : 从.y文件生成.go文件
//protobufs:从protocol buffer定义文件（.proto）生成.pb.go文件
//Unicode： 从UnicodeData.txt生成Unicode表
//HTML：将HTML文件嵌入到go源码
//bindata： 将形如JPEG这样的文件转成go代码中的字节数组
//
//再比如：
//string方法：为类似枚举常量这样的类型生成String()方法
//宏：为既定的泛型包生成特定的实现，比如用于ints的sort.Ints
//
//go generate 命令格式如下所示:
//go genenrate [-run regexp] [-n] [-v] [-x] [command] [build flags] [file.go...|packages]
//参数说明：
//-run 正则表达式匹配命令行，仅执行匹配的命令
//-v 输出被处理的包名和源文件名
//-n 显示，不执行命令
//-x 显示，并执行命令
//command 可以是在环境变量PATH中的任何命令
//
//执行go generate 命令时，也可以使用一些环境变量：
//GOARCH 体系架构（arm, amd64等）
//GOOS 当前的OS环境（linux，windows等）
//GOFILE 当前处理中的文件名
//GOLINE 当前命令行在文件中的行号
//GOPACKAGE 当前处理文件的包名
//DOLLAR 固定的$ ，不清楚具体用途

// ，
//
//go:generate go run main.go
//go:generate go version
func main() {
	fmt.Println("golang")
}

// 执行 go generate -x 命令
