package main

import (
	"fmt"
	"sync"
	"time"
)

type Store struct {
	DataCount int
	Max       int
	lock      sync.Mutex
	pCond     *sync.Cond //Cond 是干嘛的？没印象
	cCond     *sync.Cond
}

type Producer struct{}

func (Producer) Produce(s *Store) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.DataCount == s.Max {
		fmt.Println("生产者在等仓库拉走货")
		s.pCond.Wait()
	}
	fmt.Println("开始生产+1")
	s.DataCount++
	s.cCond.Signal() //Signal发送后，接受这个Signal的锁就打开了。
}

type Consumer struct{}

func (Consumer) Consume(s *Store) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.DataCount == 0 {
		fmt.Println("消费者在等货")
		s.cCond.Wait()
	}
	fmt.Println("消费者消费-1")
	s.DataCount--
	s.pCond.Signal()
}

func main() {
	s := &Store{
		Max: 10,
	}
	s.pCond = sync.NewCond(&s.lock) //指针类型初始化 该condition是由系统的sync包中的NewCond函数初始化得来的。
	s.cCond = sync.NewCond(&s.lock) //NewCond的参数是个接口，实现了该接口的对象，都可以放在这个参数的位置上？

	pCount, cCount := 50, 50
	for i := 0; i < pCount; i++ {
		go func() {
			for {
				time.Sleep(100 * time.Millisecond)
				Producer{}.Produce(s)
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
