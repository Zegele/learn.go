package main

//www.kancloud.cn/imdszxs/golang/1509648

import (
	"fmt"
	"runtime"
)

// 环境变量设置使用CPU核数
// 设置环境变量 GOMAXPROCS的值来控制使用多少个CPU核心。
// 具体操作方法是通过直接设置环境变量GOMAXPROCS的值，或者在代码中启动goroutine之前先调用以下这个语句以设置使用cpu核心：
// runtime.GOMAXPROCS(16) // 使用16个核心

// 到底应该设置多少个CPU核心呢，其实runtime包中还提供了另外一个NumCPU()函数来获取核心数，示例：

func main() {
	cpuNum := runtime.NumCPU() // 获取当前设备的cpu核心数
	fmt.Println("cpu核心数：", cpuNum)

	runtime.GOMAXPROCS(cpuNum) // 设置需要用到的cpu数量
}

/*
// 模拟完全可以并行的计算任务
// 计算N个整数的总和。 我们可以将所有整形数分成M份，M即CPU的个数。
// 让每个CPU开始计算分给它的那份计算任务，最后将每个CPU的计算结果再做一次累加，
// 这样就可以得到所有N个整型数的总和

type Vector []float64

// 分配给每个CPU的计算任务
func (v Vector) DoSome(i, n int, u Vector, c chan int) {
	for ; i < n; i++ {
		v[i] += u.Op(v[i])
	}
	c <- 1 // 发信号告诉任务管理者我已经计算完成了

}

const NCPU = 16 // 假设总共有16核

func (v Vector) DoAll(u Vector) {
	c := make(chan int, NCPU) // 用于接收每个CPU的任务完成信号
	for i := 0; i < NCPU; i++ {
		go v.DoSome(i*len(v)/NCPU, (i+1)*len(v)/NCPU, u, c)
	}
	//等待所有CPU的任务完成
	for i := 0; i < NCPU; i++ {
		<-c // 获取到一个数据， 表示一个cpu计算完成
	}
	// 到这里表示所有计算已经结束
}

*/

/*
func main() {
	for i := 0; i < 5; i++ {
		go AsyncFunc(i)
	}
	time.Sleep(time.Second * 1)
}

func AsyncFunc(index int) {
	sum := 0
	for i := 0; i < 10000; i++ {
		sum += 1
	}
	fmt.Printf("线程%d, sum为：%d\n", index, sum)
}


*/
