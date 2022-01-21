package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

func main() {
	fmt.Println("开始数")
	var totalCount int64 = 0
	var t float64 = 3.14
	fmt.Println(int(t))
	for p := 0; p < 5000; p++ {
		fmt.Print("正在统计第", p, "页")
		time.Sleep(1 * time.Second)
		r, _ := rand.Int(rand.Reader, big.NewInt(800))
		fmt.Println("有", r.Int64(), "字") //r.Int64() 为什么这种方式把r转化为int64类型。 big包的方法
		totalCount += r.Int64()
	}
}
