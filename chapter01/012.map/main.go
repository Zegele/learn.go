package main

import "fmt"

func main() {
	fmt.Println("定义Map")

	var m1 map[string]int //只是声明，没有初始化(实例化)，是nil状态的
	//等同于 var m1 map[string]int = nil
	//m1["a"] = 1                            //panic on nil map
	delete(m1, "a")                        //没有初始化可以删除
	fmt.Println("m1 没有实例化，直接取数：", m1["a"]) //没有初始化可以读

	m2 := map[string]float64{}
	m3 := map[string]int{"jj": 50, "kk": 40}
	fmt.Println(m1, m2, m3)
	fmt.Println("jj的分数：", m3["jj"])
	fmt.Println("kk的分数：", m3["kk"])

	//断言
	//value, ok:=map[key] 可选返回值
	mmScore, ok := m3["mm"]
	fmt.Println(mmScore, "mm>>>", ok)

	m3["mm"] = 99
	fmt.Println("mm的分数：", m3["小强"])
	mmScore, ok = m3["mm"]
	fmt.Println(mmScore, "mm>>>", ok)

	for name, score := range m3 {
		fmt.Println(name, "=", score)
	}

	//mSurprise := map[string]map[string]map[float64]int{}
}
