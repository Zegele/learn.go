package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	filePath := "e:/Geek/src/learn.go/chapter08/01.read-file/小强.txt"
	//filePath := "e://Geek//src//learn.go//chapter08//01.read-file//小强.txt"
	//filePath := "e:\\Geek\\src\\learn.go\\chapter08\\01.read-file\\小强.txt"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("无法打开文件：", filePath, "错误信息是：", err)
		os.Exit(1) //如果程序正常退出，Exit内的值是0，如果不是正常退出，给一个非0的数。
		// 可以在命令终端 用 $? 看程序是否正常退出。 如 echo $?  打印出的值如果是0 说明程序可以正常退出。
	}
	defer file.Close()
	b := make([]byte, 10) //这个slice不能是空的，即不能是：make([]byte, 0, 1024)
	// 一般设置为4096的整数倍 这样内存更高效一些
	for i := 0; i < 2; i++ { //for 循环，可能会存在读取断开。如一个中文 读取断开后就错了。
		n, err := file.Read(b) //对file进行读取，读到b字节切片中。n表示读到了多少个。
		if err != nil {
			fmt.Println("无法读取文件：", err)
			os.Exit(2)
		}
		fmt.Println("读出的内容(原生)：", b)
		fmt.Println("读出的内容（转string后）：", string(b))
		fmt.Println("读出的内容（转string后），做了截取：", string(b[:n])) //这种更严谨。因为string()转的时候对字节切片的0值是忽略的，但该位置是有0值的。
		// 所以读了多少，打印多少这样更严谨。
		// 一定要记得给后续程序使用时，截取到实际读取到的数据，而不是全部。否则后续程序会把无效读取也作为正常数据处理。
		fmt.Println("n的大小（读到多少）：", n)
		//	file.Seek(0, io.SeekStart) //io.SeekStart 从头开始读  //offset 0 表示偏移0个位置。
		file.Seek(3, io.SeekCurrent) //io.SeekCurrent 从目前游标的位置开始， //offset 3 表示再向后偏移3个位置，再读取。
		// todo handle error
	}
}

// 终端命令： du -h -d 1 深度为1的文件（夹）的大小
// du -sh 显示当前所在的文件夹的大小
// du -s 也是显示当前所在文件夹的大小（人不好识别）
