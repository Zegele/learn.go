// Go语言JSON文件的读写操作
// c.biancheng.net/view/4545.html
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

//JSON(JavaScript Object Notation)是一种轻量级的数据交换格式，易于阅读和编写
//同时也易于机器解析和生成。
//JSON是一种使用UTF-8编码的纯文本格式，采用完全独立于语言的文本格式，
//由于写起来比XML格式方便，并且更为紧凑，同时所需的处理时间也更少，只是JSON格式越来越流行
//特别实在通过网络连接传送数据方面
//
//开发人员可以使用JSON传出简单的字符串，数字，布尔值，也可以传输，一个数组或一个更复杂的复合结构
//在Web开发领域，JSON被广泛应用于Web服务端程序和客户端之间的数据通信
//
//Go语言内键对JSON的支持，使用内置的encoding/json标准库，开发人员可以轻松使用Go程序生成和解析JSON格式的数据
//
//JSON结构图如下：
//{"key1":"value1","key2":"value2","key3":["value3","value4","value5"]}，
//写JSON文件
//使用Go语言创建一个json文件非常方便，如下：

type Website struct {
	Name   string `xml:"name,attr"` // 这个tag 在例子中没有用
	Url    string
	Course []string
}

func main() {
	info := []Website{{"golang", "http://c.biancheng.net/golang/",
		[]string{"http://c.biancheng.net/cplus/", "http://c.biancheng.net/linux_tutouial/"}},
		{"Java", "http//c.biancheng.net/java/", []string{"http://c.biancheng.net/socket/", "http://c.biancheng.net/python/"}}}

	// 创建文件
	filePtr, err := os.Create("info.json") //文件名info.json
	// os.Create("info.json")文件在项目下， os.Create("./info.json")文件还是在项目下
	if err != nil {
		fmt.Println("文件创建失败", err.Error())
		return
	}
	defer filePtr.Close()

	// 创建Json编码器
	encoder := json.NewEncoder(filePtr)

	// 开始把数据写入
	err = encoder.Encode(info)
	if err != nil {
		fmt.Println("编码错误", err.Error())

	} else {
		fmt.Println("编码成功")
	}
}

//运行代码，会在项目下生成一个info.json文件。
