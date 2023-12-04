// Go语言结构体标签（Struct Tag）
// c.biancheng.net/view/112.html
package main

import (
	"fmt"
	"reflect"
)

// 通过reflect.Type 获取结构体成员信息reflect.StuctField结构中的Tag被称为结构体标签（Struct Tag）
// 结构体标签是对结构体字段的额外信息标签
//JSON， BSON等格式进行序列化及对象关系映射（Object Relational Mapping，简称ORM）系统都会用到结构体标签，
//这些系统使用标签设定字段，在处理时应该具备的特殊属性和可能发生的行为
//这些信息都是静态的，无需实例化到结构体，可以通过反射获取到
//
//结构体标签的格式：
//`key1:"value1" key2:"value"`
//结构体标签
//1：由一个或多个键值对组成；
//2.键与值使用冒号分隔，且无空格，且值用双引号括起来
//3. 键值对之间使用一个空格分隔。

// 从结构体标签中获取值
// func(tag StructTag)Get(key string)string
// 根据Tag中的键获取对应的值，可以传入"key1"获得"value1"
// func(tag StructTag)Lookup(key string)(value string, ok bool)
// 根据Tag中的键，查询值是否存在
//
// 结构体标签格式错误导致的问题
// 编写Tag时，必须严格遵守键值对的规则。
// 结构体标签的解析代码的容错能力很差，一旦格式写错，编译和运行时都不会提示任何错误，如下：。
func main() {
	type cat struct {
		Name string
		Type int `json:"type" id:"100"` // 不要乱加空格，尤其冒号的左右。
	}

	typeOfCat := reflect.TypeOf(cat{})
	if catType, ok := typeOfCat.FieldByName("Type"); ok {
		fmt.Println(catType.Tag.Get("json"))
		fmt.Println(catType.Tag.Get("id"))
	}
	if catType, ok := typeOfCat.FieldByName("Type"); ok {
		if v, ok := catType.Tag.Lookup("json"); ok {
			fmt.Println(v, ok)
		}
	}
}
