package main

import "fmt"

func main() {
	a := []int{} //这是slice 可添加元素
	fmt.Println(a)
	fmt.Println("追加元素到a中，a是切片")
	a = append(a, 333)
	fmt.Println(a)

	b := [0]int{} //这是arr 没有空间的arr，是没用的东西
	fmt.Println(b)
	fmt.Println("追加元素到b中，b是数组")
	//b = append(b, 333)// 编译错误，无法追加固定长度数组

	xqInfo := []string{"小强", "男", "在职"}
	fmt.Println(xqInfo)
	for i, v := range xqInfo {
		fmt.Println(i, v)
	}
	fmt.Println(xqInfo[0])

	fmt.Println("================")
	fmt.Println("删除切片中的元素")
	a = []int{111, 222, 333, 444, 555}
	fmt.Println("删除前：", a)
	a = append(a[:1], a[2:4]...) //todo ...是啥？
	fmt.Println("删除后", a)
	a = append(a, a[:]...)
	fmt.Println("double后：", a)

	backup := append([]int{}, a[1:]...)
	a = append(a[:1], 999)
	a = append(a, backup...)
	fmt.Println(a)

}
