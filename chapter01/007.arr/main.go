package main

import "fmt"

func main() {
	var age int = 30
	fmt.Println("age是：", age)

	var ages [5]int = [5]int{30, 31, 32, 33, 34}
	fmt.Println("ages是：", ages)
	var ages2 = [5]int{18, 19, 20, 21, 22}
	fmt.Println("ages2是", ages2)
	ages3 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("ages3是", ages3)
	ages2 = ages3
	fmt.Println("二次赋值后的ages2:", ages2)
	// ages3 = [6]int{35, 34, 33, 32, 31, 30} // 错误：数组长度不能变

	//var xxxx type
	var ages4 [3]int
	fmt.Println("ages4:", ages4)
	ages4[0] = 1000
	ages4[1] = 2000
	ages4[2] = 3000
	fmt.Println("ages4:", ages4)
	// ages4[-1] = -1 // 错误，越界
	// ages4[99] = -1 // 错误，越界

	for i := 0; i < len(ages4); i++ {
		fmt.Println(ages4[i])
	}

	for i, val := range ages4 {
		fmt.Println(ages4[i], "====>", i, "->", val)
	}
}
