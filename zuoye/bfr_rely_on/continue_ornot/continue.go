package continue_ornot

import "fmt"

func ContinueOrNot() bool {
	var continueornot string
	fmt.Println("是否继续？(y/n)")
	fmt.Scanln(&continueornot)
	for {
		if continueornot == "n" {
			return false
		} else if continueornot == "y" {
			return true
		} else {
			fmt.Println("输入有误，请输入 y 或 n 。")
		}
	}
}
