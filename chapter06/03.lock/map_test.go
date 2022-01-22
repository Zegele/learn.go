package main

import (
	"fmt"
	"testing"
	"time"
)

//map遇到goroutine一定要小心，就是这个错。map不能同步 写
//fatal error: concurrent map read and map write (map的读和写，同步了。)
func TestMap(t *testing.T) {
	m := map[int]int{}
	for i := 0; i < 100; i++ {
		go func() {
			for {
				v := m[i]
				fmt.Println(v)
				m[i] = v + 1
				fmt.Println("i = ", i, m[i])
			}
		}()
	}
	time.Sleep(10 * time.Second)
}
