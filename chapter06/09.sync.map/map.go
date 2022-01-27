package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	m := sync.Map{} //sync.Map{}是线程安全的
	for i := 0; i < 10; i++ {
		go func(i int) {
			m.Store(i, 1)
			for {
				v, ok := m.Load(i) //v 本质是个空接口 Load函数就是读取值
				if !ok {
					continue
				}
				m.Store(i, v.(int)+1) //v.(int) 把v转成int类型。 不能用int(v)这种方式转换成int类型
				//Store 就是存储值 i相当于map中的key，v.(int)+1 相当于map中的值
				fmt.Println("i=", v)
			}
		}(i)
	}
	time.Sleep(10 * time.Second)
}
