package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

var totalCompare int = 0

func main() {
	arr := sampleData
	//fmt.Println(arr)

	startTime := time.Now()
	for i := 0; i < 1000000; i++ { //总用时：5.6527008s 总比较次数： 3165000000
		search(&arr, 501)
		search(&arr, 888)
		search(&arr, 900)
		search(&arr, 3)
	}
	finishTime := time.Now()
	fmt.Println("总比较次数：", totalCompare)
	fmt.Println("总用时：", finishTime.Sub(startTime))
}

func search(arrP *[]int64, targetNum int64) bool {
	for _, v := range *arrP {
		totalCompare++
		if v == targetNum {
			return true
		}
	}
	return false
}

func generateRandomData(size int) []int64 {
	arr := make([]int64, 0, size)

	for i := 0; i < size; i++ {
		i, _ := rand.Int(rand.Reader, big.NewInt(int64(size))) // big.NewInt(i) 的参数，需要int64的
		arr = append(arr, i.Int64())                           // i是 *big.Int类型， i.Int64()函数，将i转成int64类型。
	}
	return arr
}
