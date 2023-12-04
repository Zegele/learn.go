package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// 先看bufio包相关的Reader函数方法：
//
//首先定义了一个用来缓冲io.Reader对象的结构体，
//同时该结构体拥有以下相关的方法
//type Reader struct{
//}
//
//NewReader 函数用来返回一个默认大小buffer的Reader对象（默认大小是4096）等同于NewReaderSize(rd, 4096)
//func NewReader(rd io.Reader)*Reader
//
//该函数返回一个指定大小buffer（size最小为16）的Reader对象，
//如果io.Reader参数已经是一个足够大的Reader，它将返回该Reader
//func NewReaderSize(rd io.Reader, size int)*Reader
//
//Buffered方法返回从当前buffer中能被读到的字节数
//func (b *Reader)Buffered()int
//
//Discard方法跳过后续的n个字节的数据，返回跳过的字节数。
//如果0<=n<=b.Buffered(),该方法将不会从io.Reader 中成功读取数据
//func(b *Reader)Discaard(n int)(discarded int, err error)
//
//Peek方法返回缓存的一个切片，该切片只包含缓存中的前n个字节的数据
//func(b *Reader)peek(n int)([]byte, error)
//
//Read方法，把Reader缓存对象中的数据读入到[]byte类型的p中，并返回读取的字节数
//读取成功，err将返回空值
//func(b *Reader)Read(p []byte)(n int, err error)
//
//ReadByte返回单个字节，如果没有数据返回err
//func(b *Reader)ReadByte()(byte, error)
//
//ReadSlice:该方法在b 中读取delim之前的数据
//返回的切片是已读出的数据的引用（注意！！是引用），切片中的数据在下一次的读取操作之前是有效的
//如果未找到delim，将返回查找结果并返回nil空值
//因为缓存的数据可能被下一次的读写操作修改，因此一般使用ReadBytes或者ReadString，他们返回的都是数据拷贝
//func(b *Reader)ReadSlice(delim byte)(line []byte, err error)
//
//ReadBytes 功能同ReadSlice，返回数据的拷贝
//func (b *Reader)ReadBytes(delim byte)([]byte, error)
//
//功能同ReadBytes，返回字符串
//func(b *Reader)ReadString(delim byte)(string, error)
//
//ReadLine：该方法是一个低水平的读取方式
//一般建议使用ReadBytes('\n')或ReadString('\n')，或者使用一个Scanner来代替
//ReadLine通过调用ReadSlice方法实现，返回的也是缓存的切片，
//用于读取一行数据，不包括行尾标记（\n 或 \r\n）
//func(b *Reader)ReadLine() (line []byte, isPrefix bool, err error)
//
//读取单个 UTF-8 字符并返回一个rune和字节大小
//func(b *Reader)ReadRune() (r rune, size int, err error)。

func main() {
	fileObj, err := os.Open("demo.txt")
	if err != nil {
		fmt.Println("文件打开失败：", err)
		return
	}
	defer fileObj.Close()
	// 一个文件对象本身是实现了io.Reader的使用bufio.NewReader去初始化一个Reader对象
	//存在buffer中，读取一次就会被清空
	reader := bufio.NewReader(fileObj) // 文件的一些信息就对应reader了？
	buf := make([]byte, 1024)
	fmt.Println(buf)
	//读取Reader对象中的内容到[]byte类型的buf中
	info, err := reader.Read(buf) //读到的东西存在buf了

	fmt.Println(info, buf)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("读取的字节数：" + strconv.Itoa(info))
	//fmt.Println("读取的字节数：",info) // 这样也行
	// 这里的buf是一个[]byte，因此如果需要只输出内容，仍然需要将文件内容的换行符替换掉
	fmt.Println("读到的文件内容：", string(buf))
}
