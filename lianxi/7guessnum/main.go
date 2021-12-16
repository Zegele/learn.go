package main

import (
	"fmt"
)

var min int
var max int
var guessnum int
var testnum int
var putout string

func main() {
	//设置猜数字范围
	setRange()

	//设置要猜的数字（防止忘记该数字或中途改变主意等）
	setNum()

	//开始猜数字
	fmt.Println("猜数字游戏开始： ")
	guess(min, max)
}

func setNum() {
	fmt.Print("现在写下你心中的数字：")
	//todo ***表示要猜的数
	for {
		fmt.Scanln(&guessnum)
		if guessnum > max || guessnum < min {
			fmt.Printf("请重新输入要猜的数,%d -- %d。", min, max)
		} else {
			break
		}
	}
}

func setRange() { //设置猜数字的范围
	fmt.Print("我们要玩猜数字的游戏。\n现在需要你确定猜数字的范围，请先输入最小值（整数）： ")
	fmt.Scanln(&min)
	fmt.Print("请再输入最大值（整数）：")
	for {
		fmt.Scanln(&max)
		if max < min {
			fmt.Print("输入有误，请输入更大的值：")
		} else {
			fmt.Printf("OK! 设置好了，范围是：%d -- %d。", min, max)
			break
		}
	}
}

func guess(min, max int) int {
	testnum = max - (max-min)/2 //猜测的数字 符合有负数
	fmt.Println("我猜是： ", testnum)
	fmt.Print("是大了，还是小了？可输入“大了”或“小了”：")
	// for { 递归就像一种循环，这里不用for循环
	fmt.Scanln(&putout)
	if putout == "大了" || putout == "b" {
		if testnum == min {
			fmt.Println("不猜了。 你怕是在玩儿我！")
			return testnum
		}
		guess(min, testnum-1)
	} else if putout == "小了" || putout == "s" {
		if testnum == max {
			fmt.Println("不猜了。 你怕是在玩儿我！")
			return testnum
		}
		guess(testnum+1, max)
	} else if putout == "yes" || putout == "y" {
		fmt.Printf("我猜对了！！！果然是：%d", testnum)
	} else {
		println("输入有误，请输入“大了”或“小了”")
	}
	//}不要for循环
	return testnum
}
