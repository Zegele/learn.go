package main

import (
	"fmt"
	"os"
)

func main() {
	filePath := "e:/Geek/src/learn.go/chapter08/02.write-file/小强.txt"
	//filePath := "e://Geek//src//learn.go//chapter08//01.read-file//小强.txt"
	//filePath := "e:\\Geek\\src\\learn.go\\chapter08\\01.read-file\\小强.txt"
	//writeFile(filePath)
	writeFileWithAppend(filePath)

}

func writeFile(filePath string) {
	file, err := os.Create(filePath) //创建一个文件
	if err != nil {
		fmt.Println("无法打开文件：", filePath, "错误信息是：", err)
		os.Exit(1) //如果程序正常退出，Exit内的值是0，如果不是正常退出，给一个非0的数。
		// 可以在命令终端 用 $? 看程序是否正常退出。 如 echo $?  打印出的值如果是0 说明程序可以正常退出。
	}
	defer file.Close() // close

	_, err = file.Write([]byte("this is first line---")) //无换行
	fmt.Println(err)
	_, err = file.Write([]byte("this is first line\n"))
	fmt.Println(err)
	_, err = file.WriteString("第二行内容\n")
	fmt.Println(err)
	_, err = file.WriteAt([]byte("CHANGE"), 0) // offset 0 表示在开头开始写 (会产生内容覆盖的效果)
	file.Sync()                                //告诉操作系统，强制把已写到缓存（内存）的内容写到磁盘上  这个写入到磁盘的速度是比较慢的
	fmt.Println(err)
}

func writeFileWithAppend(filePath string) {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777) //创建一个文件
	if err != nil {
		fmt.Println("无法打开文件：", filePath, "错误信息是：", err)
		os.Exit(1) //如果程序正常退出，Exit内的值是0，如果不是正常退出，给一个非0的数。
		// 可以在命令终端 用 $? 看程序是否正常退出。 如 echo $?  打印出的值如果是0 说明程序可以正常退出。
	}
	defer file.Close() // close

	_, err = file.Write([]byte("this is first line---")) //无换行
	fmt.Println(err)
	_, err = file.Write([]byte("this is first line\n"))
	fmt.Println(err)
	_, err = file.WriteString("第二行内容\n")
	fmt.Println(err)
	//_, err = file.WriteAt([]byte("CHANGE"), 0) // Append打开的文件，不能用WriteAt写在前面。只能继续往后写。

	//fmt.Println(err)
}
