package main

//没有使用编码的情况下，要自己处理。很可能会出错。

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	filePath := "e:/Geek/src/learn.go/chapter08/03.format/小强.self.infomation"
	writeFile(filePath)
	readFile(filePath)
}

func readFile(filePath string) {
	data, err := ioutil.ReadFile(filePath) // data 是读到的内容
	if err != nil {
		fmt.Println("读取文件失败：", err)
		return
	}
	infos := strings.Split(string(data), ",") //表示数据遇到 ","(逗号)，分割开。(小强，男 --> 小强 男) 分隔开就可以struct等使用。
	fmt.Println("开始计算体脂信息：", infos)
}

func writeFile(filePath string) {
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("无法打开文件", filePath, "错误信息是：", err)
		os.Exit(1)
	}
	defer file.Close()

	_, err = file.Write([]byte("小强，男，1.70，..."))
	fmt.Println(err)
}
