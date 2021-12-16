package main

import (
	"fmt"
)

func main() {
	var min int
	var max int
	var guessnum int

	//设置猜数字范围
	fmt.Print("我们要玩猜数字的游戏。现在需要你确定猜数字的范围，请先输入最小值（整数）： ")
	fmt.Scanln(&min)
	fmt.Print("请再输入最大值（整数）：")
	for {
		fmt.Scanln(&max)
		if max < min {
			fmt.Print("输入有误，请输入更大的值：")
		} else {
			fmt.Printf("OK! 设置好了，范围是：%d - %d。", min, max)
			break
		}
	}

	//设置要猜的数字（防止忘记该数字或中途改变主意等）
	fmt.Print("写下你心中的数字：")
	//todo ***表示要猜的数
	for {
		fmt.Scanln(&guessnum)
		if guessnum > max || guessnum < min {
			fmt.Printf("请重新输入要猜的数,%d -- %d。", min, max)
		} else {
			break
		}
	}

	//开始猜数字
	fmt.Print("猜数字游戏开始： ")
	guess(min, max, guessnum)
	return

}

func guess(min, max, guessnum int) int {
	var testnum int
	var putout string
	for {
		testnum = max - (max-min)/2
		fmt.Println("我猜是： ", testnum)
		fmt.Scanln(&putout)
		if putout == "大了" || putout == "b" {
			if testnum == min {
				fmt.Println("不猜了。 你怕是在玩儿我！")
				return testnum
			}
			fmt.Println(min, testnum-1)
			guess(min, testnum, guessnum)
		} else if putout == "小了" || putout == "s" {
			if testnum == max {
				fmt.Println("不猜了。 你怕是在玩儿我！")
				return testnum
			}
			fmt.Println(testnum+1, max)
			guess(testnum, max, guessnum)
		} else if putout == "yes" || putout == "y" {
			fmt.Println("我猜对了！！！")
			return testnum
		} else if testnum > max || testnum < min {
			return testnum
			break
		} else {
			println("输入有误，请输入“大了”或“小了”")
		}
	}

	if testnum == guessnum {
		fmt.Println("我猜对了！！！")
		fmt.Println(testnum, "dgagadsg")
	}
	return testnum
}
