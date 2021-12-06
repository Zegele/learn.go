package main

import "fmt"

func main() {
	var arr1 = [5]int{1, 2, 3, 4, 5}

	var temparr [5]int
	for i := 0; i < len(arr1); i++ {
		temparr[i] = arr1[len(arr1)-1-i]
	}
	fmt.Println(temparr)

	for j, _ := range temparr {
		arr1[j] = temparr[j]
	}
	fmt.Println(arr1)
}
