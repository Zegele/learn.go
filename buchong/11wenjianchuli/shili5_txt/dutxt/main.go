package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 读纯文本文件
func main() {
	// 打开文件
	file, err := os.Open("./output.txt")
	if err != nil {
		fmt.Println("文件打开失败 = ", err)
	}
	// 及时关闭file句柄，否则会有内存泄露
	defer file.Close()
	// 创建一个*Reader ，是带缓冲的
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n') // 读到一个换行就结束
		// 下次循环，会接着上次的地方继续读
		if err == io.EOF { // io.EOF表示文件的末尾
			break
		}
		fmt.Print(str)
	}
	fmt.Println("文件读取结束。。。")
}
