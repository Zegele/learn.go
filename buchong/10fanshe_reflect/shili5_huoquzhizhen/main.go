// Go语言reflect.Elem() ———— 通过反射获取指针指向的元素类型
// c.biancheng.net/view/110.html
package main

import (
	"fmt"
	"reflect"
)

// Go语言程序中对指针获取反射对象后，可以通过reflect.Elem()方法获取这个指针指向的元素类型。
// 这个获取过程被称为取元素，等效于对指针类型变量做了一个*操作。如下：
func main() {
	// 声明一个空结构体
	type cat struct {
	}
	// 创建cat的实例
	ins := cat{}
	insptr := &cat{} // 创建了cat结构体的实例，insptr是一个*cat（cat指针）类型的指针变量。

	// 获取结构体实例的反射类型对象
	typeOfCat := reflect.TypeOf(ins)
	typeOfCatPtr := reflect.TypeOf(insptr) // 指针变量获取反射类型信息

	// 显示反射类型对象的名称和种类
	fmt.Printf("name:'%v' kind:'%v'\n", typeOfCat.Name(), typeOfCat.Kind())
	fmt.Printf("name:'%v' kind:'%v'\n", typeOfCatPtr.Name(), typeOfCatPtr.Kind())
	// 输出指针变量的类型名称和种类。
	//go语言的反射中对所有指针变量的种类都是Ptr，但注意指针变量的类型名称是空，不是*cat

	// 取类型的元素
	typeOfCatElem := typeOfCatPtr.Elem()
	// 取指针类型的元素类型，也就是cat类型。这个操作不可逆，不可以通过一个非指针类型获取它的指针类型。

	// 显示反射类型对象的名称和种类
	fmt.Printf("element name:'%v', element kind:'%v'\n", typeOfCatElem.Name(), typeOfCatElem.Kind())
	// 输出指针变量指向元素的类型名称和种类，得到了cat的类型名称（cat）和种类（struct）
}
