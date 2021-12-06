package main

import "fmt"

func main() {
	var m1 map[string]int                  //只是声明，没有初始化，是nil状态的
	m1["a"] = 1                            //panic on nil map
	delete(m1, "a")                        //没有初始化可以删除
	fmt.Println("m1 没有实例化，直接取数：", m1["a"]) //没有初始化可以读
	//断言
	//value, ok:=map[key] 可选返回值

	m2 := map[string]int{}
	m3 := map[string]int{"jj": 50, "kk": 40}

	mSurprise := map[string]map[string]map[float64]int{}
}
