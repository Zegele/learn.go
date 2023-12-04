// Go语言自定义二进制文件的读写操作
// c.biancheng.net/view/4570.html
package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
)

//虽然Go语言的encoding/gob包非常易用 而且所需代码量也非常少，
//但是我们仍由可能需要创建自定义的二进制格式。
//自定义的二进制格式有可能做到最紧凑的数据表示
//并且读写速度可以非常快
//
//不过，在实际使用中，我们发现以Go语言二进制格式的读写通常比自定义格式要快非常多
//而且创建的文件也不会大很多
//但如果我们必须通过满足 gob.GobEncoder和gob.GobDecoder接口来处理一些不可被gob编码的数据
//这些优势就有可能会失去
//
//在有些情况下我们可能需要与一些使用自定义二进制格式的软件交互，因此了解如何处理二进制文件就非常有用。
//
//写自定义二进制文件
//Go语言的encoding/binary 包中的binary.Write()函数使得以二进制格式写数据非常简单，函数原型如下：
//func Write(w io.Writer, order ByteOrder, data interface{})error
//Write函数可以将参数data的binary编码格式写入参数w中，参数data必须示定长值、定长值的切片、定长值的指针。
//参数order指定写入数据的字节序，写入结构体时，名字中有_的字段会置为0
//
//下面演示Write函数的使用：

type Website struct {
	Url int32
}

func main() {
	file, err := os.Create("output.bin")
	if err != nil {
		fmt.Println("文件创建失败", err.Error())
		return
	}
	defer file.Close()

	for i := 1; i <= 100; i++ {
		info := Website{
			int32(i),
		}

		var bin_buf bytes.Buffer
		binary.Write(&bin_buf, binary.LittleEndian, info)
		//把info写到bin_buf缓存了

		b := bin_buf.Bytes()   // 把缓存的数据换成[]byte格式
		_, err = file.Write(b) //把[]byte格式的数据写入文件 //这是继续写，不会抹掉 之前的内容
		if err != nil {
			fmt.Println("编码失败", err.Error())
			return
		}
	}
	fmt.Println("编码成功")
}
