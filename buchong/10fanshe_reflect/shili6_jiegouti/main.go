// go语言通过反射获取结构体的成员类型
// c.biancheng.net/view/111.html
package main

import (
	"fmt"
	"reflect"
)

//任意值（非接口类型）通过reflect.TypeOf()获得反射对象信息后，如果它的类型（type）是结构体，
//可以通过反射值对象（reflect.Type）的NumFiled()和Field()方法获得结构体成员的详细信息
// Field(i int)StructField 索引结构体字段的信息
// NumField()int 返回结构体字段数量
// FieldByName(name string)(StructField, bool) 根据name返回字段信息
// FieldByIndex(index []int)SturctField 多层访问时，根据[]int提供的每个结构体的字段索引，返回字段信息。
// FieldByNameFunc(match func(string)bool)(StructField, bool) 根据匹配函数，匹配需要的字段。
//
//结构体字段类型
//reflect.Type的Field()方法返回StructField结构，这个结构体描述结构体的成员信息
//通过这个StructField信息可以获取成员与结构体的关系，如偏移，是否为匿名字段，结构体标签（struct tag）等
//而且还可以通过StructField的Type字段进一步获取结构体成员（字段）的类型。
//StructField的结构如下：
//type StructField struct{
//Name string  //字段名
//PkgPath string //字段路径
//Type Type //字段反射类型对象
//Tag StructTag //字段的结构体标签
//Offset uintptr //字段在结构体中的相对偏移
//Index []int //type.FieldByIndex中的返回的索引值
//Anonymous bool //是否为匿名字段.

func main() {
	// 声明一个空结构体
	type cat struct {
		Name string

		// 带有结构体tag的字段
		Type int `json:"type" id:"100"` //带有tag
	}

	// 创建cat实例
	ins := cat{Name: "mimi", Type: 1}

	// 获取结构体的反射类型对象
	typeOfCat := reflect.TypeOf(ins)                                                                        // 传入是cat
	typeOfCatPtr := reflect.TypeOf(&ins)                                                                    //传入是*cat
	fmt.Printf("typeOfCat's Kind:'%v', typeOfCatPtr's Kind: '%v'\n", typeOfCat.Kind(), typeOfCatPtr.Kind()) //ptr
	fmt.Printf("typeOfCatPtr-第一个字段的 StructField信息：%v\n", typeOfCatPtr.Elem().Field(0))

	// 遍历结构体所有成员
	for i := 0; i < typeOfCat.NumField(); i++ {
		// 获取每个成员的结构体字段类型
		fieldType := typeOfCat.Field(i)
		fieldTypePtr := typeOfCatPtr.Elem().Field(i)

		// 输出成员名和tag
		fmt.Printf("name: %v tag: %v\n", fieldType.Name, fieldType.Tag)
		fmt.Printf("*name: %v *tag: %v\n", fieldTypePtr.Name, fieldTypePtr.Tag)
	}

	// 通过字段名， 找到字段类型信息
	if catType, ok := typeOfCat.FieldByName("Type"); ok { //FieldByName(name string)根据name查找结构体字段信息，返回StructField
		//从tag中取出需要的tag
		fmt.Printf("json's tag:'%v', id's tag:'%v'\n", catType.Tag.Get("json"), catType.Tag.Get("id"))
		//Tag.Get()使用StructField中的Tag的Get()方法，根据Tag中的名字进行信息获取。
	}

	if catType, ok := typeOfCatPtr.Elem().FieldByName("Type"); ok {
		//从tag中取出需要的tag
		fmt.Printf("json ptr's tag:'%v', id ptr's tag:'%v'\n", catType.Tag.Get("json"), catType.Tag.Get("id"))
	}
}
