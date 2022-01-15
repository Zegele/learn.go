package main

import "fmt"

func main() {
	var data string
	{
		var equipment IOInterface = &Soft{} //定义equipment是IOI接口类型，并与Soft结构体建立联系。这样Soft结构体就可以用接口中的方法。
		data = equipment.Read()
		fmt.Println(data)
	}
	{
		var equipment IOInterface = &Mag{}
		data = equipment.Read()
		fmt.Println(data)
	}
	{
		var equipment IOInterface = &Sata{}
		data = equipment.Read()
		fmt.Println(data)
	}
}

type IOInterface interface {
	Read() (data string)
}

type Sata struct{}

func (Sata) Read() string { //这是Sata结构体的方法（函数）（Read）。
	return "安安静静的sata"
}

type Soft struct {
}

func (Soft) Read() string { //这是Soft结构体的方法（Read）
	return "软盘数据，啦啦啦啦"
}

type Mag struct {
}

func (Mag) Read() string {
	return "滋滋滋， 磁带"
}

type Paper struct {
}

func (Paper) Read() string {
	return "从纸带读取01010....."
}
