// go test 命令 （Go 语言测试命令）完全攻略
// c.biancheng.net/view/124.html
package main

//go语言拥有一套单元测试和性能测试系统
//仅需要添加很少的代码就可以快速测试一段需求代码
//go test命令，会自动读取源码目录下面名为xxx_test.go的文件
//生成并运行测试用的可执行文件，输出的信息如下：
//ok archive/tar 0.011s
//FAIL archive/tip 0.22s
//ok compress/gzip 0.033s
//
//性能测试系统可以给出代码的性能数据，帮助测试者分析性能问题
//提示: 单元测试（unit testing），指对软件中的最小可测试单元进行检查和验证
//单元测试实在软件开发过程中要进行的最低级别的测试活动，软件的独立单元将在与程序的其他部分相隔离的情况下进行测试
//单元测试————测试和验证代码的框架
//要开始一个单元测试，需要准备一个go源码文件
//在命名文件时需要让问价你必须以_test结尾
//默认的情况下，go test 命令不需要任何的参数，它会自动把你源码包下面所有test文件测试完毕，当然你也可以带参数
//
//常用参数：
//-bench regexp 执行相应的benchmarks ， 例如 -bench=.
//-cover 开启测试覆盖率
//-run regexp 只运行regexp匹配的函数，例如：-run=Array那么就执行包含有Array开头的函数
//-v 显示测试的详细命令
//单元测试源码文件可以由多个测试用例组成，每个测试用例函数需要以Test为前缀，例如：
//func TestXXX(t *testing.T)
//1. 测试用例文件不会参与正常源码编译，不会被包含到可执行文件中
//2. 测试用例文件使用go test指令来执行，没有也不需要main()作为函数入口
//所有在以_test结尾的源码内以Test开头的函数会自动被执行
//3. 测试用例可以不传入*testing.T参数

//示例1：见hello_test.go

func main() {

}