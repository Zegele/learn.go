// go语言inject库：依赖注入
// c.biancheng.net/view/5132.html
package main

import (
	"fmt"
	"github.com/codegangsta/inject"
)

//先了解“依赖注入”和“控制反转”
//控制反转
//正常状况下，对函数或方法的调用是我们的主动直接行为，在调用某个函数之前我们需要清楚地知道被调用函数的名称是什么，参数有哪些类型等
//所谓的控制反转就是将这种主动行为变成间接行为，我们不用直接调用函数或对象，
//而是借助框架代码进行间接的调用和初始化，这种行为称为“控制反转”，
//库和框架能很好的解释控制反转的概念
//
//依赖注入
//依赖注入是实现控制反转的一种方法，如果说控制反转是一种设计思想，那么依赖注入就是这种思想的一种实现
//通过注入参数或实例的方式实现控制反转。
//如果没有特殊说明，我们可以认为依赖注入和控制反转是一个东西
//
//控制反转的价值在于解耦，有了控制反转就不需要将代码写死，可以让控制反转的框架代码读取配置，动态的构建对象，这一点在Java和Spring框架中体现的有位突出.

//inject实践
//inject是依赖注入的Go语言实现，它能在运行时注入参数，调用方法，是Martini框架（Go语言中著名的Web框架）的基础核心
//在介绍具体实现之前，先来想一个问题，如何通过一个字符串型的函数名来调用函数？
//Go语言没有Java中的Class.forName方法可以通过类名直接构造对象，所以这种该方法是行不通的
//能想到的方法就是使用map实现一个字符串到函数的映射，如下：
/*
func f1(){
	println("f1")
}

func f2(){
	println("f2")
}

funcs := make(map[string]func())
funcs ["f1"] = f1
funcs ["f2"] = f2
funcs["f1"]()
funcs["f2"]()

*/

// 但是这有个缺陷，就是map的Value类型被写成func(),不同参数和返回值的类型的函数并不能通用
//将map的Value定义为interface{}空接口类型即可解决该问题
//但需要借助类型断言或反射来实现，通过类型断言实现等于又绕回去了，反射是一种可行的办法。

/*
type S1 interface{}
type S2 interface{}

func Format(name string, company S1, level S2, age int) {
	fmt.Printf("name = %s, company=%s, level=%s, age = %d\n", name, company, level, age)
}

func main() {
	// 控制实例的创建
	inj := inject.New()
	// 实参注入
	inj.Map("tom")
	inj.MapTo("tencent", (*S1)(nil)) // 给S1空接口写入“tencent”
	inj.MapTo("T4", (*S2)(nil)) // 给S2空接口写入“T4”  装入什么就对应是什么类型
	inj.Map(23)
	//函数反转调用
	inj.Invoke(Format)
// 运行结果： name = tom, company=tencent, level=T4, age = 23
}
*/
//可见，inject提供了一种注入参数调用函数的通用功能，inject.New()想当于创建了一个控制实例，
//由其来实现对函数的注入调用。
//inject包不但提供了对函数的注入，还实现了对struct类型的注入。如下：
//.
/*
type S1 interface{}
type S2 interface{}
type Staff struct {
	Name    string `inject`
	Company S1     `inject`
	Level   S2     `inject`
	Age     string `inject`
}

func main() {
	// 创建被注入实例
	s := Staff{}
	//控制实例的创建
	inj := inject.New()
	// 初始化注入值
	inj.Map("tom") //Map() , MapTo()用于注入参数
	inj.MapTo("T4", (*S2)(nil))
	inj.MapTo("tencent", (*S1)(nil))
	inj.Map("23")
	// 实现对struct注入
	inj.Apply(&s) // 怎么一一对应的？？？如果有两个string类型的，怎么确保正确对应？？
	//inject.Apply()用于注入结构体

	// 打印结果
	fmt.Printf("s = %v\n", s)
	// 运行结果：s = {tom tencent T4 23}

}

*/

//可以看到inject提供了一种对结构类型的通用注入方法。
//至此，我们仅仅从宏观层面了解inject能做什么，下面从源码实现角度分析inject

//inject原理分析
//inject 包中只有2个文件，一个是inject.go，一个inject_test.go
//inject.go短小精悍，包括注释和空行在内才157行，定义了4个接口，一个父接口，3个子接口
//type Injector interface{
//	Applicator
//	invoker
//	TypeMapper
//	SetParaent(Injector)
//}
//
//type Applicator interface{
//	Apply(interface{})error
//}
//
//type Invoker interface{
//	Invoke(interface{})([]reflect.Value, error)
//}
//
//type TypeMapper interface{
//	Map(interface{})TypeMapper // 返回值的类型还是自己？
//	MapTo(interface{}, interface{})TypeMapper
//	Get(reflect.Type)reflect.Value
//}
//Injector 接口是Applicator，Invoker，TypeMapper接口的父接口，所以实现Injector接口的类型，也必然实现了Applicator，Invoker和TypeMapper接口：
//1. Applicator接口只规定了Apply成员，它用于注入struct
//2. Invoker 接口只规定了Invoke成员，它用于执行被调用者
//3. TypeMapper接口规定了三个成员，Map和MapTo都用于注入参数，但他们有不同的用法，Get用于调用时获取被注入的参数
//
//另外，Injector 还规定了SetParent行为，它用于设置父Injector，其实它相当于查找继承
//也即通过Get方法在获取被注入参数时会一直追溯到parent，这是个递归过程，知道查找到参数或为nil终止
//
//type injector struct{
//	values map[reflect.Type]reflect.Value
//	parent Injector
//}
//
//func InterfaceOf(value interface{}) reflect.Type{
//	t := reflect.TypeOf(value)
//
//	for t.Kind() == reflect.Ptr{
//		t= t.Elem()
//	}
//
//	if t.Kind() != reflect.Interface{ // reflect.Interface是什么玩意儿？？？
//		panic("Called inject.InferfaceOf with a value that is not a pointer to an interface. (*MyInterface)(nil)")
//	}
//
//	return
//}
//
//func New() Injector {
//	return &injector{
//	values: make(map[reflect.Type]reflect.Value),
//	}
//}
//
//injector是inject包中唯一定义的struct，所有的操作都是基于injector struct来进行的
//他有两个成员values和parent。values用于保存注入的参数，是一个用reflect.Type当键，refect.Value为值的map
//理解这点将有助于理解Map和MapTo
//
//New方法用于初始化injector struct， 并返回一个指向injector struct的指针，
//但是这个返回值被Injector接口包装了
//
//interfaceOf方法虽然只有几句实现代码，但它是Injector的核心。
//InterfaceOf方法的参数必须是一个接口类型的指针，如果不是则引发panic
//InterfaceOf方法的返回类型是reflect.Type, injector的成员values就是一个reflect.Type类型当键的map
//这个方法的作用其实只是获取参数的类型，而不关心它的值
//
//.
/*
type SpecialString interface{}

func main() {
	fmt.Println(inject.InterfaceOf((*interface{})(nil)))
	fmt.Println(inject.InterfaceOf((*SpecialString)(nil)))
}

*/

// InterfaceOf方法就是用来得到参数类型，而不关心它具体存储的是什么值
//func (i *injector) Map(var interface{})TypeMapper{
//	i.values[reflect.TypeOf(val)] = refelct.ValueOf(val)
//	return i
//}
//
//func (i *injector) MapTo(val interface{}, ifacePtr interface{})TypeMapper{
//	i.values[InterfaceOf(ifacePtr)] = reflect.ValueOf(val)
//	return i
//}
//
//func (i *injector) Get(t reflect.Type)reflect.Value{
//	val := i.values[t]
//	if !val.IsValid() && i.parent != nil{
//		val = i.parent.Get(t)//递归
//	}
//	return val
//}
//
//func (i *injector) SetParent(parent Injector){
//	i.parent = parent
//}
//Map和MapTo方法都用于注入参数，保存于injector的成员values中。
//这两个方法的功能完全相同，唯一的区别就是Map方法用参数值本身的类型当键
//而MapTo方法有一个额外的参数可以指定特定的类型当键
//但是MapTo方法的第二个参数ifacePtr必须是接口类型指针，因为最终ifacePtr会作为InterfeOf方法的参数
//
//为什么需要有MapTo方法？
//因为注入的参数是存储在一个以类型为键的map中，可想而知，当一个函数中有一个以上的参数的类型是一样时，后执行Map进行注入的参数将会覆盖前一个通过Map注入的参数。
//
//SetParent方法用于给某个Injector指定父Injector。
//Get方法通过reflect.Type从injector的values成员中取出对应的值
//它可能会检查是否设置了parent
//直到找到或返回无效的值，
//最后Get方法的返回值会经过IsValid方法的校验

/*
type SpecialString interface{}

func main() {
	inj := inject.New()
	inj.Map("GOGOGO")
	inj.MapTo("Golang", (*SpecialString)(nil))
	inj.Map(20)
	fmt.Println("字符串是否有效？", inj.Get(reflect.TypeOf("G")).IsValid())
	fmt.Println("特殊字符串是否有效？", inj.Get(inject.InterfaceOf((*SpecialString)(nil))).IsValid())
	fmt.Println("[]byte是否有效：", inj.Get(reflect.TypeOf([]byte("Golang"))).IsValid())
	inj2 := inject.New()
	inj2.Map([]byte("test"))
	inj.SetParent(inj2)
	fmt.Println("[]byte是否有效?", inj.Get(reflect.TypeOf([]byte(""))).IsValid())

	fmt.Println(inj.MapTo(123, (*SpecialString)(nil)).Get(reflect.TypeOf(123)).Int())
	//没看懂这里为什么输出 20 ？？
	fmt.Println(inj.MapTo("golang", (*SpecialString)(nil)).Get(reflect.TypeOf("golang")).Interface())
	//没看懂这里为什么输出 “GOGOGO” ？？
}

*/
// 通过以上例子应该直到SetParent是什么样的行为，是不是像面向对象中的查找链？
//func(inj *injector)Invoke(f interface)([]reflect.Value, error){
//	t := reflect.TypeOf(f)
//
//	var in = make([]reflect.Value, t.NumIn()) // Panic if t is not kind of Func
//	for i := 0; i <t.NumIn(); i++{
//		argType := t.In(i)
//		val := inj.Get(argType)
//		if !val.IsValid(){
//			return nil, fmt.Errorf("value not found for type %v", argType)
//		}
//		in[i] = val
//	}
//	return reflect.ValueOf(f).Call(in), nil
//}
//
//Invoke方法用于动态执行函数，当然执行前可以通过Map或MapTo来注入参数，
//因为通过Invoke执行的函数会取出已注入的参数，然后通过reflect包中的Call方法来调用
//Invoke接收的参数f是一个接口类型，但是f的底层类型必须为func，否则会panic.

/*
type SpecialString interface{}
type cityy interface{}

func Say(name string, gender SpecialString, city cityy, age int) {
	fmt.Printf("My name is %s, gender is %s, age is %d\n", name, gender, age)
	fmt.Printf("My name is %s, gender is %s,city is %s, age is %d\n", name, gender, city, age)
}

func main() {
	inj := inject.New()
	inj.Map("张三")
	inj.MapTo("男", (*SpecialString)(nil))
	inj.MapTo("beijing", (*cityy)(nil))
	inj2 := inject.New()
	inj2.Map(25)
	inj.SetParent(inj2)
	inj.Invoke(Say)
	//运行后：My name is 张三, gender is 男, age is 25
}

*/

//上面的例子如果没有定义SpecialString接口作为gender参数的类型，而把name和gender都定义为string类型，那么gender会覆盖name的值

/*
func (inj *injectot)Apply(val interface{}) error{
	v := reflect.ValueOf(val)

	for v.Kind() == reflect.Ptr{
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct{
		return nil
	}

	t := v.Type()

	for i:=0; i<v.NumField(); i++{
		f := v.Field(i)
		structField := t.Field(i)
		if f.CanSet()&&structField.Tag == "inject"{
			ft := f.Type()
			v := inj.Get(ft)
			if !=v.IsValid(){
				return fmt.Errorf("Value not found for type %v", ft)
			}
			f.Set(v)
		}
	}
	return nil
}

*/

// Apply方法是用于对struct的字段进行注入，参数为指向底层类型为结构体的指针。
//可注入的前提是：字段必须是导出的（即字段名以大写字母开头），并且此字段的tag设置为`inject`。

type SpecialString interface {
}

type TestStruct struct {
	Name   string `inject`
	Nick   []byte
	Gender SpecialString `inject`
	uid    int           `inject`
	Age    int           `inject`
}

func main() {
	s := TestStruct{}
	inj := inject.New()
	inj.Map("张三")
	inj.MapTo("男", (*SpecialString)(nil))
	inj2 := inject.New()
	inj2.Map(26)
	inj.SetParent(inj2)
	inj.Apply(&s) //把上面的数据写入到s结构体中。
	fmt.Println("s.Name =", s.Name)
	fmt.Println("s.Gender =", s.Gender)
	fmt.Println("s.Age =", s.Age)
	fmt.Println("s=", s)

}
