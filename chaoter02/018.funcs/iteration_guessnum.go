package main

import (
	"fmt"
)

var min int
var max int

func main() {
	fmt.Print("我们要玩猜数字的游戏。现在需要你确定猜数字的范围，请先输入最小值（整数）： ")
	fmt.Scanln(&min)
	fmt.Print("请再输入最大值（整数）：")
	for {
		fmt.Scanln(&max)
		if max < min {
			fmt.Print("输入有误，请输入更大的值。")
		}
		break
	}

}
func guess(left, right uint) {

}
