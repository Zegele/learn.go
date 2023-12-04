// Go语言refelct.ValueOf()和reflect.Value（通过反射获取值信息）
// c.biancheng.net/view/113.html
package main

import (
	"fmt"
	"reflect"
)

// 当我们将一个接口值传递给一个reflect.ValueOf函数调用时，此调用返回的是代表着此接口值的动态值的一个reflect.Value值。
//我们必须通过间接的图途径获得一个代表一个接口值的refelct.Value值。
//reflect.Value类型有很多方法（golang.google.cn/pkg/reflect/）我们可以调用这些方法来观察和操作一个reflect.Value属主值表示的Go值。
//这些方法中的有些适用于所有种类类型的值
//
//通过不合适的reflect.Value属主值调用某个方法将在运行时产生一个恐慌。
//请阅读reflect代码库中各个方法的文档来获取如何正确地使用这些方法
//
//一个reflect.Value值的CanSet方法将返回此reflect.Value值代表的Go值是否可以被修改（可以被赋值）
//如果一个go值可以被修改，则我们可以调用对应的reflect.Value值的Set方法来修改此go值。
//注意：reflect.ValueOf 函数直接返回的reflect.Value值都是不可修改的
//
//反射不仅可以获取值的类型信息，还可以动态地获取或者设置变量的值
//go语言中使用reflect.Value获取和设置变量的值
//
//使用反射值对象包装任意值
//Go语言中，使用refeclt.ValueOf()函数获得值的反射值对象（reflect.Value）
//value := reflect.ValueOf(rawValue)
//reflect.ValueOf返回reflect.Value类型，包含由rawValue的值信息
//reflect.Value与原值间可以通过值包装和值获取相互转化。
//reflect.Value是一些反射操作的重要类型，如反射调用函数
//
//
//从反射值对象获取被包装的值
//go语言中更可以通过reflect.Value重新获得原始值
//1. 从反射值对象（reflect.Value）中获取值的方法
//可以通过下面几种方法从反射值对象reflect.Value中获取原值，如下表所示：
//Interface()interface{}  //将值以interface{}类型返回，可以通过类型断言转换为指定类型
//Int()int64  // 以int类型返回，所有有符号整型均可以此方式返回
//Uint()uint64  //
//Float()float64  //
//Bool()bool  //
//Bytes()[]bytes  // 将值以字节数组类型返回
//String()string  //
//
//2. 从反射值对象（reflect.Value）中获取值的例子
//下面代码，将整型变量中的值使用reflect.ValueOf()获取反射值对象（reflect.Value）
//再通过relect.Value的Interface()方法获得interface{}类型的原值，
//通过int类型对应的refelct.Value的Int()方法获得整型值。

func main() {
	// 声明整型变量a并赋初始值
	var a int = 1024
	// 获取变量a的反射值对象
	valueOfA := reflect.ValueOf(a) //获取变量a的反射值对象，类型为reflect.Value
	v := valueOfA.Interface()
	fmt.Printf("v's value: '%v', v's type:'%v', v's kind:'%v'\n", v, valueOfA.Type(), valueOfA.Kind())
	// 获取interface{}类型的值，通过类型断言转换
	var getA int = valueOfA.Interface().(int) //接口类型断言，参考 // ：c.biancheng.net/view/83.html

	//获取64位的值，强制类型转换为int类型
	var getA2 int = int(valueOfA.Int())

	fmt.Println(getA, getA2)
}
