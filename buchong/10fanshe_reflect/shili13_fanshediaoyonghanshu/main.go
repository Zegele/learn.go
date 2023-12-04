//Go语言通过反射调用函数
//c.biancheng.net/view/118.html

package main

import (
	"fmt"
	"reflect"
)

// 如果反射值对象（reflect.Value）中值的类型为函数时，可以通过reflect.Value调用该函数
//使用反射调用函数时，需要将参数使用反射值对象的切片[]reflect.Value 构造后传入Call()方法中，
//调用完成时，函数的返回值通过[]reflect.Value返回
//
//下面的代码声明一个加法函数，传入两个整型值，返回两个整型值的和。
//将函数保存到反射值对象（refelct.Value）中，然后将两个整型值构造为反射值对象的切片（[]reflect.Value）
//使用Call()方法进行调用.

// 普通函数
func add(a, b int) int {
	return a + b
}

func main() {

	// 将函数包装为反射值对象
	funcValue := reflect.ValueOf(add)

	// 构造函数参数，传入两个整型值
	paramList := []reflect.Value{reflect.ValueOf(10), reflect.ValueOf(20)}
	// 将add需要的参数，构造成[]reflect.Value切片，且本身也要是reflect.Value类型

	// 反射调用函数
	retList := funcValue.Call(paramList) // 通过Call()函数，传入构造的[]reflect.Value切片
	// 返回值也是个切片，[]Value

	// 获取第一个返回值，取整数值
	fmt.Println(retList[0])       // 通过retList[0]取返回值的第一个参数，
	fmt.Println(retList[0].Int()) // 通过retList[0]取返回值的第一个参数，
	// 并使用Int取返回值的整数型值。
}

// 提示：
// 反射调用函数的过程需要构造大量的reflect.Value和中间变量，对函数参数值进行逐一检查，
//还需要将调用参数复制到调用函数的参数内存中
//调用完毕后，还需要将返回值转换为refelct.Value
//用户还需要从中取出调用值
//因此，反射调用函数的性能问题尤为突出，不建议大量使用反射函数调用.
