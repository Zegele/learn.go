package main

import (
	"fmt"
	"sync"
)

type rank struct {
	standard []string
}

var globalRank = &rank{}
var once sync.Once = sync.Once{}

//var globalRankInitialized bool
//var globalRankInitializedLock sync.Mutex

//func init() {
//	globalRank.standard = []string{"Asia"} //初始化rank
//}

func initGlobalRankStandard(standard []string) { //初始化rank // 将该函数放入go routine 中执行
	once.Do(func() { //Do的参数是一个函数
		globalRank.standard = standard //要执行初始化的内容
	})
}

func main() {
	fmt.Println(globalRank)
	standard := []string{"asia"}
	for i := 0; i < 10; i++ { //即使这里有很多个goroutine，但只会执行一次initGlobalRankStandard函数。
		go func() {
			initGlobalRankStandard(standard)
		}()
	}
	fmt.Println(globalRank)
}
