package main

import "fmt"

func main() {
	var money int = 20
	var busy bool = true

	//实例1：
	switch money {
	case 20:
		fmt.Println("点个外卖")
		fallthrough //穿过向下执行。这里穿过后，执行case 200:
	case 200:
		fmt.Println("下个馆子")
		if busy {
			break //立刻结束
		} else {
			fmt.Println("再吃点零食")
		}
		//...
	}

	//实例2：
	switch {
	case money == 20: //也可以把条件写在case后
		fmt.Println("点个外卖")
		fallthrough //穿过向下执行。这里穿过后，执行case 200:
	case money == 200:
		fmt.Println("下个馆子")
		if busy {
			break //立刻结束循环
		} else {
			fmt.Println("再吃点零食")
		}
		//...
	}
	fmt.Println("end")
}
