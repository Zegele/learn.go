package main

import "fmt"

func main() {
	a := []int{4, 4, 4, 4, 4, 5, 5}
	// 1,2,3,4,5,6,7,8,9

	for i := 0; i < len(a); i++ { //4
		//fmt.Println("i", i, a[i])
		for j := 1 + i; j < len(a); { //5 4 ->45544
			//fmt.Println("j:", j, a[j])
			if a[i] == a[j] {
				a = append(a[:j], a[j+1:]...)
				//fmt.Println(a)
			} else {
				j++
				break
			}
		}
	}
	fmt.Println(a)

	fmt.Println("sdfg")

	//去重
	a = []int{4, 5, 4, 5, 4, 4, 5, 5}
	quchong := SliceQuChong(a)
	fmt.Println(quchong)
	//排序从小到大
	//smallToBig := SliceSmallToBig(quchong)
	//fmt.Println(smallToBig)
	//bigToSmall := SliceBigToSmall(smallToBig)
	//fmt.Println(bigToSmall)
	var t int
	fmt.Println("teacher", t)

}

//去重函数
func SliceQuChong(yuanShiSlice []int) (quChongSlice []int) {
	for i := 0; i < len(yuanShiSlice); i++ {
		fmt.Println("i", i, yuanShiSlice[i])
		for j := i + 1; j < len(yuanShiSlice); {
			fmt.Println("j:", j, yuanShiSlice[j])
			if yuanShiSlice[i] == yuanShiSlice[j] {
				yuanShiSlice = append(yuanShiSlice[:j], yuanShiSlice[j+1:]...)
				fmt.Println(yuanShiSlice)
			} else {
				j++
			}
		}
	}
	quChongSlice = yuanShiSlice
	return quChongSlice
}

//从小到大排序函数
func SliceSmallToBig(quChongSlice []int) (smallToBigSlice []int) {
	for i := 0; i < len(quChongSlice)-1; i++ {
		for j := 0; j < len(quChongSlice); j++ {
			if quChongSlice[j] < quChongSlice[j+1] {
				quChongSlice[j], quChongSlice[j+1] = quChongSlice[j+1], quChongSlice[j]
			}
		}
	}
	smallToBigSlice = quChongSlice
	return
}

//从大到小排序函数
func SliceBigToSmall(quChongSlice []int) (bigToSmallSlice []int) {
	for i := 0; i < len(quChongSlice)-1; i++ {
		for j := 0; j < len(quChongSlice)-1-i; j++ {
			if quChongSlice[j] > quChongSlice[j+1] {
				quChongSlice[j], quChongSlice[j+1] = quChongSlice[j+1], quChongSlice[j]
			}
		}
	}
	bigToSmallSlice = quChongSlice
	return
}
