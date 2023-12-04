package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
)

// 读取和创建一样简单，只是如果归档文件中包含带有路径的文件名，就必须重建目录结构
func main() {
	//打开一个zip格式文件
	r, err := zip.OpenReader("file.zip")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer r.Close()

	// 迭代压缩文件中的文件，打印出文件中的内容
	for _, f := range r.File {
		fmt.Printf("文件名：%s\n", f.Name)
		rc, err := f.Open()
		if err != nil {
			fmt.Printf(err.Error())
		}
		_, err = io.CopyN(os.Stdout, rc, int64((f.UncompressedSize64)))
		// 把rc(文件内容)拷贝到os.Stdout标准输出
		fmt.Println()
		if err != nil {
			fmt.Println(err.Error())
		}
		rc.Close()
	}
}
