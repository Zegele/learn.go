// Go语言使用Gob传输数据
// c.biancheng.net/view/4597.html
package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

// 为了让某个数据结构能够在网络上传输或能够保存至文件， 它必须被编码然后再解码
//go语言encoding/gob包提供了一种
//Gob是go语言自己以二进制形式序列化和反序列化程序数据的格式，可以在encoding包中找到。
//这种格式的数据简称为Gob（Go binary）
//类似于Python的“pickle“和Java的Serilization
//
//Gob由发送端使用encoder对数据结构进行编码，在接受到消息后，接收端使用decoder将序列化的数据便化成本地变量
//
//go语言可以通过JSON或Gob来序列化struct对象，虽然JSON的序列化更为通用，
//但利用Gob编码可以实现JSON所不能支持的struct的方法序列化，利用Gob包序列化struct保存到本地也十分简单
//
//Gob不是可外部定义，语言无关的编码方式，它的首选的是二进制格式，而不是像JSON或XML那样的文本格式
//Gob并不是一种不同于go的语言，
//而是在编码和解码过程中用到了go的反射
//
//Gob通常用于远程方法调用参数和结果的传输，以及应用程序和机器之间的数据传输。
//它和JSON或XML有什么不同？
//Gob特定的用于纯Go的环境中，例如两个用Go语言写的服务之间的通信。这样的话服务可以被实现得更加高效和优化。
//
//Gob文件或流是完全子描述的，它里面包含的所有类型都有一个对应的没描述，并且都是可以用Go语言解码，而不需要了解文件的内容。
//
//只有可导出的字段会被编码，零值会被忽略。
//在解码结构体的时候，只有同时匹配名称和可兼容类型的字段才会被解码
//当源数据类型增加新字段后，Gob解码客户端任然可以以这种方式正常工作。
//解码客户端会继续识别以前存在的字段，并且还提供了很大的灵活性
//比如在发送者看来，整数被编码成没有固定长度的可变长度，而忽略具体的Go类型
//假如有下面那这样一个结构体
//type T stuct{X,Y,X int}
//var t = T{X:7, Y:0, Z:8}
//而在接收时可以用一个结构体U类型的变量u来接收这个值：
//type U struct{X, Y *int8}
//var u U
//在接收时，X的值是7，Y的值是0（Y的值并没有从t中传递过来，因为它是零值）和JSON的使用方式一样，
//Gob使用通用的io.Writer接口，通过NewEncoder()函数创建Encoder对象并调用Encode()
//相反的过程使用通用的io.Reader接口，通过NewDecoder()含数创建Decoder对象并调用Decode，

// 创建Gob文件
// 下面通过简单的实例来演示Go语言是如何创建gob文件的：
type Book struct {
	Name string
	Page int
}

func main() {
	//info := map[string]string{ //map类型
	//	"name":    "golang",
	//	"website": "http://c.biancheng.net/golang/",
	//}
	info := Book{Name: "gogogo", Page: 100} // 结构体类型，必须大写
	name := "demo.gob"
	File, _ := os.OpenFile(name, os.O_RDWR|os.O_CREATE, 0777)
	defer File.Close()

	enc := gob.NewEncoder(File)
	if err := enc.Encode(info); err != nil {
		fmt.Println(err)
	}
}
