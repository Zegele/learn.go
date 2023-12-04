// Go文件的写入、追加、读取、复制操作
// c.biancheng.net/view/5729.html
package main

import (
	"bufio"
	"fmt"
	"os"
)

// Go语言的os包下有一个OpenFile函数，其原型如下：
//func OpenFile(name string, flag int, perm FileMode)(file *File, err error)
//其中name是文件的文件名，如果不是在当前路径下运行需要加上具体路径
//flag是文件的处理参数，为int类型，根据系统的不同具体值可能有所不同，但是作用是相同的
//下面列举一些常用的flag文件处理参数：
//O_RDONLY: 只读模式打开文件
//O_WRONLY:只写模式打开文件
//O_RDWR: 读写模式打开文件
//O_APPEND: 写操作时将数据附加到文件尾部（追加）
//O_CREATE: 如果不存在将创建一个新文件
//O_EXCL: 和O_CREATE配合使用，文件必须不存在，否则返回一个错误
//O_SYNC: 当进行一系列写操作时，每次都要等待上次的I/O操作完成再进行
//O_TRUNC:如果可能，在打开时清空文件。

// 示例1：创建一个新文件golang.txt, 并在其中写入5句URL
func main() {
	// 创建一个新文件，写入内容5句
	filePath := "E://Geek/src/learn.go/buchong/11wenjianchuli/shili13_wenjian/golang.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	// 及时关闭file句柄
	defer file.Close()

	// 写入文件时，使用带缓存的*Writer
	write := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		write.WriteString("http://c.biancheng.net/golang/\n")
	}
	// Flush将缓存的文件真正写入到文件中
	write.Flush()
}
