package main

import (
	"fmt"
	"sync"
)

/*
func main() {
	iSlice := []int{1, 2, 3, 4, 5, 6, 7}
	wg := sync.WaitGroup{}
	wg.Add(len(iSlice))
	for _, item := range iSlice {//由于goroutine 共享这个item变量，导致错误
		go func() {
			defer wg.Done()
			for i := 0; i < 10; i++ {
				fmt.Println(item)
			}
		}()
	}
	wg.Wait()
}

*/

func main() {
	iSlice := []int{1, 2, 3, 4, 5, 6, 7}
	wg := sync.WaitGroup{}
	wg.Add(len(iSlice))
	for _, item := range iSlice { //由于goroutine 共享这个item变量，导致错误
		go func(newItem int) {
			defer wg.Done()
			for i := 0; i < 10; i++ {
				fmt.Println(i, newItem)
			}
		}(item)
	}
	wg.Wait()
}
