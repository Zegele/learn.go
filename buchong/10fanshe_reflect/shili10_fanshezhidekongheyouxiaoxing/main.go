// Go语言IsNil()和IsValid()————判断反射值的空和有效性
// c.biancheng.net/view/115.html
package main

import (
	"fmt"
	"reflect"
)

// 反射值对象（reflect.Value）提供一系列方法进行零值和空判定，如下：
// IsNil()bool //返回值是否为nil。
// 如果值类型不是通道（channel），函数，接口，map，指针或切片时发生panic，类似于语言层的 v == nil操作 。
// 没理解上一句
// IsValid()bool //判断值是否有效。当值本身非法时，返回false，例如reflect Value不包含任何值，值为nil等
//
// 下例，将会对各种方式的空指针进行IsNil()和IsValid()的返回值判定检测。
// 同时对结构体成员及方法查找map键值对的返回值进行IsValid()判定：.
func main() {
	// *int的空指针
	var a *int
	fmt.Printf("%v,%+v,%#v\n", a, a, a) //(*int)(nil) 的含义是将nil转换为*int，也就是*int类型的空指针。
	// %v 默认格式输出； %+v 在%v基础上额外输出字段名； %#v 在%+v基础上，额外输出类型名。
	var b int = 100
	fmt.Println("var a *int:", reflect.ValueOf(a).IsNil())

	// nil值
	fmt.Println("nil:", reflect.ValueOf(nil).IsValid())

	// *int类型的空指针
	fmt.Println("(*int)(nil):", reflect.ValueOf(a).Elem().IsValid())
	fmt.Println("(*int)(nil):", reflect.ValueOf((*int)(nil)).Elem().IsValid())
	//(*int)(nil) 的含义是将nil转换为*int，也就是*int类型的空指针。  Elem()是取指针指向的元素
	//(*int)nil 意思是指针类型的int是nil空。

	a = &b
	fmt.Println(*a)
	fmt.Println("(*int)(nil):", reflect.ValueOf(a).Elem().IsValid())
	fmt.Println("(*int)(nil):", reflect.ValueOf(a).IsValid())

	// 实例化一个结构体
	s := struct{}{}

	// 尝试从结构体中查找一个不存在的字段
	fmt.Println("不存在的结构体成员：", reflect.ValueOf(s).FieldByName("").IsValid())
	// 查找结构体中一个空字符串的成员，如果成员不存在，IsValid()返回false

	// 尝试从结构体中查找一个不存在的方法
	fmt.Println("不存在的结构体方法：", reflect.ValueOf(s).MethodByName("").IsValid())
	// 查找结构体中一个空字符串的方法，如方法不存在，IsValue()返回false

	// 实例化一个map
	m := map[int]int{}

	// 尝试从map中查找一个不存在的键
	fmt.Println(reflect.ValueOf(m).IsNil())
	fmt.Println("不存在的键：", reflect.ValueOf(m).MapIndex(reflect.ValueOf(3)).IsValid())
	// MapIndex()方法能根据给定的reflect.Value类型的值查找map，并且返回查找到结果。
}

//！！！注意：
//InNil() 常被用于判断指针是否为空； IsValid()常被用于判定返回值是否有效。
