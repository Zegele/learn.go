package main

import "fmt"

func main() {
	//year month day week
	//普通年
	//闰年

	//计算闰年
	var year int

	var yeararr [12][31]int

	for {
		fmt.Print("请输入年份：")
		fmt.Scanln(&year)

		for i := 1; i <= 12; i++ {
			if i == 1 || i == 3 || i == 5 || i == 7 || i == 8 || i == 10 || i == 12 { //31
				for t := 1; t <= 31; t++ {
					yeararr[i-1][t-1] = t
				}
			} else if i == 2 { //2月
				if (year%4 == 0 && year%100 != 0) || year%400 == 0 { //闰年公式
					for t := 1; t <= 29; t++ { //29
						yeararr[1][t-1] = t
					}
				} else {
					for t := 1; t <= 28; t++ { //28
						yeararr[1][t-1] = t
					}
				}
			} else {
				for t := 1; t <= 30; t++ { //30
					yeararr[i-1][t-1] = t
				}
			}
		}
		break
	}

	fmt.Println(year)
	fmt.Println(yeararr)
}
