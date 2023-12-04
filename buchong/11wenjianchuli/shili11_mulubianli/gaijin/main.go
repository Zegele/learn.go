// 改进
//www.cdsy.xyz/computer/programme/golang/20210307/cd161507567010717.html

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// 下面为每一个walkDir的调用创建一个新的goroutine
//它使用sync.WaitGroup来为当前存活的walkDir调用计数，
//一个goroutine在计数器减为0的时候关闭fileSizes通道。

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

	// 并行遍历每一个文件树
	fileSizes := make(chan int64)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, &n, fileSizes)
	}
	go func() {
		n.Wait()
		close(fileSizes)
	}()

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
}

// walkDir 递归地遍历以dir为根目录的整个文件树，并在filesizes上发送每个已找到文件的大小
func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

// sema是一个用于限制目录并发数的计数信号量
var sema = make(chan struct{}, 20)

// dirents 返回directory目录中的条目
func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}        // 获取令牌
	defer func() { <-sema }() // 释放令牌
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}
