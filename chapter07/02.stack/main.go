package main

import "fmt"

// 栈： 先进后出
type Stack struct {
	data []interface{}
}

func (s *Stack) Push(data interface{}) {
	s.data = append([]interface{}{data}, s.data...) //和队列不同的是，栈是给data装入s.data，队列是给s.data装入data
}

func (s *Stack) Pop() (interface{}, bool) {
	if len(s.data) > 0 {
		o := s.data[0]
		s.data = s.data[1:]
		return o, true
	}
	return nil, false
}

func main() {
	s := &Stack{}
	s.Push(1)
	s.Push(2)
	s.Push(3)

	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())

	//Add(1, 2) //stack overflow	展示堆栈溢出

}

//func Add(a, b int) {
//	Add(a, b)递归调用，且没有设置边界
//}
