// Go语言tar归档文件的读写操作
// c.biancheng.net/view/4584.html
package main

import (
	"archive/tar"
	"fmt"
	"io"
	"os"
)

// 创建tar归档文件
// tar是一种打包格式，但不对文件进行压缩，所以打包后的文档一般远远大于zip和tar.gz
//因额外i不需要压缩的原因，所以打包的速度是非常快的，打包时cpu占用率也很低
//
//tar的目的在于方便文件的管理，将零散的东西放一起
//
//创建tar归档文件与创建zip类似，主要不同点在于我们将所有数据都写入相同的writer中，
//并且在写入文件的数据之前必须写入完整的头部，
//而非仅仅是一个文件名
//
//tar打包实现原理如下：
//1.创建一个文件x.tar,然后向x.tar写入tar头部信息
//2. 打开要被tar的文件，向x.tar写入头部信息，然后向x.tar写入文件信息
//3. 当有多个文件需要被tar时，重复第二步直到所有文件都被写入到x.tar中
//4. 关闭x.tar，完成打包。 如下：
/*
func main() {
	f, err := os.Create("./output.tar") // 创建一个tar文件
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	tw := tar.NewWriter(f)
	defer tw.Close()

	fileinfo, err := os.Stat("./main.exe") //获取文件相关信息
	if err != nil {
		fmt.Println(err)
	}
	hdr, err := tar.FileInfoHeader(fileinfo, "")
	if err != nil {
		fmt.Println(err)
	}

	err = tw.WriteHeader(hdr) // 写入头文件信息
	if err != nil {
		fmt.Println(err)
	}

	f1, err := os.Open("./main.exe")
	if err != nil {
		fmt.Println(err)
		return
	}
	m, err := io.Copy(tw, f1) // 将main.exe文件中的信息写入压缩包中
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(m)
}


*/
func main() {
	f, err := os.Create("./output.tar") //创建一个 tar 文件
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	tw := tar.NewWriter(f)
	defer tw.Close()
	fileinfo, err := os.Stat("E:/Geek/src/learn.go/buchong/11wenjianchuli/shili9_tar/main.exe") //获取文件相关信息
	// os.Stat(要打包的文件) 获取要打包的文件
	if err != nil {
		fmt.Println(err)
	}
	hdr, err := tar.FileInfoHeader(fileinfo, "") // 获得文件信息的头部信息？
	if err != nil {
		fmt.Println(err)
	}
	err = tw.WriteHeader(hdr) //先写入头文件信息  头部信息对应文件信息
	if err != nil {
		fmt.Println(err)
	}
	f1, err := os.Open("E:/Geek/src/learn.go/buchong/11wenjianchuli/shili9_tar/main.exe")
	if err != nil {
		fmt.Println(err)
		return
	}
	m, err := io.Copy(tw, f1) //将main.exe文件中的信息写入压缩包中
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(m)
}
