package main

import (
	"fmt"
	"sync"
	"time"
)

type Store struct { //定义了一个仓库，设置一个最大值，一个实际数量，一个锁
	DataCount int
	Max       int
	lock      sync.Mutex //锁类型也能这么用？
}

type Producer struct{} //定义了一个生产者的结构体

func (Producer) Produce(s *Store) {
	s.lock.Lock()         //在生产者方法中，给仓库上锁
	defer s.lock.Unlock() //给仓库解锁
	if s.DataCount == s.Max {
		fmt.Println("生产者看到库存满了，不生产")
		return
	}
	fmt.Println("开始生成 +1")
	s.DataCount++
}

type Consumer struct{} // 定义了一个消费者的结构体

func (Consumer) Consume(s *Store) { //参数是仓库结构体
	s.lock.Lock()         // 在消费者方法中，给仓库上锁
	defer s.lock.Unlock() // 给仓库解锁
	if s.DataCount == 0 {
		fmt.Println("消费者看到没库存了，不消费")
		return
	}
	fmt.Println("消费者消费-1")
	s.DataCount--
}

func main() {
	s := &Store{
		Max: 10, //设置了最大值，但是没有设置初始的实际值，默认是0个。
	}
	pCount, cCount := 50, 50 //定义了50个消费者(50goroutine)，50个生产者(50个goroutine)？
	for i := 0; i < pCount; i++ {
		go func() {
			for {
				time.Sleep(100 * time.Millisecond)
				Producer{}.Produce(s) //居然能这么直接用结构体
			}
		}()
	}
	for i := 0; i < cCount; i++ {
		go func() {
			for {
				time.Sleep(100 * time.Millisecond)
				Consumer{}.Consume(s)
			}
		}()
	}
	time.Sleep(1 * time.Second)
	fmt.Println(s.DataCount)
}
