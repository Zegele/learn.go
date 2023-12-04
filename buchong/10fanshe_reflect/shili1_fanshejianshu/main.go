// go 反射（reflect）简述
// c.biancheng.net/view/4407.html
package main

import (
	"fmt"
	"reflect"
)

/*
// 一、反射的类型（Type）与种类（Kind）
// 1.1 type和kind的区别
func main() {
	var a int
	typeOfA := reflect.TypeOf(a) // 取得变量a的类型对象
	fmt.Println(typeOfA.Name(), typeOfA.Kind())

	type KKK func(int) int // KKK是类型
	// type是类型 kind是种类
	// Type指的是系统原生数据类型： 如 int， string， bool， float32等类型
	// 以及使用type关键字定义的类型。这些类型的名称就是其类型本身的名称。
	// 如type A struct{}  ，A就是struct{}的类型。

	// kind 指对象归属的品种。a := make(chan int) a属于chan种类。
	//type Kind uint
	//const (
	//	Invalid Kind = iota  // 非法类型
	//	Bool                 // 布尔型
	//	Int                  // 有符号整型
	//	Int8                 // 有符号8位整型
	//	Int16                // 有符号16位整型
	//	Int32                // 有符号32位整型
	//	Int64                // 有符号64位整型
	//	Uint                 // 无符号整型
	//	Uint8                // 无符号8位整型
	//	Uint16               // 无符号16位整型
	//	Uint32               // 无符号32位整型
	//	Uint64               // 无符号64位整型
	//	Uintptr              // 指针
	//	Float32              // 单精度浮点数
	//	Float64              // 双精度浮点数
	//	Complex64            // 64位复数类型
	//	Complex128           // 128位复数类型
	//	Array                // 数组
	//	Chan                 // 通道
	//	Func                 // 函数
	//	Interface            // 接口
	//	Map                  // 映射
	//	Ptr                  // 指针
	//	Slice                // 切片
	//	String               // 字符串
	//	Struct               // 结构体
	//	UnsafePointer        // 底层指针
	//)
	// type A struct{} 定义的结构体，属于Struct种类，*A数据Ptr
	//
	var k KKK

	k = ttt
	fmt.Println(k(1))

	typeOfk := reflect.TypeOf(k)

	fmt.Println("name:", typeOfk.Name(), "kind", typeOfk.Kind())
	fmt.Printf("k的类型：%T,反射的类型：%T\n", k, typeOfk)
}

func ttt(i int) int {
	return i
}


*/

/*
// 1.2 从类型对象中获取类型名称和种类
// go语言中的类型名称对应的反射获取方法是reflect.Type中的Name()方法，返回表示类型名称的字符串
// 类型归属的种类（Kind）使用的是reflect.Type中的Kind() 方法，返回reflect.Kind类型的常量。

type Enum int // 声明Enum类型

const (
	Zero Enum = 0 // 声明Zero是Enum类型的，并实例化赋值0
)

func main() {
	// 声明一个空结构体
	type cat struct { //声明结构体类型cat
	}
	// 获取结构体实例的反射类型对象
	typeOfCat := reflect.TypeOf(cat{}) //将cat实例化，并且使用reflect.TypeOf()获取被实例化后的cat的反射类型对象。

	// 显示反射类型对象的名称和种类
	fmt.Println(typeOfCat.Name(), typeOfCat.Kind())
	// 类型名称就是cat，而cat属于一种结构体忠烈，因此种类为struct

	// 获取Zero常量的反射类型对象
	typeOfA := reflect.TypeOf(Zero) // Zero是一个Enum类型的常量。

	// 显示反射类型对象的名称和种类
	fmt.Println(typeOfA.Name(), typeOfA.Kind())
}

*/

/*
// 二、指针与指针指向的元素
// go语言中对指针获取反射对象时，可以通过reflect.Elem()方法获取这个指针指向的元素类型。
// 这个获取过程被称为取元素，等效于对指针类型变量做了一个*操作。如下：

func main() {
	// 声明一个空结构体
	type cat struct {
	}
	// 创建cat的实例
	ins := &cat{} // ins是一个*cat类型的指针变量
	// 获取结构体实例的反射类型对象
	typeOfins := reflect.TypeOf(ins) // 对指针变量获取反射类型信息。
	// 显示反射类型对象的名称和种类
	fmt.Printf("name: '%v' kind:'%v'\n", typeOfins.Name(), typeOfins.Kind())
	// 输出指针变量的类型名称和种类
	// go语言的反射中对所有指针变量的种类都是Ptr，但需要注意的是，指针变量的类型名称是空，不是*cat

	// 取类型的元素
	typeOfCat := typeOfins.Elem() // 取指针类型的元素类型，也就是cat类型。
	// 这个操作不可逆，不可以通过一个非指针类型获取它的指针类型。

	//显示反射类型对象的名称和种类
	fmt.Printf("element name:'%v', element kind: '%v'\n", typeOfCat.Name(), typeOfCat.Kind())
	// 输出指针变量指向元素的类型名称和种类，得到了cat的类型名称（cat），和种类（struct）
}


*/

// 使用反射获取结构体的成员类型
// Field(i int) SturctField, NumField()int, FieldByName(name string)(StructField, bool), FieldByIndex(index []int)StructFiled, FieldByNameFunc(match func(string)bool)(StructField, bool)

func main() {
	// 声明一个结构体
	type cat struct { // 声明带有两个成员的cat结构体
		Name string
		// 带有结构体tag的字段
		Type int `json:"type" id:"100"` // 带了 // 标签（tag） 注意标签的 :
		// 结构体标签由一个或多个键值对组成，键与值使用冒号分割（注意不要乱加空格），值用双引号括起来；键值对之间使用一个空格分隔。
	}
	// 创建cat的实例
	ins := cat{Name: "mimi", Type: 1}
	// 获取结构体实例的反射类型对象
	typeOfCat := reflect.TypeOf(ins)
	inss := &cat{Name: "nn", Type: 100}
	typeOfCatPtr := reflect.TypeOf(inss)
	fmt.Println(typeOfCatPtr.Elem().Name(), typeOfCatPtr.Elem().Kind(), "-------")
	// 遍历结构体所有成员
	for i := 0; i < typeOfCat.NumField(); i++ { //NumField()获得一个结构体类型共有多少个字段。如果不是结构体，会宕机
		// 获取每个成员的结构体字段类型
		fieldType := typeOfCat.Field(i) // Field() 返回的是StructField结构体， 不是reflect.Type
		/*
			type StructField struct {
			    Name string          // 字段名
			    PkgPath string       // 字段在结构体中的路径
			    Type      Type       // 字段反射类型对象
			    Tag       StructTag  // 字段的结构体标签
			    Offset    uintptr    // 字段在结构体中的相对偏移
			    Index     []int      // Type.FieldByIndex中的返回的索引值
			    Anonymous bool       // 是否为匿名字段
			}
		*/
		// 输出成员名和tag
		fmt.Printf("name: %v tag: '%v' PkgPath: '%v'\n", fieldType.Name, fieldType.Tag, fieldType.PkgPath)
		fmt.Println(fieldType.Type)
	}
	// 通过字段名，找到字段类型信息
	if catType, ok := typeOfCat.FieldByName("Type"); ok {
		// FieldByName（）函数根据字段名查找结构体字段信息。
		// 从tag中取出需要的tag
		fmt.Println(catType.Tag.Get("json"), catType.Tag.Get("id"))
		// func(tag StrcutTag)Get(key string)string 根据Tag中的键获取对应的值
		// func(tag StructTag)Lookup(key string)(value string, ok bool) 根据Tag中的键，查询值是否存在。
	}
}
