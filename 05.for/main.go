package main

import "fmt"

func main() {
	fmt.Println("round 1")
	for i := 0; i < 5; i++ {
		fmt.Println("golang")
	}

	fmt.Println("round 2")
	j := 0
	for ; j < 5; j++ {
		fmt.Println("round2, golang")
	}

	fmt.Println("round 3")
	k := 0
	for k < 5 {
		fmt.Println("round3 golang")
		k++
	}

	fmt.Println("round 4")
	l := 0
	for {
		fmt.Println("round4 golang")
		l++
		if l >= 3 {
			break
		}
	}

	fmt.Println("round 5")
	m := 0
	for {
		fmt.Println("round 5 golang", m)
		m++
		if m >= 10 {
			break //直接结束循环
		}

		if m%2 == 0 {
			fmt.Println(m, "被continue了")
			continue // 不执行以下了，继续下次循环
		}
		fmt.Println("练习跳过", m)

	}
}
