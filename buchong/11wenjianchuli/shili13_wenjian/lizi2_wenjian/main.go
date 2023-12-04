package main

import (
	"bufio"
	"fmt"
	"os"
)

// 打开一个存在的文件，在原来的内容追加内容“c网”
func main() {
	filePath := "E://Geek/src/learn.go/buchong/11wenjianchuli/shili13_wenjian/golang.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	// 及时关闭file句柄
	defer file.Close()

	// 写入文件时，使用带缓存的*Writer
	write := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		write.WriteString("c网\r\n")
	}
	// Flush
	write.Flush()
}
