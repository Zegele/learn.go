// GO语言通过反射访问结构体成员的值
// c.biancheng.net/view/114.html
package main

import (
	"fmt"
	"reflect"
)

// 反射值对象（reflect.Value）提供对结构体访问的方法，通过这些方法可以完成对结构体任意值的访问：
// 反射值 对象的成员访问方法： 注意：与reflect.Type的方法做区别。
//Field(i int)Value // 根据索引，返回对应的结构体成员字段的反射值对象。
//NumField()int // 返回结构体字段成员数量
//FieldByName(name string) Value // 根据给定字符串返回字符串对应的结构体字段
//FieldByIndex(index []int)Value // 多层成员访问，[]int相当于字段路径，根据索引，返回字段的值。
//FieldByNameFunc(match func(string)bool)Value // 根据匹配函数 匹配需要的字段

// 定义结构体
type dummy struct {
	a int
	b string

	// 嵌入字段
	float32
	bool

	next *dummy
}

func main() {
	// 值包装结构体
	d := reflect.ValueOf(dummy{
		a:       1,
		b:       "khgk",
		float32: 3.0, // 匿名字段这样赋值！！
		bool:    true,
		next:    &dummy{},
	})
	fmt.Printf("d's value:'%v', d's type:'%v', d's kind:'%v'\n", d, d.Type(), d.Kind())

	// 获取字段数量
	fmt.Println("NumField", d.NumField())

	// 获取索引为2的字段（float32字段）
	floatField := d.Field(2)

	// 输出字段类型
	fmt.Println("Field", floatField.Type())

	// 根据名字查找字段
	fmt.Println("FieldByName(\"b\")的值：", d.FieldByName("b"))
	fmt.Println("FieldByName(\"b\").Type", d.FieldByName("b").Type())

	// 根据索引查找值中，next字段的int字段的值
	fmt.Println("FieldByIndex([]int{4,0})的值：", d.FieldByIndex([]int{4, 0}))
	fmt.Println("FieldByIndex([]int{4,0}).Type()", d.FieldByIndex([]int{4, 0}).Type())
	// dummy结构体中索引值为4的成员，也就是next。next的类型为*dummy，也是一个结构体
	//因此使用[]int{4,0}中的0继续再next值的基础上索引
	//*dummy中索引值为0的a字段，类型为int。
}
