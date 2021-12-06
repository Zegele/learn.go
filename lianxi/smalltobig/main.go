package main

import "fmt"

func main() {
	//var arr = []int{7, 8, 5, 1, 3, 9, 6, 0, 2, 4}
	var arr = []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	fmt.Println(arr)

	for j := 0; j < len(arr)-1; j++ { //循环次数
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				//temp := arr[i]
				//arr[i] = arr[i+1]
				//arr[i+1] = temp
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
	}

	fmt.Println(arr)
}
