// Go语言reflect.TypeOf()和reflect.Type（通过反射获取类型信息）
// c.biancheng.net/view/109.html
package main

import (
	"fmt"
	"reflect"
)

// 可以为一个任何非接口类型的值创建一个reflect.Type类型的值。

// 那如何得到一个表示着某个接口类型的reflect.Type值呢？
//类型reflect.Type为一个接口类型，它指定了若干方法（golang.google.cn/pkg/reflect/#Type）
//通过这些方法，我们能观察到一个reflect.Type值所表示的Go类型的各种信息。
//这些方法中的适用于所有种类（golang.google.cn/pkg/reflect/#Kind）的类型，有的只适用于一种或几种类型
//通过不合适的reflect.Type属主值调用某个方法将在运行时产生一个恐慌。

// 使用reflect.TypeOf()函数可以获得任意值的类型对象（reflect.Type），程序通过类型对象可以访问任意值的类型信息。

/*
func main() {
	var a int
	typeOfA := reflect.TypeOf(a) // 取a的反射类型对象给typeOfA，类型为reflect.Type()
	valueOfA := reflect.ValueOf(a)
	fmt.Println(typeOfA.Name(), typeOfA.Kind()) // 通过typeOfA类型对象的成员函数，获取typeOfA的类型名（Name）为int（该Type的Name是int），种类（Kind）为int（该Type的Kind是int）
	fmt.Println(valueOfA.Interface())
}


*/
// 理解反射的类型（Type）与种类（Kind）
//当需要区分一个大品种的类型时，就会用到种类（Kind），例如：需要统一判断类型中的指针时，使用种类（Kind）信息就较为方便。

//1. 反射种类（kind）的定义
//Go 程序中的类型（Type）指的是系统原生数据类型，如int，string，bool，float32等类型
//以及type关键字定义的类型，这些类型的名称就是其类型本身的名称。例如：使用type A struct{} ，A 就是structural{}的类型
//种类（Kind）指的是对象归属的品种，在reflect包中如下定义：
//type Kind uint
//const (
//    Invalid Kind = iota  // 非法类型
//    Bool                 // 布尔型
//    Int                  // 有符号整型
//    Int8                 // 有符号8位整型
//    Int16                // 有符号16位整型
//    Int32                // 有符号32位整型
//    Int64                // 有符号64位整型
//    Uint                 // 无符号整型
//    Uint8                // 无符号8位整型
//    Uint16               // 无符号16位整型
//    Uint32               // 无符号32位整型
//    Uint64               // 无符号64位整型
//    Uintptr              // 指针
//    Float32              // 单精度浮点数
//    Float64              // 双精度浮点数
//    Complex64            // 64位复数类型
//    Complex128           // 128位复数类型
//    Array                // 数组
//    Chan                 // 通道
//    Func                 // 函数
//    Interface            // 接口
//    Map                  // 映射
//    Ptr                  // 指针
//    Slice                // 切片
//    String               // 字符串
//    Struct               // 结构体
//    UnsafePointer        // 底层指针
//)
//Map, Slice, Chan 属于引用类型，使用起来类似于指针，但是在种类常量定义中任然属于独立的种类，不属于Ptr。

//type A struct{} 定义的结构体， A反射种类（Kind）属于Struct种类，*A反射种类（Kind）属于Ptr

//2. 从类型对象中获取类型名称和种类的例子
//go 语言中的类型名称对应的反射获取方法是reflect.Type中的Name()方法，返回表示类型（type）名称的字符串
//类型归属的种类（Kind）使用的是reflect.Type中的Kind()方法，返回reflect.Kind类型的常量。

// 定义一个Enum类型
type Enum int

const (
	Zero Enum = 0 //Zero是Enum类型的常量
)

func main() {
	// 声明一个空结构体
	type cat struct { // 声明结构体cat
	}
	// 获取结构体实例的反射类型对象
	typeOfCat := reflect.TypeOf(cat{}) //将cat实例化，并获取cat的反射类型对象

	// 显示反射类型对象的名称和种类
	fmt.Println(typeOfCat.Name(), typeOfCat.Kind()) // 输出cat反射类型（type）的名是cat，和反射类型的种类（kind）是struct

	// 获取Zero常量 的反射类型对象
	typeOfA := reflect.TypeOf(Zero)
	// 显示反射类型对象的名称和种类
	fmt.Println(typeOfA.Name(), typeOfA.Kind()) // 输出Zero对应的类型对象的type名name（是Enum）和种类kind（是int）。
}
