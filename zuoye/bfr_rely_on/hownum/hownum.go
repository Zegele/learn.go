package hownum

import "fmt"

func Hownum() (num int) {
	fmt.Print("你要录入几个人的数据？请输入数字：")
	for {
		fmt.Scanln(&num)
		if num <= 0 {
			fmt.Println("请输入一个大于0的数字")
		} else {
			return
		}
	}
}
