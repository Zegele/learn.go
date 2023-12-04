package main

import (
	"fmt"
	"time"
)

// c.biancheng.net/view/vip_7356.html 看不了
// www.niuguwen.cn/shouce/art-8072.html
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go pump1(ch1)
	go pump2(ch2)
	go suck(ch1, ch2)
	time.Sleep(time.Second)
}

func pump1(ch chan int) {
	for i := 0; ; i++ {
		ch <- i * 2
	}
}

func pump2(ch chan int) {
	for i := 0; ; i++ {
		ch <- i + 5
	}
}

func suck(ch1, ch2 chan int) {
	for {
		select {
		case v := <-ch1:
			fmt.Printf("Received on channel 1 : %d", v)
		case v := <-ch2:
			fmt.Println("Received on channel 2: %d", v)
		}
	}
}
