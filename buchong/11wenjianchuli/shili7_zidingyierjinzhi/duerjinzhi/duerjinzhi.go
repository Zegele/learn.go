package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
)

type Website struct {
	Url int32
}

func main() {
	//file, err := os.Open("E:/Geek/src/learn.go/output.bin") //可行
	file, err := os.Open("output.bin") //直接在项目里找？
	defer file.Close()
	if err != nil {
		fmt.Println("文件打开失败", err.Error())
		return
	}
	m := Website{}
	for i := 1; i <= 100; i++ {
		data := readNextBytes(file, 4) //返回读满的4个字节
		// 为什么4就能解出，小于4没有解出乱码？大于4解出乱码
		fmt.Println(data)
		fmt.Printf("---%d\n", data)

		buffer := bytes.NewBuffer(data)                    // 把读到的4个字节，放到缓存里
		err = binary.Read(buffer, binary.LittleEndian, &m) // 把缓存的数据读到&m中
		//binary.BigEndian (大端模式)：内存的低位地址存放着数据高位
		//binary.LittleEndian (小端模式)：内存的地位地址存放着数据地位
		// 如果编码用小端，解码用大端，就会出错，反之依然，所以要一致
		if err != nil {
			fmt.Println("二进制文件读取失败", err)
			return
		}
		fmt.Println("第", i, "个值为：", m)
	}
}

func readNextBytes(file *os.File, number int) []byte {
	bytes := make([]byte, number)
	_, err := file.Read(bytes) // 把文件读到bytes中
	if err != nil {
		fmt.Println("解码失败", err)
	}
	return bytes
}

//至此，我们完成了自定义二进制数据的读和写操作
//只要小心选择表示长度的整数符号和大小
//并将该长度值写在变长值（如切片）的内容之前，
//那么使用二进制数据进行工作并不难
//
//go语言对二进制文件的支持还包括随机访问
//这种情况下，我们必须使用os.OpenFile()函数来打开文件(而非os.Open())
//并给它传入合理的权限标志和模式（如：os.O_RDWR表示可读写）参数
//
//然后，就可以使用os.File.Seek()方法在文件中定位并读写，
//或者使用os.File.ReadAt() 和 os.File.WriteAt()方法来从特定的字节偏移中读取或写入数据
//
//Go语言还提供了其他常用的方法，包括os.File.Stat()方法
//它返回的os.FileInfo包含了文件大小，权限以及日期时间等细节信息。.
