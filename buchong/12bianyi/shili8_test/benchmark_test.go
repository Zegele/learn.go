package main

import (
	"fmt"
	"testing"
)

//二、基准测试：获取代码内存占用和运行效率的性能数据(接benchmark_test.go)
//基准测试可以测试一段程序的运行性能及耗费cpu的程度
//Go语言中提供了基准测试框架
//使用方法类似单元测试，使用者无须准备高精度的计时器和各种分析工具
//基准测试本身既可以打印出非常标准的测试报告

//1.基础测试基本使用，如下：

func Benchmark_Add(b *testing.B) {
	var n int
	for i := 0; i < b.N; i++ {
		n++
	}
}

//这段代码使用基准测试框架，测试加法性能。
// b.N由基准测试框架提供
// 测试代码需要保证函数可重入性及无状态
//测试代码不使用全局变量等带有记忆性质的数据结构体
//避免多次运行统一段代码时的环境不一致，不能假设N值范围
//go test -v -bench=. hello_test.go
//> go test -v -bench=Benchmark_Add benchmark_test.go
//Benchmark_Add
//Benchmark_Add-4         1000000000               0.3490 ns/op
//PASS
//ok      command-line-arguments  0.677s

// -bench=. 表示运行benchmark_test.go文件里的所有基准测试
//和单元测试用的-run类似， -bench=函数名（要测试的）
//
//1000000000表示测试的次数，也就是testing.B结构中提供给程序使用的N
//0.349ns/op 表示每一次操作耗费多少时间（纳秒）
//注意：windows下使用go test 命令时，
//-bench=. 应写成 -bench="." 或 -bench='.'。

//2. 基准测试原理
//基准测试框架对一个测试用例的默认测试时间是1秒。
//开始测试时，当以Benchmark开头的基准测试用例函数返回时还不到1秒，
//那么testing.B中的N值将按1、2、5、10、20、50...递增
//同时以递增后的值重新调用基准测试用例函数

//3. 自定义测试时间
//通过 -benchtime参数可以自定义测试时间，如：
//go test -v -bench='.' -benchtime=5s benchmark_test.go
//就是按5秒测试

// 4. 测试内存
// 基准测试可以对一段代码可能存在的内存分配进行统计
// 下面是使用字符串格式化的函数，内部会进行一些分配操作
// .
func Benchmark_Alloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%d", i)
	}
}

// 在命令行中添加-benchmen参数以显示内存分配情况，如
// go test -v -bench=Alloc -benchmem benchmark_test.go
// -bench=Alloc 指定测试Alloc 函数。 -bench=支持正则表达
// 16 B/op 表示每一次调用需要分配16个字节，
// 2 allocs/op 表示每次调用由两次分配
//
// 开发者根据这些信息可以迅速找到可能的分配点，进行优化和调整
//
// 5. 控制计数器
// 有些测试需要一定的启动和初始化时间
// 如果从Benchmark() 函数开始计时会很大程度上影响测试结果的精准性
// testing.B提供了一系列的方法可以方便地控制计时器
// 从而让计时器只在需要的区间进行测试。
// 同门通过下面的代码来了解计时器的控制
// ：
func Benchmark_Add_TimeControl(b *testing.B) {
	var n int
	//停止计时器
	b.StopTimer()

	// 重置计时器
	b.ResetTimer()

	// 开始计时器
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		n++
	}
}

// Benchmark()函数开始，Timer就开始计数
//StopTimer()可以停止这个计数过程，然后可做一些耗时的操作
//通过StartTimer() 重新开始计时，
//ResetTimer()可以重置计数器的数据
//计数器内部不仅包含耗时数据，还包括内存分配的数据，
