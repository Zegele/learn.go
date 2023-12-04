// Json数据编码和解码
// www.kancloud.cn/imdszxs/golang/1509744
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

//数据结构要在网络中传输或保存到文件，就必须对其编码和解码
//目前存在很多编码格式：JSON，XML，gob，Google缓冲协议等
//go支持多有这些编码格式
//结构可能包含二进制数据，如果将其作为文本打印， 那么可读性是很差的，
//另外结构内部可能包含匿名字段，而不清楚数据的用意
//通过把数据转换成纯文本，使用命名的字段来标注，让其具有可读性。这样的数据格式可以通过网络传输
//而且是与平台无关的，任何类型的应用都能够读取和输出
//不与操作系统和编程语言的类型相关
//数据结构--> 指定格式 ： 序列化 或 编码（传输之前）/
//指定格式--> 数据格式 ： 翻序列化 或 编码（传输之后）
//序列化是在内存中把数据转换成指定格式（data->string）,
//反之依然(string->data structure)编码也是一样的
//只是输出一个数据流（实现了io.Writer接口）
//解码是从一个数据流（实现了io.Reader）输出到一个数据结构

//JSON被作为首选，主要是由于其格式上非常简单。
//通常JSON被用于web后端和浏览器之间的通讯/
//{"Person":{"FirstName": "Laura","LastName": "Lynn"}}

//JSON更加简洁，轻量（占用更少的内存，磁盘及网络带宽） 更好的可读性
//go语言的json包可以在程序中方便的读取和写入JSON数据

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

func main() {
	pa := &Address{"private", "Aartselaar", "Belgium"}
	wa := &Address{"work", "boom", "Belgium"}
	vc := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}
	fmt.Printf("%v:\n", vc)

	js, _ := json.Marshal(vc)
	fmt.Printf("------JSON format: %s\n", js) // js是字节切片，居然可以直接用%s？？？
	// using an encoder:
	file, _ := os.OpenFile("vcard.json", os.O_CREATE|os.O_WRONLY, 0)
	defer file.Close()
	enc := json.NewEncoder(file) //json格式的
	err := enc.Encode(vc)        //直接对结构体进行json格式的编码，并写入文件

	if err != nil {
		log.Println("Error in encoding json")
	}
}

//json.Marshal() 函数签名是 func Marchal(v interface{})([]byte, error)
//返回的是编码后的JSON文本，实际上是个[]byte
//出于安全考虑，在web应用中最好使用json.MarshlforHTML()函数
//其对数据执行HTML转码，所以文本可以被安全地嵌在HTML标签中
//JSON与go类型对应如下：
//bool 对应JSON的 booleans
//float64 对应JSON的 numbers
//string 对应 JSON的 strings
//nil 对应JSON的 null
//不是所有的数据都可以编码为JSON 类型：只有验证通过的数据结构才能被编码 ：
//1. JSON对象只支持字符串类型的key； 要编码一个Go map类型，map必须是map[string]T (T是json包中支持的任何类型)
//2. Channel， 复杂类型和函数类型不能被编码
//3. 不支持循环数据结构，它将引起序列化进入一个无限循环
//4. 指针可以被编码，实际上是对指针指向的值进行编码（或者指针是nil）

//反序列化
//UnMarshal()的函数签名是func Unmarshal(data []byte, v interface{})error
//把JSON解码为数据结构
//我们首先创建一个结构Message用来保存解码的数据： var m Message并调用Unmarshal()
//解析[]byte中的JSON数据并将结果存入指针m指向的值。
//虽然反射能够让JSON字段去尝试匹配目标结构字段
//但是只有真正匹配上的字段才会填充数据
//字段没有匹配不会报错，而是直接忽略掉

//解码任意的数据：
//json包使用map[string]interface{}和interface{}存储任意的JSON对象和数组；
//其可以被反序列化为任何的JSON blob 存储到接口值中
//来看这个JSON数据，被存储在变量b中：
//b == []byte({"Name":"Wednesday","Age":6, "Parents":["Gomez","Morticia"]})
//不用理解这个数据的结构，我们可以直接使用Unmarshal把这个数据编码并保存在接口值中：
//var f interface{}
//err := json.Unmarshal(b, &f)
//f指向的值是一个map，key是一个字符串，value是自身存储作为空接口类型的值
//map[string]interface{"Name":"Wednesday", "Age":6, "Parents":[]interface{}{"Gomaz", "Morticia",},}
//要访问这个数据，可以使用类型断言
//m := f.(map[string]interface{}). // 参考interface接口类型断言
//我们可以通过for range 语法和type switch来访问其实际类型
//for k, v := range m{
//	swich vv := v.(type){
//		case string:
//			fmt.Println(k, "is string", vv)
//		case int:
// 			fmt.Println(k, "is int", vv)
//  	case []interface{}
//			fmt.Println(k, "is a slice:")
//			for i, u := range vv {
//				fmt.Println(fmt.Println(i, u))
//			}
//		default: fmt.Println(k, "is of a type I don't know how to handle")
//	}
//}
//通过这种方式，可以处理位置的JSON数据，同时可以确保类型安全

//解码数据到结构：
//如果我们事先知道JSON数据，可以定义一个适当的结构并对JSON数据反序列化
//type FamilyMember struct {
//	Name string
// 	Age int
// 	Parents []string
//}
//并对其反序列化
//var m FamilyMembererr := json.Ummarshal(b, &m)
//程序实际上是分配一个新的切片，
//这是一个典型的反序列化引用类型（指针、切片和map）的例子

//编码和解码流
//json包提供Decoder和Encoder类型来支持常用JSON数据流读写
//NewDecoder和NewEncoder函数分别封装了io.Reader和io.Writer接口
//func NewDecoder(r io.Reader)*Decoder
//func NewEncoder(w io.Writer)*Encoder
//要想把JSON直接写入文件，
//可以使用json.NewEncoder初始化文件（或者任何实现io.Writer的类型） ，并调用Encoder()
//反过来，使用json.NewDecoder 和Decode()函数，解码
//func NewDecoder(r io.Reader)*Decoder
//func (dec *Decoder)Decode(v interface{})error
//接口是如何对现实进行抽象的：数据结构可以是任何类型，只要其实现了某种接口。
//目标或源数据要能够被编码就必须实现io.Writer或io.Reader接口
//由于Go语言中导出都实现了Reader和Writer,
//因此Encoder和Decoder可被应用的场景非常广泛
//例如读取或写入HTTP连接，websockets或文件、
