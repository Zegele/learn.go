package main

import (
	"fmt"
	"io/ioutil"
)

// 编写一个程序，将一个文件的内容复制到另外一个文件（这两个文件都已存在）
func main() {
	file1Path := "E://Geek/src/learn.go/buchong/11wenjianchuli/shili13_wenjian/golang.txt"
	file2Path := "E://Geek/src/learn.go/buchong/11wenjianchuli/shili13_wenjian/copyed.txt"
	data, err := ioutil.ReadFile(file1Path)
	if err != nil {
		fmt.Printf("文件打开失败=%v\n", err)
		return
	}
	err = ioutil.WriteFile(file2Path, data, 0666)
	if err != nil {
		fmt.Printf("文件打开失败=%v\n", err)
	}
}

// ioutil是干啥的？？搞一搞清楚
//www.jianshu.com/p/0e0ad501a145
