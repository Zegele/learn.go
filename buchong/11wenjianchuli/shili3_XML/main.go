// Go语言XML文件的读写操作
// c.biancheng.net/view/4551.html
package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

// XML(extensible Markup Language)格式被广泛用作一种数据交换格式，
//并且自成一种文件格式，比JSON复杂得多，而且手动写起来相对乏味。
//在JSON还未像现在广泛使用时，XML的使用相当广泛。
//XML作为一种数据交换和信息传递的格式，使用还是很广泛的，现在很多开放平台接口，基本都会支持XML格式
//
//go内置的encoding/xml包可以用在结构体和XML格式之间进行编码，
//其方式跟encoding/json包类似
//然而与JSON相比XML的编码和解码在功能上更苛刻得多，
//这是由于encoding/xml包要求结构体的字段包含格式合理的标签，而JSON格式却不需要.

//写XML文件

// 使用encoding/xml包可以很方便的将xml数据存储到文件中，如下：

type Website struct {
	Name   string `xml:"name,attr"` // 这个tag 编码后： name="golang" 而不是只有个 golang
	Url    string
	Course []string
}

func main() {
	//info := []Website{{"golang", "http://c.biancheng.net/golang/",
	//	[]string{"http://c.biancheng.net/cplus/", "http://c.biancheng.net/linux_tutouial/"}},
	//	{"Java", "http//c.biancheng.net/java/",
	//		[]string{"http://c.biancheng.net/socket/", "http://c.biancheng.net/python/"}}}
	//xml包针对结构体，如果是结构体类型的切片。存储是全部，但是读取只是第一个结构体

	info := Website{"golang", "http://c.biancheng.net/golang/",
		[]string{"http://c.biancheng.net/cplus/", "http://c.biancheng.net/linux_tutouial/"}}

	// 创建文件
	filePtr, err := os.Create("./info.xml") //文件名info.xml
	if err != nil {
		fmt.Println("文件创建失败", err.Error())
		return
	}
	defer filePtr.Close()

	// 序列化到文件中  创建xml编码器
	encoder := xml.NewEncoder(filePtr)

	// 开始把数据写入
	err = encoder.Encode(info)
	if err != nil {
		fmt.Println("编码错误", err.Error())
		return
	} else {
		fmt.Println("编码成功")
	}
}

//运行代码，会在项目下生成一个info.xml文件。
