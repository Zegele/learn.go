// go语言使用buffer读取文件
// c.biancheng.net/view/4595.html
package main

import (
	"bufio"
	"fmt"
	"os"
)

//buffer是缓冲器的意思，Go语言要实现缓冲读取需要使用到bufio包
//bufio包本身包装了io.Reader和io.Writer 对象
//同时创建了另外的Reader和Writer对象
//因此对于文本I/O来说，bufio包提供了一定的便利性
//
//buffer缓冲器的实现原理是：将文件读取进缓冲（内存）之中，再次读取的时候就可以避免文件系统的i/o从而提高速度
//同理在进行写操作时，先把文件写入缓冲（内存），然后由缓冲写入文件系统
//
//使用bufio包写入文件
// bufio和io包中有很多操作都是相似的，唯一不同的地方是bufio提供了一些缓冲的操作，
//如果对文件i/o操作比较频繁，使用bufio包能够提高一定的性能
//
//在bufio包中，有一个Writer结构体，而其相关的方法支持一些写入操作，如下：
//
//Writer是一个空的结构体，一般需要使用NewWriter或者NewWriterSize来初始化一个结构体对象
//type Writer struct{
// // contains filtered or unexported fields
//}
//
//NewWriterSize 和 NewWriter 函数
//返回默认缓冲大小的Writer对象（默认是4096）
//func NewWriter(w io.Writer)*Writer
//
//指定缓冲大小创建一个Writer对象
//func NewWriterSize(w io.Writer, size int)*Writer
//
//Writer 对象相关的写入数据方法
//
//把p中的内容吸入buffer，返回写入的字节数和错误信息
//如果nn<len(p)， 返回错误信息中会包含为什么写入数据比较短
//func (b *Writer)Write(p []byte)(nn int, err error)
//
//将buffer中的数据写入io.Writer
//func(b *Writer)Flush()error
//
//以下三个方法可以直接写入到文件中
//写入单个字节
//func (b *Writer)WriteByte(c byte)error
//写入单个Unicode指针返回写入字节错误信息
//func (b *Writer)WriteRune(r rune)(size int, err error)
//写入字符串并返回写入字节数和错误信息
//func (b *Writer)WriteString(s string)(int, error).

func main() {
	name := "demo.txt"
	content := "http://c.biancheng.net/golang/"

	fileObj, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}

	defer fileObj.Close()

	writeObj := bufio.NewWriterSize(fileObj, 4096)

	// 使用Write方法，需要使用Writer对象的Flush方法将buffer中的数据刷到磁盘
	buf := []byte(content)
	if _, err := writeObj.Write(buf); err == nil { //写到缓冲
		if err = writeObj.Flush(); err != nil { //把缓冲里的写到文件
			panic(err)
		}
		fmt.Println("数据写入成功")
	}
}
