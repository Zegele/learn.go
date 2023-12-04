package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	quit := make(chan bool)

	// 新开一个协程
	go func() {
		for {
			select {
			case num := <-ch:
				fmt.Println("num = ", num)
			case <-time.After(3 * time.Second):
				fmt.Println("超时")
				quit <- true
			}
		}
	}()

	for i := 0; i < 5; i++ {
		ch <- i
		time.Sleep(time.Second)
	}

	<-quit
	fmt.Println("程序结束")
}

/*
// www.kancloud.cn/imdszxs/golang/1509644
ch := make(chan int, 1)
for{
	select{
	case ch <- 0:
		case ch <-1:
}
}
i := <-ch
fmt.Println("Value received: ", i)

*/
