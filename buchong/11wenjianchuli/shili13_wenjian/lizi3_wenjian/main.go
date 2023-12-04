package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 示例3：打开一个存在的文件，将原来的内容读出来，显示在终端，并且追加5句 hello
func main() {
	filePath := "E://Geek/src/learn.go/buchong/11wenjianchuli/shili13_wenjian/golang.txt"
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}

	defer file.Close()

	//读原来文件的内容，并且显示在终端
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF { //io.EOF读到结尾了
			break
		}
		fmt.Print(str) // 文件内容自带换行
		//fmt.Println(str)// 多了一个换行
	}

	//写入文件时，使用带缓存的*Writer
	write := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		write.WriteString("Hello\r\n")
	}
	write.Flush()
}
