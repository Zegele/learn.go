// go pprof 命令（go语言性能分析命令）完全攻略
// c.biancheng.net/view/125.html
package main

import (
	"github.com/pkg/profile"
	"time"
)

//go语言工具链中的go pprof 可以帮助开发者快速分析及定位各种性能问题
//如cpu消耗，内存分配及阻塞分析
//
//性能分析首先需要使用runtime.pprof包嵌入到待分析程序的入口和结束处
//runtime.pprof包在运行时对程序进行每秒100次的采样，最少采样1秒。
//然后将生成的数据输出，让开发者写入文件或者其他媒介上进行分析
//
//go pprof工具链配合Graphviz图形化工具可以将runtime.pprof包生成的数据转换为PDF格式，
//以图片的方式展示程序的性能分析结果
//
//安装第三方图形化显示分析数据工具（Graphviz）
//Graphviz是一套通过文本描述的方法生成图形的工具包，
//描述文本的语言叫做DOT
//在www.graphviz.org获取最新的Graphviz各平台的安装包
//
//安装第三方性能分析来分析代码包
//runtime.pprof提供基础的运行时分析的驱动，但是这套接口使用起来还不是太方便，如：
//1.输出数据使用io.Writer接口，虽然扩展性很强，但是对于实际使用不够方便，
//不支持写入文件
//默认配置项较为复杂
//很多第三方的包在系统包runtime.pprof的技术上进行便利性封装
//让整个测试过程更为方便
//如果发现编程的时候，包导入不了，那就一步步检查，
//包是否get并安装成功，mod文件是否更新并可用，可能还需要 go vendor,
//总之，问题总会被找到并解决。

//这里使用github.com/pkg/profile包进行例子展示，使用下面代码安装这个包：
//go get github.com/pkg/profile
//
//性能分析代码
//下面代码故意制造了一个性能问题，同时使用github.com/pkg/profile包进行性能分析
//
//。

func joinSlice() []string {
	var arr []string
	for i := 0; i < 100000; i++ {
		//故意造成多次的切片添加（append）操作，由于每次操作可能会由内存重新分配和移动，性能较低
		arr = append(arr, "arr")
		// 为了性能分析，这里在已知元素大小的情况下，还是使用append()函数不断地添加切片
		//性能较低，在实际中应该避免，这里为了性能分析，故意这样写。

	}
	return arr
}

func main() {
	// 开始性能分析，返回一个停止接口
	stopper := profile.Start(profile.CPUProfile, profile.ProfilePath("."))
	//frofile.Start调用github.com/pkg/profile包的开启性能分析接口
	//这个Start函数的参数都是可选项，这里需要指定的分析项目是profile.CPUProfile，
	//也就是CPU耗用，profile.ProfilePath(".") 指定输出的分析文件路径，
	//这里指定为当前文件夹
	//profile.Start()函数会返回一个Stop接口，方便在程序结束时结束性能分析

	// 在main()结束时停止性能分析
	defer stopper.Stop()
	//将性能分析在main函数结束时停止

	// 分析的核心逻辑
	joinSlice()
	//开始执行分析的核心（也就是分析的对象，也就是要分析的东西）

	//让程序至少运行1秒
	time.Sleep(time.Second)
	// 为了保证性能分析数据的合理性，分析的最短时间是1秒
	//使用time.Sleep()在程序结束前等待1秒，
	//如果你的层序默认可以运行1秒以上，这个等待可以去掉

}

// go build -o cpu.exe main.go  // build main.go 生成cpu.exe
// ./cpu.exe //运行.exe，当前目录输出cpu.pprof文件  //直接goland中，或双击.exe才好用？ // 这步git中不好用
// go tool pprof --pdf cpu.exe cpu.pprof > cpu.pdf // windows别忘了是cpu.exe，带后缀 //第三步在git上运行，goland运行不了
//使用go tool 工具输入 cpu.pprof和cpu.exe可执行文件，
//生成PDF格式的输出文件
//将输出文件，重定向为cpu.pdf 文件
//这个过程中会调用Graphviz工具.

// 图中的每一个框为一个函数调用的路径，第三个方框中joinSlice函数耗费了50%的cpu时间
//存在性能瓶颈，重新优化代码，在已知切片元素数量的情况下，直接分配内存，如下：
//func joinSlice2() []string {
//	const count = 100000
//	var arr []string = make([]string, count)
//	for i := 0; i < count; i++ {
//		arr[i] = "arr"
//	}
//	return arr
//}

// 查看了pdf貌似不对劲
