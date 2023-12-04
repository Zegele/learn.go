// 反射规则浅析
// c.biancheng.net/view/5131.html
package main

import (
	"fmt"
	"reflect"
)

/*
// 回顾接口类型
//接口是 一个重要的类型，它意味着一个确定的方法集合，一个接口变量可以存储任何实现了接口的方法的具体值（除了接口本身）
// 例如：io.Reader 和 io.Writer
// 如果一个类型声明实现了Reader（或Writer）方法，那么它便实现了io.Reader (或 io.Writer)
// 这意味着一个io.Reader的变量可以持有任何一个实现了Read方法的类型的值.如下：
// Reader is the interface that wraps the basic Read method.
type Reader interface {
	Read(p []byte) (n int, err error)
}

var r io.Reader// r可以接住 任何实现了Reader接口中方法（read函数）的值。
r = os.Stdin
r = bufio.NewReader(r)
r = new(bytes.Buffer)

// and so on
// 必须要弄清楚的一点是，不管变量r中的具体值是什么，r的类型永远是io.Reader，
//由于Go语言是静态类型的，r的静态类型就是io.Reader
//在运行时，接口变量的类型是不会变.


*/

//*** 反射三定律
// 反射第一定律：反射可以将“接口类型变量”转换为“反射类型对象”
//注：这里反射类型指reflec.Type 和 reflect.Value。

/*
func main() {
	var x float64 = 3.4
	fmt.Println("type:", reflect.TypeOf(x))
	fmt.Println("type:", reflect.ValueOf(x))
	// TypeOf returns the reflection Type of the value in the interface{}
	// func TypeOf(i interface{})Type
	// 我们调用reflect.TypeOf(x)时，x被存储在一个空接口变量中传递过去，然后reflect.TypeOf对空接口变量进行拆解，恢复其类型信息。

}

*/

/*
// 类型reflect.Type和reflect.Value都有很多方法，举几个例子：
// 类型 reflect.Value由一个方法Type(), 它会返回一个reflect.Type类型的对象。
// Type和Value都有一个kind的方法，会返回一个常量，表示底层数据的类型，常见：Uint，Float64，Slice等

// Value 类型也有一些类似于Int， Float的方法，用来提取底层的数据：
//Int方法用来提取int64
//	Float方法用来提取float64,实例如下：

func main() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	y := reflect.TypeOf(x)
	fmt.Println("type:", v.Type())
	fmt.Println("v.Type==y?", v.Type() == y)
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64) //reflect.Float64表示反射类型中的Float64类型
	fmt.Println("value:", v.Float())//v.Float() 取float64的值
	fmt.Println("value:", v)
}

*/

/*
// 可修改性
//reflect.Value类型的getter和setter方法，为了保证API的精简，这两个方法操作的时某一组类型范围最大的那个。比如：处理任何含号整型数，都是int64
//也就是说，Value类型的Int()方法返回值为int64类型，SetInt方法接收的参数类型也是int64类型.
func main() {
	var x uint8 = 'x'
	fmt.Println(x)

	v := reflect.ValueOf(x)
	fmt.Println("--------", v)
	fmt.Println("type:", v.Type())                           //uint8
	fmt.Println("kind is uint8:", v.Kind() == reflect.Uint8) //true
	// Kind 方法描述的是基础类型，而不是静态类型
	xx := v.Uint() // xx是int64类型
	x = uint8(xx)  //

	type MyInt int
	var a MyInt = 7
	valueOfa := reflect.ValueOf(a)
	fmt.Printf("vv‘s kind:%s. vv's type:%s\n", valueOfa.Kind(), valueOfa.Type())
	// 注意区分Kind 和 Type 理解这两个东西。
	typeOfa := reflect.TypeOf(a)
	fmt.Println("vv‘s kind:", typeOfa.Kind(), "  vv‘s type name:", typeOfa.Name())
}

*/

/*
// 反射第二定律
//反射可以将“反射类型对象”转换为“接口类型变量”
//根据一个reflect.Value类型的变量，我们可以使用Interface方法恢复其接口类型的值。
//事实上，这个方法会把type和value信息打包并填充到一个接口变量中，然后返回
//Interface returns v's value as an interface{}
//func (v Value)Interface()interface{}
//然后，我们可以通过断言，恢复底层的具体值
//y:=v.Interface().(float64)//y will have type float64
//fmt.Println(y)
//上面会打印出一个float64类型的值，也就是反射类型变量v所代表的值。
// 利用fmt包
//fmt.Println(v.Interface())
//为什么不直接使用fmt.Println(v)?因为v的类型是reflect.Value，我们需要的是它的具体值，由于值的类型是float64，我们也可以用浮点格式打印它。
// fmt.Printf("value is %7.1e\n", v.Interface())
func main() {
	var a float32 = 3.912345
	v := reflect.ValueOf(a)
	fmt.Println(v)
	vfloat64 := v.Float()
	fmt.Printf("%T\n", vfloat64)
	vv := v.Interface().(float32) //vv will have type float64.
	vvv := v.Interface()
	fmt.Println("vv", vv)
	fmt.Println("vvv:", vvv)
	fmt.Printf("%T\n", vv)
	fmt.Printf("value is %10.1e\n", v.Interface()) // %10.1e 里面的10是什么意思？代表距离？？应该是。  .1代表保留1位小数； e表示小数部分使用科学计数法
	fmt.Printf("value is %20.4f\n", v.Interface()) //可参考 go语言的格式化打印
}

*/

//反射第三定律
//如果要修改“反射类型对象”其值必须是“可写的”
//var x float64 = 3.4
// v := reflect.ValueOf(x)
// v.SetFloat(7.1) // Error: will panic
//运行会报错：panic:reflect:reflect.flag.mustBeAssignable using unaddressable value
//这里问题不在于值7.1不能被寻址，而是因为变量v是“不可写的”，“可写性”是反射类型变量的一个属性，但不是所有的反射类型变量都拥有这个属性
//可以通过CanSet方法检查一个reflect.Value类型变量的“可写性”。
/*
func main() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	// 这里传递给reflect.ValueOf函数的是变量x的一个拷贝，而非x本身，
	//所以 执行 v.SetFloat(7.1) 会报错。 因为修改的值对x没影响
	//如果想通过反射修改变量x，就要把想要修改的变量的指针传递给反射库。
	fmt.Println("settability of v:", v.CanSet()) //false 所以v不可写
	// 所以对于一个不具有“可写性”的Value类型变量，调用Set方法会报错。

}

// 什么是可写性？
// 可写性有类似于寻址能力，但是更严格，它是反射类型变量的一种属性
//赋予该变量修改底层存储数据的能力
//可写性最终是由一个反射对象是否存储了原始值而决定的
//

*/

////如果想通过反射修改变量x，就要把想要修改的变量的指针传递给反射库。.
/*
func main() {
	var x float64 = 3.4
	p := reflect.ValueOf(&x) // Note: take the address of x
	fmt.Println("type of p:", p.Type(), "  kind of p:", p.Kind())
	fmt.Println("settability of p:", p.CanSet()) //false 还是不可写
	// 反射对象p是不可写的，但是我们也不想修改p，事实上我们要修改的是*p
// 继续看下面例子：
}
*/

/*
//为了得到p指向的数据，可以调用Value类型的Elem方法。
//Elem方法能够对指针进行“解引用，然后将结果存储到反射Value类型对象v中。
func main() {
	var x float64 = 3.14
	p := reflect.ValueOf(&x) // Note: take the address of x.
	v := p.Elem()            // 注意这里是reflect.value的Elem()方法 区分reflect.type的Elem()方法
	fmt.Println("改值前：", v.Interface())
	fmt.Println("settability of v:", v.CanSet()) //true 可以改值
	v.SetFloat(7.14)
	fmt.Println("改值后：", v.Interface()) //为什么不直接fmt.Println(v) 因为v是reflect.Value类型。
	fmt.Println(x)
}
//反射不容易理解，reflect.Type 和 reflect.Value 会混淆
//但是它做的事情正是编程语言做的事情。
//只记住：只要反射对象要修改它们表示的对象，就必须获取它们表示的对象的地址。

*/

//结构体：
//一般使用反射修改结构体的字段，只要有结构体的指针，我们就可以修改它的字段
//下面是一个解析结构体变量t的例子，用结构体的地址创建反射变量，再修改它。
//然后对它的类型设置了typeOfT，并用调用简单的方法迭代字段。
//注意：我们从结构体的类型中提取了字段的名字，但每个字段本身是正常的reflect.Value对象.

func main() {
	type T struct { // 注意 T 以及 T的字段名 首字母都要大写，因为结构体中只有可导出的字段是”可设置“的。
		ID   int
		Name string
	}
	t := T{23, "dgaldkjg"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}

	fmt.Println("t is before:", t)
	s.Field(0).SetInt(1000)
	s.Field(1).SetString("GO GO GO")
	fmt.Println("t is after:", t)
}

// 总结：
// 1. 反射可以将”接口类型变量“转换为”反射类型对象“
//2. 反射可以将”反射类型对象“转换为”接口类型变量“
//3. 如果要修改”反射类型对象“，其值必须是”可写的“。
