package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Runner struct {
	Name string
}

func (r Runner) Run(startPointWg, wg *sync.WaitGroup) { // 参数是两个等待组
	defer wg.Done()
	startPointWg.Wait()
	start := time.Now()
	fmt.Println(r.Name, "开始跑@", start)
	rand.Seed(time.Now().UnixNano())                          //用时间戳去随机数的Seed
	time.Sleep(time.Duration(rand.Uint64()%10) * time.Second) //随机数在Uint64中取（非负），%10是取10以内的。
	//Sleep需要Duration类型的数据
	finish := time.Now()
	fmt.Println(r.Name, "跑到终点，用时：", finish.Sub(start))
}

func main() {
	runnerCount := 10
	runners := []Runner{} //多个runner的切片

	wg := sync.WaitGroup{}
	wg.Add(runnerCount) //wg是一个等待组（等待数量是runnerCount）

	startPointWg := sync.WaitGroup{}
	startPointWg.Add(1) //startPointWg 是另一个等待组（等待数量是1）

	//生成Runner结构体，和Runners切片
	for i := 0; i < runnerCount; i++ {
		runners = append(runners, Runner{
			Name: fmt.Sprintf("*%d*", i),
		})
	}

	for _, runnerItem := range runners {
		go runnerItem.Run(&startPointWg, &wg)
	}

	fmt.Println("各就位")
	//time.Sleep(1 * time.Second)
	fmt.Println("预备：跑")

	startPointWg.Done()

	wg.Wait()
	fmt.Println("赛跑结束")

}
