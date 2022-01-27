package main

import (
	"fmt"
	"sync"
)

type rank struct {
	standard []string
}

var globalRank = &rank{}
var globalRankInitialized bool
var globalRankInitializedLock sync.Mutex

//func init() {
//	globalRank.standard = []string{"Asia"} //初始化rank
//}

func initGlobalRankStandard(standard []string) { //初始化rank
	globalRankInitializedLock.Lock()
	defer globalRankInitializedLock.Unlock()
	if globalRankInitialized { //如果为true 说明已经被初始化了。return
		return
	} //如果没有被初始化，就向下执行初始化，然后把false变为true
	globalRank.standard = standard
	globalRankInitialized = true
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
