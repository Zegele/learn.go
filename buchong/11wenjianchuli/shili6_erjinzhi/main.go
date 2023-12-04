// Go语言二进制文件的读写操作
// c.biancheng.net/view/4563.html
package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

// Go语言的二进制（gob）格式是一个自描述的二进制序列
// 从其内部表示来看，Go语言的二进制格式由一个0块或者更多块的序列组成，
// 其中的每一块都包含一个字节数，一个由0个或者多个typeld-typeSpecification对组成的序列，以及一个typeld-value对
//
// 如果typeld-value对的typeld是语言定义好的（例如bool,int 和string等），
// 则这些typeld-typeSpecification对可以省略
// 否则就用类型对来描述一个自定义类型（如一个自定义的结构体）
// 类型对和值之间的typeld没有区别
//
// 正如我们将看到的，我们无需了解其内部结构就可以使用gob格式，
// 因为encoding/gob包会在幕后为我们打理好一切底层细节
//
// Go语言中的encoding/gob包也提供了与encoding/json包一样的编码解码功能，且易用
// 通常如果肉眼可读性不做要求，gob格式是Go语言上用于文件存储和网络传输最为方便的格式
//
// 写Go语言二进制文件 ：
func main() {
	info := "http://c.biancheng.net/golang"
	file, err := os.Create("./output.gob")
	if err != nil {
		fmt.Println("文件创建失败", err.Error())
		return
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(info)
	if err != nil {
		fmt.Println("编码错误", err.Error())
		return
	} else {
		fmt.Println("编码成功")
	}
}
