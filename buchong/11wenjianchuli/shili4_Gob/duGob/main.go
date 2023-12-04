package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

// 读取gob文件

func main() {
	type Book struct {
		Name string
		Page int
	}
	var M Book //结构体类型 名字和字段必须首字母大写
	//var M map[string]string // map类型
	File, _ := os.Open("demo.gob")
	D := gob.NewDecoder(File)
	D.Decode(&M)
	fmt.Println(M)
}
