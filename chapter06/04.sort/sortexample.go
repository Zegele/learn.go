package main

// 从大到小，或从小到大排序
import (
	"fmt"
	"sort"
)

type Button struct {
	Floor int
}
type Buttons []*Button //指针类型的结构体 切片类型

type Elevator struct {
	buttons  Buttons //切片  结构体嵌套了一个结构体类型的切片
	position int
}

func (b Buttons) Len() int {
	return len(b)
}

func (b Buttons) Less(i, j int) bool {
	return b[i].Floor < b[j].Floor // 左边比右边小，则不做交换（不调用swap）。否则，做交换（调用下面的swap）。所以出来的是从小到大。
	// 改变成 > 号，则是左边比右边大，则不做交换。否则，做交换。所以出来的是从大到小。
}

func (b Buttons) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func main() {
	ev := &Elevator{
		position: 2,
		buttons: Buttons{
			{Floor: 3},
			{Floor: 1},
			{Floor: 5},
			{Floor: 2},
			{Floor: 4},
		},
	}
	sort.Sort(ev.buttons)
	for _, item := range ev.buttons {
		fmt.Println(item.Floor)
	}

	sort.Sort(sort.Reverse(ev.buttons)) // sort.Reverse()函数是反转，返回值是Interface

	fmt.Printf("%+v\n", ev.buttons)
	for _, item := range ev.buttons {
		fmt.Println(item.Floor)
	}
}
