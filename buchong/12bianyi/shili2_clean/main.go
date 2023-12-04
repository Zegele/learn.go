// go clean命令————清除编译文件
// c.biancheng.net/view/4440.html
package main

import "fmt"

//go 语言中go clean 命令可以移除当前源码包和关联源码包里面编译生成的文件
//这些文件包括以下：
//1.执行go build命令是在当前目录下生成的与包名或者go源码文件同名的可执行文件。
//在windows下，则是与包名或go源码文件同名且带有.exe后缀的文件
//2. 执行go test 命令并加入 -c 标记是在当前目录下生成以包名加.test后缀为名的文件。
//在windows下，则是以包名加.test.exe后缀的文件
//3. 执行go install命令安装当前代码包时产生的结果
//如果当前代码包中只包含库源码文件，则结果文件指的就是在工作区pkg目录下相应的归档文件
//如果当前代码包中只包含一个命令源文件，则结果文件指的就是在工作区bin目录下的可执行文件
//4. 在编译Go或c源码文件时遗留在相应目录中的文件或目录
//包括："_obj"和"_test"目录，名称为"_testmain.go" "test.out"  "build.out"或"a.out"的文件名
//名称以".5" ".6"  ".a"  ".o"   ".os"为后缀的文件
//这些目录和文件是在执行go build命令时生成在临时目录中的
//
//go clean命令就像java中的maven clean命令一样， 会清除掉编译过程中产生的一些文件
// go clean -i -n
//通过上面的示例可以看出，go clean还可以指定一些参数
//对应的参数的含义如下：
//-i 清除关联的安装的包和可运行文件，也就是通过go install安装的文件
//-n 把需要执行的清除命令打印出来，但是不执行，这样就可以很容易的知道底层是如何运行的
//-r 循环的清除在import中引用的包
//-x 打印出来执行的详细命令，其实就是-n打印的执行版本
//-cache 删除所有go build命令的缓存
//-testcache 删除当前所有的测试结果
//
//实际开发中go clean命令使用的可能不是很多，一般都是利用go clean命令清除编译文件，
//然后再将源码递交到github上，方便管理
//
//go clean -n
//rm -f xxx.exe  xxx.test  xxx.test.exe main  main.exe ....
//
//go clean -x
//把可执行文件就删掉了。

func main() {
	fmt.Println("hello")
}
