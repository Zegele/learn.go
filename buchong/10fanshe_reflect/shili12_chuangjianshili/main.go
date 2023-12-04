// Go语言通过类型信息创建实例
// c.biancheng.net/view/117.html
package main

import (
	"fmt"
	"reflect"
)

// 当已知reflect.Type时，可以动态地创建这个类型的实例，
//实例的类型为指针。例如:reflect.Type的类型为int时，
//创建int的指针，即*int，代码如下：.

func main() {
	var a int = 1
	// 取变量a的反射类型对象
	typeOfA := reflect.TypeOf(a)
	fmt.Println("a's kind:", typeOfA.Kind())
	fmt.Println("a's Interface:", reflect.ValueOf(a).Interface())

	//根据反射类型对象创建类型实例
	aIns := reflect.New(typeOfA) // 使用reflect.New()函数传入变量a的反射类型对象，创建这个类型的实例值
	//值以reflect.Value类型返回。这步操作等效于：new(int)  ,因此返回的是*int类型的实例。
	fmt.Println(aIns.Elem())
	//输出Value的类型和种类
	fmt.Println(aIns.Type(), aIns.Kind()) // *int ptr
}

/*
//reflect.New(typeOfA)创建实例后赋值。
func main() {
	var A int64 = 1
	// 取变量a的反射类型对象
	typeOfA := reflect.TypeOf(&A)

	//根据反射类型对象创建类型实例
	aIns := reflect.New(typeOfA)

	aIns = reflect.ValueOf(&A)
	fmt.Println(aIns.Elem())

	aIns.Elem().SetInt(100)
	fmt.Println(aIns.Elem())
	//输出Value的类型和种类
	//fmt.Println(aIns.Type(), aIns.Kind())
}

*/
