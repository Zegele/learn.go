// 并发目录遍历
// www.cdsy.xyz/computer/programme/golang/20210307/cd161507567010717.html
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

//我们将构建一个程序，根据命令行指定的输入，
//报告一个或多个目录的磁盘使用情况
//类似于UNIX的du命令
//该程序大多数工作是由下面的worldDir函数完成，
//它使用dirents辅助函数来枚举目录中的条目，如下。

// walkDir 递归地遍历以dir为根目录的整个文件树，并在filesizes上发送每个已找到文件的大小
func walkDir(dir string, fileSizes chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

// dirents 返回dir目录中的条目
func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

//ioutil.ReadDir函数返回一个os.FileInfo类型的slice
//针对单个文件同样的信息可以通过调用os.Stat函数来返回
//对每一个子目录，walkDir递归调用它自己，
//对于每一个文件，walkDir发送一条消息到fileSizes通道，消息的内容为文件所占用的字节数
//
//程序的完整代码如下所示，代码中main函数使用两个goroutine，
//后台goroutine调用walkDir遍历命令行指定的每一个目录
//最后关闭fileSizes 通道
//主goroutine计算从通道中接收的文件的大小的和，最后输出总数。
/*
func main() {
	// 确定初始目录
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// 遍历文件树
	fileSizes := make(chan int64)
	go func() {
		for _, root := range roots {
			walkDir(root, fileSizes)
		}
		close(fileSizes)
	}()

	//输出结果
	var nfiles, nbytes int64
	for size := range fileSizes {
		nfiles++
		nbytes += size
	}
	printDiskUsage(nfiles, nbytes)// 最终总数
}

// go run main.go D:/
// 就是统计D盘  这个 ’D:/‘就是命令行的第一个参数

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files %.1f GB\n", nfiles, float64(nbytes)/1e9) //1mb=1024kb 1kb =1024bytes
	// 1e9 = 10的9次方 （10亿） ≈ 1024*1024*1024
}

*/

// 如果程序可以通知它的进度，将会更友好，但是仅把printDiskUsage调用移动到循环内部会使它输出数千行结果
//所以这里对上面的程序进行一些调整 ，在有-v标识的时候周期性的输出当前目录的总和，
//如果只想看到最终的结果省略-v即可，

var verbose = flag.Bool("v", false, "显示详细进度")

func main() {

	// 启动后台 goroutine
	//确定初始目录
	flag.Parse()         //对命令行参数进行解析，所以要在第一个？
	roots := flag.Args() // 返回命令行参数后的其他参数，以[]string类型
	fmt.Println(roots)
	fmt.Println(flag.NArg())  // 返回命令行参数 后的 其他参数个数
	fmt.Println(flag.NFlag()) // 返回使用的命令行参数 如： -v 就是一个命令行参数

	//flag包：恢复记忆
	//flag.TypeVar(Type指针, flag名, 默认值, 帮助信息)
	//flag.IntVar(&age, "age", 18, "年龄+")

	if len(roots) == 0 {
		roots = []string{"."} // 搞了一个默认值？
	}

	// 遍历文件树
	fileSizes := make(chan int64)
	go func() {
		for _, root := range roots {
			walkDir(root, fileSizes)
		}
		close(fileSizes)
	}()

	//增加的！！
	// 定期打印结果
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(80 * time.Millisecond)
	}

	//输出结果
	var nfiles, nbytes int64
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop // fileSizes 关闭
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}

	printDiskUsage(nfiles, nbytes) // 最终总数
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files %.1f GB\n", nfiles, float64(nbytes)/1e9) //1mb=1024kb 1kb =1024bytes
	// 1e9 = 10的9次方 （10亿） ≈ 1024*1024*1024
}

//因为这个程序没有使用range循环，
//所以第一个select情况必须显示判断fileSizes通道是否已经关闭
//使用两个返回值的形式进行接收操作
//如果通道已经关闭，程序退出循环
//标签化的break语句将跳出select和for循环逻辑
//没有标签的break只能跳出select的逻辑 ，导致循环的下一次迭代
//
//此程序的弊端很明显，它耗时太长
//
//。
