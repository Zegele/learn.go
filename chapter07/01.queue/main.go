package main

//先进先出 类似羽毛球桶
import (
	"fmt"
	"sync"
)

type Queue struct {
	sync.Mutex
	data []interface{}
}

func (q *Queue) Push(data interface{}) {
	q.Lock()
	defer q.Unlock()
	q.data = append(q.data, data)
}

func (q *Queue) Pop() (interface{}, bool) {
	q.Lock()
	defer q.Unlock()
	if len(q.data) > 0 {
		o := q.data[0]
		q.data = q.data[1:]
		return o, true //每次切出slice的第一个元素
	}
	return nil, false //slice中没东西了
}

func main() {
	q := &Queue{}
	q.Push(111)
	q.Push("a")
	q.Push(3)
	q.Push(4)
	q.Push(5)
	q.Push(nil)

	fmt.Println(q.Pop()) //222 true
	fmt.Println(q.Pop()) //a true
	fmt.Println(q.Pop()) //3 true
	fmt.Println(q.Pop()) //4 true
	fmt.Println(q.Pop()) //5 true
	fmt.Println(q.Pop()) //<nil> true
	fmt.Println(q.Pop()) //<nil> false
}
