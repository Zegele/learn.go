package main

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	t.Log("hello")
}

//2. 运行指定单元测试用例
//go test指定文件时默认执行文件内的所有测试用例
//可以使用-run 参数选择需要的测试用例单独执行，如下 ：
//一个文件包含多个测试用例

func TestA(t *testing.T) {
	t.Log("A")
}
func TestAK(t *testing.T) {
	t.Log("AK")
}
func TestB(t *testing.T) {
	t.Log("B")
}
func TestC(t *testing.T) {
	t.Log("C")
}

//go test -v -run TestA hello_test.go
//...PASS TestA  TestAK
//发现TestA和TestAK的测试用例都被执行
//原因是-run跟随的测试用例的名称支持正则表达式
//使用-run TestA$ 既可执行TestA测试用例

// 3. 标记单元测试结果
// 当需要终止当前测试用例时，可以使用t.FailNow()，参考：
func TestFailNow(t *testing.T) {
	t.FailNow() // 测试终止，后面的不执行
	fmt.Println("--------")
}

// 还有一种只标记错误，不终止测试的方法————t.Fail()，如下：
func TestFail(t *testing.T) {
	fmt.Println("before fail")
	t.Fail() //测试错误标记，继续测试后面的
	fmt.Println("after fail")
}

//4.单元测试日志
//每个测试用例可能并发执行，
//使用testing.T提供的日志输出可以保证日志跟随这个测试上下文一起打印输出
//testing.T提供了几种日志输出方法：
//1.  t.Log() 打印日志，同时结束测试
//2.  Logf() 格式化打印日志，同时结束测试
//3.  Error() 打印错误日志，同时测试结束
//4.  Errorf() 格式化打印错误日志，同时结束测试
//5.  Fatal() 打印致命日志，同时结束测试
//6.  Fatalf() 格式化打印致命日志，同时结束测试
//
//
//
//重要：
//二、基准测试：获取代码内存占用和运行效率的性能数据(接benchmark_test.go)
