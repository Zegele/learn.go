package main

import "fmt"

func main() {
	leftMap, rightMap := map[string]int{}, map[string]int{} //带{}就算实例化了，可以使用。
	/*//没有{} 没有实例化，不能直接使用，需要再make一遍才可以
	var amap map[string]int
	amap = make(map[string]int)
	fmt.Println(amap)
	*/

	leftMap["语文"] = 80 //普通添加
	rightMap["数学"] = 90
	for k, v := range rightMap { //for range 添加
		leftMap[k] = v
	}
	fmt.Println(leftMap)

}
