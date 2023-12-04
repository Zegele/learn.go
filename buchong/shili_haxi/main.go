//go语言哈希函数
//www.kancloud.cn/imdszxs/golang/1509749

package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
)

//go语言提供了MD5，SHA-1等几种哈希函数，下面我们用例子做一个介绍：
/*
func main() {
	TestString := "http://c.biancheng.net/golang/"
	Md5Inst := md5.New()
	Md5Inst.Write([]byte(TestString))
	Result := Md5Inst.Sum([]byte(""))
	fmt.Printf("md5: %x\n", Result)

	Sha1Inst := sha1.New()
	Sha1Inst.Write([]byte(TestString))
	Result = Sha1Inst.Sum([]byte(""))
	fmt.Printf("sha1:%x\n", Result)
}

*/

// 文件内容计算MD5和SHA1,示例：
func main() {
	TestFile := "E:/Geek/src/learn.go/buchong/shili_haxi/123.txt"
	infile, inerr := os.Open(TestFile)
	if inerr == nil {
		md5h := md5.New()
		io.Copy(md5h, infile)
		fmt.Printf("MD5: %x %s\n", md5h.Sum([]byte("")), TestFile)

		sha1h := sha1.New()
		io.Copy(sha1h, infile)
		fmt.Printf("SHA1: %x %s\n", sha1h.Sum([]byte("")), TestFile)
	} else {
		fmt.Println(inerr)
		os.Exit(1)
	}

}
