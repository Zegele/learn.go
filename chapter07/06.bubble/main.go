package main

import (
	"fmt"
	"math/rand"
	"time"
)

func bubble(sli *[]int) {
	for i := 0; i < len(*sli)-1; i++ {
		for j := 0; j < len(*sli)-i-1; j++ {
			if (*sli)[j] > (*sli)[j+1] {
				(*sli)[j], (*sli)[j+1] = (*sli)[j+1], (*sli)[j]
			}
		}
		//fmt.Println("中间状态：", *sli)
	}
	//fmt.Println("最终状态：", *sli)
}

func main() {
	arrSize := 10
	sli := make([]int, 0, arrSize)
	for i := 0; i < arrSize; i++ {
		sli = append(sli, rand.Intn(50))
	}

	start := time.Now()
	fmt.Println(sli)
	bubble(&sli)
	finish := time.Now()
	fmt.Println(finish.Sub(start))
	fmt.Println(sli)
}
