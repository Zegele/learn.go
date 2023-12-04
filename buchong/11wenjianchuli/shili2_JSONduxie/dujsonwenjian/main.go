package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// 读JSON数据与写JSON数据一样简单
type Website struct {
	Name   string `xml:"name, attr""` // 这个tag干嘛？没用
	Url    string
	Course []string
}

func main() {
	filePtr, err := os.Open("E:/Geek/src/learn.go/info.json") // 路径要正确
	if err != nil {
		fmt.Printf("文件打开失败[Err:%s]", err.Error())
		return
	}
	defer filePtr.Close()
	var info []Website
	//创建json解码器
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(&info) //按&info格式解码
	if err != nil {
		fmt.Println("解码失败", err.Error())

	} else {
		fmt.Println("解码成功")
		fmt.Println(info)
	}
}

// 还有一种叫BSON（Binary JSON）的格式与JSON非常类似，与JSON相比
//BSON着眼于提高存储和扫描效率。
//BSON文档中的大型元素以长度字段为前缀以便于扫描
//在某些情况下，由于长度前缀和显示数组索引的存在，BSON使用的空间会多于JSON，
