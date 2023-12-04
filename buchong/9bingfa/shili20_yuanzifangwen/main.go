package main

import (
	"fmt"
	"sync/atomic"
)

// www.likecs.com/show-306390439.html
// www.kancloud.cn/imdszxs/golang/1509650
// c.biancheng.net/view/vip_7353.html
var (
	// 序列号
	seq int64
)

// 序列号生成器
func GenID() int64 {
	// 尝试原子的增加序列号
	return atomic.AddInt64(&seq, 1)
	/*
		atomic.AddInt64(&seq, 1)
		return seq // 会产生问题
	*/

}

func main() {
	// 生成10个并发序列号
	for i := 0; i < 10; i++ {
		go GenID()
	}

	fmt.Println(GenID())
}

// go run -race reflect_test.go
// -race 参数，开启运行时（runtime）对竞态问题的分析
