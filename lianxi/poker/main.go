package main

import (
	"fmt"
	"strconv"
)

func main() {
	var poker = []string{}
	//生成poker 52张
	for i := 0; i < 13; i++ {
		poker = append(poker, "红"+strconv.Itoa(i+1))
		poker = append(poker, "黑"+strconv.Itoa(i+1))
		poker = append(poker, "花"+strconv.Itoa(i+1))
		poker = append(poker, "方"+strconv.Itoa(i+1))
	}
	fmt.Println("扑克原始顺序")
	fmt.Println(poker)

	//打乱poker
	fmt.Println("开始洗牌")

	var n int
	fmt.Print("你要洗几遍？（输入数字）")
	fmt.Scanln(&n)

	for k := n; k > 0; k-- {
		j := 0
		var temppocker []string
		for i := 25; i >= 0; i-- { //均匀分成2份后，由下向上各添加一张牌（第26张，和第52张叠加，然后是第25和第51叠加。。。）  类似现实中的洗牌。
			temppocker = append(temppocker, poker[i], poker[51-j])
			j++
		}

		//把洗好顺序的temppoker，再写入原poker。同时也方便下次循环
		for l := 0; l < 52; l++ {
			poker[l] = temppocker[l]
		}
	}
	fmt.Println("pocker是：", poker)

}
