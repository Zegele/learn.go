// Go语言zip归档文件的读写操作
// c.biancheng.net/view/4583.html
package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"os"
)

// Go语言的标准库提供了对几种压缩格式的支持，包括gzip
// 因此go程序可以无缝地读写.gz扩展名的gzip压缩文件或非.gz扩展名的非压缩文件。
// 此外标准库也提供了读和写.zip，tar包文件（.tar和.tar.gz）
// 以及读.bz2文件（即.tar.bz2文件）的功能
//
// 创建zip归档文件
// go语言提供了archive/zip包来操作压缩文件 ，示例如下： .
func main() {
	//创建一个缓冲区来保存压缩文件内容
	buf := new(bytes.Buffer)

	// 创建一个压缩文档
	w := zip.NewWriter(buf)
	//fmt.Println(w)
	// 将文件加入压缩文档
	var files = []struct {
		Name, Body string
	}{
		{"Golang.txt", "http://c.biancheng.net/golang/"},
		{"shouye.txt", "http://c.biancheng.net/"},
	}

	for _, file := range files {
		f, err := w.Create(file.Name) //在压缩文档中创建数据文档
		if err != nil {
			fmt.Println(err)
		}
		_, err = f.Write([]byte(file.Body)) // 写到了buf 也就是缓存
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println(w)
	}

	// 关闭压缩文档
	err := w.Close()
	if err != nil {
		fmt.Println(err)
	}

	//将压缩文档内容写入压缩文件
	f, err := os.OpenFile("file.zip", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
	}
	buf.WriteTo(f) //把缓存的数据写到压缩文件
}
