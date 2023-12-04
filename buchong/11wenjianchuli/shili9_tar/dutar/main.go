package main

import (
	"archive/tar"
	"fmt"
	"io"
	"os"
)

// 解压tar归档文件比创建tar归档文件稍微简单些
//首先需要将其打开，然后从这个tar头部中循环读取存储在这个归档文件内的文件头部信息，从这个文件头里读取文件名，
//以这个文件名创建文件，然后向这个文件里写入数据即可。

func main() {
	f, err := os.Open("output.tar")
	if err != nil {
		fmt.Println("文件打开失败", err)
		return
	}

	defer f.Close()

	r := tar.NewReader(f)
	for hdr, err := r.Next(); err != io.EOF; hdr, err = r.Next() {
		if err != nil {
			fmt.Println(err)
			return
		}
		fileinfo := hdr.FileInfo()
		fmt.Println(fileinfo.Name())
		f, err := os.Create("123" + fileinfo.Name())
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()
		_, err = io.Copy(f, r)
		if err != nil {
			fmt.Println(err)
		}
	}
}

// 完成对tag介绍
//Go语言使用io.Reader，io.ReadClose, io.Writer和io.WriteCloser等
//接口处理文件的方式让开发者可以使用相同的编码模式来读写文件或者其它流（如网络流或者甚至时字符串）
//从而大大降低了难度。
