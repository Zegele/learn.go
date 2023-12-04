package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

//读XML文件比写XML文件稍微复杂，特别是在必须处理一些我们自定义字段的时候（例如日期）
//但是，如果我们使用合理的打上XML标签的结构体，就不会复杂，如下：。

type Website struct {
	Name   string `xml:"name,attr"` // 这个tag干嘛？把name对应的读出来
	Url    string
	Course []string
}

func main() {
	// 打开xml
	filePtr, err := os.Open("E:/Geek/src/learn.go/info.xml") // 路径要正确
	if err != nil {
		fmt.Printf("文件打开失败[Err:%s]", err.Error())
		return
	}
	defer filePtr.Close()

	//var info []Website //也可以这样写 注意这是结构体的切片
	info := Website{} //这是结构体，不是切片
	//创建json解码器
	decoder := xml.NewDecoder(filePtr)
	err = decoder.Decode(&info) //按&info格式解码
	if err != nil {
		fmt.Println("解码失败", err.Error())

	} else {
		fmt.Println("解码成功")
		fmt.Println(info) // 只能读取切片的第一个结构体。所以XML更适合结构体？
	}
}

// 正如写XML时一样，我们无需关心对所读取的XML数据进行转义
//xml.NewDecoder.Decode()函数会自动处理这些
//
//xml包还支持跟为复杂的标签，包括嵌套。
//例如标签名为'xml:"Books>Author"'产生的是 <Books><Author>content</Author></Books>这样的XML内容
//'xml:"xxx,attr"' 把属性值符给xxx 例如：xxx="golang"
//该包还支持'xml:",chardata"'这样的标签表示将该字段当做字符数据来写，
//支持'xml:",innerxml"'这样的标签表示按照字面量来写该字段，
//'xml:"comment"'这样的标签表示将该字段当做XML注释，
//因此，通过使用标签化的结构体，我们可以充分利用好这些方便得到编码解码函数，同时合理控制如何读写XML数据。
