// Go语言纯文本文件的读写操作
// c.biancheng.net/view/4556.html
package main

import (
	"bufio"
	"fmt"
	"os"
)

// 写纯文本文件
func main() {
	// 创建一个新文件，写入内容
	filePath := "./output.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("打开文件错误=%v \n", err)
		return
	}
	// 及时关闭
	defer file.Close()
	// 写入内容
	str := "http://c.biancheng/net/golang/\r\n" // \n\r表示隔行 txt文件要看到换行效果要用\r\n
	//\r是回车符，将光标移到当前行的行首；Carriage Return。
	//\n是换行符，光标换到下一行

	// 写入时，使用带缓存的*Writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 3; i++ {
		writer.WriteString(str)
	}
	// 因为writer是带缓存的，因此在调用WriterString方法时，内容时先写入缓存的
	// 所以要调用flush方法，将缓存的数据真正写入到文件中。
	writer.Flush()
}
