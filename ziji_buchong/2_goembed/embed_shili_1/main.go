package main

import "embed"

//参考文档：
//https://blog.csdn.net/Naisu_kun/article/details/130722873#:~:text=%2F%2Fgo%3Aembed%20%E6%8C%87%E4%BB%A4%201%20%E8%A6%81%E5%B5%8C%E5%85%A5%E7%9A%84%E6%96%87%E4%BB%B6%E6%94%AF%E6%8C%81%E5%BD%93%E5%89%8D%E7%A8%8B%E5%BA%8F%20%28%20%2A.go%20%29%E6%89%80%E5%9C%A8%E7%9B%AE%E5%BD%95%E5%8F%8A%E5%AD%90%E7%9B%AE%E5%BD%95%EF%BC%9B%202,%E6%9D%A5%E5%8C%B9%E9%85%8D%E6%89%80%E6%9C%89%E6%96%87%E4%BB%B6%EF%BC%8C%E5%B9%B6%E4%B8%94%E4%BC%9A%E9%80%92%E5%BD%92%E5%AD%90%E7%9B%AE%E5%BD%95%E4%B8%AD%20.%20...%207%20%E5%8F%AF%E4%BB%A5%E4%BD%BF%E7%94%A8%20%2F%2Fgo%3Aembed%20%2A%20%E6%9D%A5%E8%A1%A8%E7%A4%BA%E5%BD%93%E5%89%8D%E7%9B%AE%E5%BD%95%EF%BC%9B
//1. 要嵌入的文件支持当前程序（xxx.go 不一定非是main.go，也可以某个包中的.go文件）所在目录及子目录
//2. 嵌入的文件系统是只读的，支持 Open ReadDir ReadFile 三个方法进行访问；
//3. 一条 //go:embed 指令后面可以写用空格隔开的多个文件，也可以用多条//go:embed 将文件内容放入一个变量中 如22-27行 .

import (
	_ "embed"
	"fmt"
)

//go:embed不能有空格
//go:embed abc.txt
var a string // 自动将abc.txt的内容嵌入到a中。

//go:embed aaa.txt
//go:embed bbb.txt
var f embed.FS //直接相当于打开了文件？
// f直接是文件，读文件就直接读f就行。
// 一个f嵌套了，两个文件

//go:embed aaa.txt bbb.txt
var ff embed.FS

//这个ff和上面的f是等价的。

func main() {
	fmt.Printf("version: %q\n", a)

	data, _ := f.ReadFile("aaa.txt")
	fmt.Println(string(data))

	data, _ = f.ReadFile("bbb.txt")
	fmt.Println(string(data))
}
