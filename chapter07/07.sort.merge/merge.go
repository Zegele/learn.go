package main

import (
	"fmt"
	"math/rand"
	"time"
)

func mergeSort(sli []int) []int {
	left, right := sli[:len(sli)/2], sli[len(sli)/2:] //老师这样写：(sli)[:len(sli)/2],(sli)[len(sli)/2:]
	if len(sli) <= 2 {
		return mergeSli(left, right)
	} else {
		return mergeSli(mergeSort(left), mergeSort(right)) //递归了
	}
}

func mergeSli(left, right []int) []int {
	out := []int{}
	leftI, rightI := 0, 0
	for {
		if leftI == len(left) || rightI == len(right) {
			break
		}
		if left[leftI] < right[rightI] {
			out = append(out, left[leftI])
			leftI++
			continue //进入下一次循环
		} else {
			out = append(out, right[rightI])
			rightI++
			continue
		}
	}
	for ; leftI < len(left); leftI++ {
		out = append(out, left[leftI])
	}
	for ; rightI < len(right); rightI++ {
		out = append(out, right[rightI])
	}
	return out
}

func main() {
	SliSize := 10000000 //一千万个数3.6秒
	sli := []int{}
	rand.Seed(time.Now().UnixNano()) //随机种子
	for i := 0; i < SliSize; i++ {
		sli = append(sli, rand.Intn(50))
	}
	//fmt.Println(sli)

	start := time.Now()
	/*sorted :=*/ mergeSort(sli)
	finish := time.Now()
	fmt.Println(finish.Sub(start))
	//fmt.Println(sorted)
}
