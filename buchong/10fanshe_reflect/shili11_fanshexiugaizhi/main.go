// go语言通过反射修改变量的值
// c.biancheng.net/view/116.html
package main

import (
	"fmt"
	"reflect"
)

//go语言中类似x , x.f[1] , *p 形式 的表达式都可以表示变量， 但是其他如x+1 和 f(2)则不是变量
//一个变量就是一个可寻址的内存空间，里面存储了一个值，并且存储的值可以通过内存地址来更新
//
//对于reflect.Values也有类似的区别。有一些reflect.Values是可取地址的； 其他一些则不可以
//x := 2 //value type variable?
//a := reflect.ValueOf(2) // 2 int no
//b := reflect.ValueOf(x) // 2 int no
//c := reflect.ValueOf(&x) // &x *int no
//d := c.Elem() // 2 int yes (x)
//a对应的变量则不可取地址，因为a中的值仅仅只是2的拷贝副本
//b中的值也同样不可取地址，是x的拷贝
//c中的值还是不可取地址，它只是一个指针&x的拷贝。
//实际上，所有通过reflect.ValueOf(x)返回的reflect.Value都是不可取地址的
//但是对于d，它是c的解引用方式生成的，指向另一个变量，因此是可取地址的
//可以通过调用reflect.ValueOf(&x).Elem()，来获取任意变量x（除了interface）对应的可取地址的Value
//
//我们可以通过调用reflect.Value的CanAddr方法来判断其是否可以被取地址：
//fmt.Println(a.CanAddr()) //"false".
//fmt.Println(b.CanAddr()) //"false".
//fmt.Println(c.CanAddr()) //"false".
//fmt.Println(d.CanAddr()) //"true"
//
//每当我们通过指针间接地获取的reflect.Value都是可取地址的
//即使开始的是一个不可取地址的Value
//在反射机制中，所有关于是否支持取地址的规则都是类似的。
//例如：slice的索引表达式e[i]将隐式得包含一个指针，它就是可取地址的，即使开始的e表达式不支持也没有关系
//
//以此类推，reflect.ValueOf(e).Index(i)对于该值也是可取地址的，即使原始的reflect.ValueOf(e)不支持也没关系
//
//使用reflect.Value对包装的值进行修改时，需要遵循一些规则。
//如果没有按照规则进行代码设计和编写，轻则无法修改对象值，重则程序在运行时会发生宕机。

//判定及获取元素的相关方法
//使用reflect.Value取元素、取地址及修改值的方法，参考下表：
//Elem()Value //取reflect.Value指向的元素值，类似于语言层* 操作。当值类型不是指针或接口时发生宕机，空指针时返回nil的Value
//Addr()Value //对可寻址的值返回其地址，类似于语言层&操作。当值不可寻址时发生宕机
//CanAdd()bool //表示值是否可寻址
//CanSet()bool //返回值能否被修改，要求值可寻址且时导出的字段 // 导出的字段？？？
//
//值修改相关方法
//使用reflect.Value修改值的相关方法：
//SetInt(x int64) // 使用int64设置值。当值不是int，int8，int16，int32，int64时发生宕机
//SetUint(x uint64) //
//SetFloat(x float64) //
//SetBytes(x []byte) //
//SetString(x string) //
//以上方法，在reflect.Value的CanSet返回false，任然修改值时会发生宕机
//在已知值的类型时，应尽量使用值对应类型的反射设置值
//
//
/*
//1. 值可修改条件之一：可被寻址
//通过反射修改变量值的前提条件之一：这个值必须可以被寻址
//简单说就是这个变量必须能被修改
func main() {
	// 声明整型变量a并赋初始值
	var a int = 1024
	fmt.Println(a)
	// 获取变量a的反射值对象
	//valueOfA := reflect.ValueOf(a)
	// 尝试将a修改为1（此处会发生崩溃）
	//valueOfA.SetInt(1)

	// 修改后
	valueOfA := reflect.ValueOf(&a) // 这里是可寻址的
	//valueOfA.SetInt(1)              // 也是错的
	valueOfA.Elem().SetInt(1) // 这样才行
	//valueOfA.Elem()取出a地址对应的值， SetInt修改值
	fmt.Println(a)

	fmt.Println("valueOfA's addr:", valueOfA.CanAddr())
	fmt.Println("valueOfA's Elem addr:", valueOfA.Elem().CanAddr())
	//Elem才是可以寻址的
	// 当reflect.Value不可寻址时，使用Addr()方法也时无法取到值的地址的，同时会发生宕机。
	// 虽然说reflect.Value的Addr()方法类似于语言层的&操作；
	//Elem()方法类似于语言层*操作，但并不代表这些方法于语言层操作等效.
}

*/

// 2. 值可修改条件之一： 被导出
//结构体成员中，如果字段没有被导出，即便不使用反射也可以被访问，但不能通过反射修改：
//.
/*
func main() {
	type dog struct {
		legCount int // 小写不可导出 大写才可导出
	}
	// 获取dog实例的反射值对象
	valueOfDog := reflect.ValueOf(dog{legCount: 4})

	// 获取legCount字段的值
	vLegCount := valueOfDog.FieldByName("legCount")
	fmt.Println(vLegCount)

	// 尝试设置legCount的值（这里会发生崩溃）
	//valueOfDog.SetInt(40)
	// panic:reflect:reflect.Value.SetInt using value obtained using unexported field
	// SetInt()使用的值来自于一个未导出的字段

}
*/
//为了能修改这个值，需要将该字段导出
//将dog中的legCount的成员首字母大写，导出LegCount让反射可以访问，修改如果下：
// type dog struct{
// LegCount int
// }
//然后，并且满足条件一，取结构体的指针，再通过reflect.Value的Elem()方法取到值的反射值对象。
//如下：

func main() {
	type dog struct { // 如果dog类型在函数外，也要可导入才能用，也就是要用大写Dog才行。
		LegCount int //首字母大写，可导出字段
	}
	// 获取dog实例地址的反射值对象
	valueOfDog := reflect.ValueOf(&dog{LegCount: 4}) // 获取dog实例指针的反射值对象

	// 取出dog实例地址的元素
	valueOfDogElem := valueOfDog.Elem() // 取dog实例的指针元素，也就是dog的实例

	// 获取legCount字段的值
	vLegCount := valueOfDogElem.FieldByName("LegCount") // 取dog结构体中LegCount字段的成员值
	fmt.Println(vLegCount)

	// 尝试设置legCount的值
	vLegCount.SetInt(40) // 修改该成员值

	fmt.Println(vLegCount.Int()) // Value.Int()int64
}
