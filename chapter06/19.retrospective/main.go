package main

import (
	"fmt"
	"time"
)

type Downloader struct {
	fileNameCh chan string
	// some kind of download worker
	finishedCh chan string
}

func (d *Downloader) Download() {
	for fileName := range d.fileNameCh { // 把fileNameCh中的数据遍历给fileName
		//	fmt.Println("开始下载文件：", fileName)
		time.Sleep(1 * time.Second)
		//	fmt.Println("开始处理文件：", fileName)
		time.Sleep(1 * time.Millisecond)
		//	fmt.Println("保存文件：", fileName)
		d.finishedCh <- fileName //遍历的文件名，又传递给了finishedCh
	}
}

func main() {
	fileNameCh := make(chan string, 20)
	finishedCh := make(chan string, 20) // 如果这个finished chan 没有buffer，则在download完成后无法将结果放进该channel，从而导致死锁。
	//没有buffer的channel必须有消费者才可以向channel内装数据。

	workerCounter := 50

	for i := 0; i < workerCounter; i++ {
		go func() {
			downloader := &Downloader{ //初始化这些也是在goroutine中进行的。。 学到没？？？
				fileNameCh: fileNameCh,
				finishedCh: finishedCh,
			}
			downloader.Download() // 在该goroutine中实例化后，再调用
			// 为什么不把实例化写到函数里？因为下载函数，就做下载的事。实例化和下载没有必然联系。
		}()
	}

	fileNumber := 100
	fileNames := make([]string, 0, fileNumber)
	for i := 0; i < fileNumber; i++ { //生成所有名字的切片
		fileNames = append(fileNames, fmt.Sprintf("file_%d.txt", i))
	}

	finishedFileCount := 0
	go func() {
		for finishedFile := range finishedCh { //这里又把finishedCh中的数据接住
			fmt.Println("文件：", finishedFile, "处理完毕")
			finishedFileCount++
			if finishedFileCount == fileNumber {
				close(finishedCh) //在遍历channel的时候，再close该channel。学到了么？？
			}
		}
	}()

	for _, fileItem := range fileNames { //把所有文件名字装到channel中
		fileNameCh <- fileItem
	}
	close(fileNameCh)

	fmt.Println("所有文件处理完毕，结束")
	time.Sleep(2 * time.Second)
}
